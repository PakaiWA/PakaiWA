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

func QuotaMiddleware(rdb *redis.Client, limit int) fiber.Handler {
	return func(c fiber.Ctx) error {

		user, ok := c.Locals("auth_user").(*model.AuthUser)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized)
		}

		userID := user.Sub
		now := time.Now()

		key := fmt.Sprintf(
			"quota:v1:%s:%s",
			userID,
			now.Format("200601021504"), // window per menit
		)

		ctx := c.Context()

		used, err := rdb.Incr(ctx, key).Result()
		if err != nil {
			return fiber.NewError(
				fiber.StatusInternalServerError,
				"quota service unavailable",
			)
		}

		// TTL > window untuk aman
		_ = rdb.Expire(ctx, key, 2*time.Minute).Err()

		remaining := limit - int(used)
		if remaining < 0 {
			remaining = 0
		}

		// ⬅⬅⬅ HEADER PENTING
		c.Set("X-Quota-Remaining", strconv.Itoa(remaining))

		if used > int64(limit) {
			return fiber.NewError(
				fiber.StatusTooManyRequests,
				"quota exceeded",
			)
		}

		return c.Next()
	}
}
