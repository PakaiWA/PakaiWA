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
		bucket := now / int64(windowSeconds)

		key := fmt.Sprintf(
			"quota:%s:%d",
			user.Sub, // user_id
			bucket,   // window bucket
		)

		ctx := c.Context()

		used, err := rdb.Incr(ctx, key).Result()
		if err != nil {
			return fiber.NewError(
				fiber.StatusInternalServerError,
				"quota service unavailable",
			)
		}

		// TTL = window_seconds * 2 (aman dari race)
		ttl := time.Duration(windowSeconds*2) * time.Second
		_ = rdb.Expire(ctx, key, ttl).Err()

		remaining := max(limit-int64(used), 0)

		// Header observability
		c.Set("X-Quota-Remaining", strconv.Itoa(int(remaining)))
		c.Set("X-Quota-Limit", strconv.Itoa(int(limit)))
		c.Set("X-Quota-Window", strconv.Itoa(int(windowSeconds)))

		if used > int64(limit) {
			return fiber.NewError(
				fiber.StatusTooManyRequests,
				"quota exceeded",
			)
		}

		return c.Next()
	}
}
