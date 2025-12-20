/*
 * Copyright (c) 2025 KAnggara
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * See <https://www.gnu.org/licenses/gpl-3.0.html>.
 *
 * @author KAnggara on Thursday 18/12/2025 19.20
 * @project PakaiWA
 * ~/work/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/http/middleware
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/delivery/http/middleware
 */

package middleware

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/model"
)

func QuotaMiddleware(rdb *redis.Client) fiber.Handler {
	return func(c fiber.Ctx) error {
		user, ok := c.Locals("auth_user").(*model.AuthUser)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized)
		}

		claims, ok := c.Locals("jwt_claims").(*model.JWTClaims)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized)
		}

		limit := claims.QuotaLimit            // contoh: 100
		windowSeconds := claims.WindowSeconds // contoh: 60

		now := time.Now().Unix()
		window := int64(windowSeconds)

		bucket := now / window
		windowStart := bucket * window
		windowEnd := windowStart + window

		key := fmt.Sprintf(
			"quota:messages:%s:%d",
			user.Sub,
			bucket,
		)

		ctx := c.Context()

		used, err := rdb.Incr(ctx, key).Result()
		if err != nil {
			return fiber.NewError(
				fiber.StatusInternalServerError,
				"quota service unavailable",
			)
		}

		if used == 1 {
			ttl := time.Duration(window*2) * time.Second
			_ = rdb.Expire(ctx, key, ttl).Err()
		}

		remaining := max(limit-used, 0)

		c.Set("X-Quota-Limit", strconv.FormatInt(limit, 10))
		c.Set("X-Quota-Window", strconv.FormatInt(window, 10))
		c.Set("X-Quota-Remaining", strconv.FormatInt(remaining, 10))

		if used > limit {
			retryAfter := max(windowEnd-now, 0)
			c.Set("Retry-After", strconv.FormatInt(retryAfter, 10))
			c.Set("X-Quota-Reset", strconv.FormatInt(windowEnd, 10))

			return fiber.NewError(
				fiber.StatusTooManyRequests,
				"quota exceeded",
			)
		}

		return c.Next()
	}
}
