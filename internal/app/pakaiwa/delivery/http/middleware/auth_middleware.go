/*
 * Copyright (c) 2025 KAnggara75
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * See <https://www.gnu.org/licenses/gpl-3.0.html>.
 *
 * @author KAnggara75 on Sat 08/11/25 22.06
 * @project PakaiWA middleware
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/delivery/http/middleware
 */

package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/model"
	"github.com/PakaiWA/PakaiWA/internal/pkg/config"
	"github.com/PakaiWA/PakaiWA/internal/pkg/logger/ctxmeta"
	"github.com/PakaiWA/PakaiWA/internal/pkg/utils"
)

func AuthMiddleware(authFailLimiter *RateLimiter) fiber.Handler {
	return func(c fiber.Ctx) error {
		log := ctxmeta.Logger(c.Context())
		ip := c.IP()

		key := "auth_fail:" + ip + ":" + c.Get("User-Agent")
		fail := func(msg string, status int) error {
			if !authFailLimiter.isAllowed(key) {
				log.WithField("ip", ip).Warn("auth rate limit exceeded")
				return utils.TooManyRequests(c)
			}

			log.WithField("ip", ip).Warn(msg)
			return c.Status(status).JSON(fiber.Map{"error": msg})
		}

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return fail("missing authorization header", 401)
		}

		parts := strings.Fields(authHeader)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return fail("invalid authorization format", 401)
		}

		secretKey := config.GetJWTKey()
		if secretKey == "" {
			log.Error("JWT_SECRET not configured")
			return c.SendStatus(500)
		}

		token, err := jwt.ParseWithClaims(parts[1], &model.JWTClaims{}, func(t *jwt.Token) (any, error) {
			if t.Method != jwt.SigningMethodHS256 {
				return nil, fiber.ErrUnauthorized
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			return fail("invalid or expired token", 401)
		}

		claims := token.Claims.(*model.JWTClaims)

		c.Locals("auth_user", model.AuthUser{
			Sub:  claims.Subject,
			Role: claims.Role,
			JTI:  claims.ID,
		})

		c.Locals("jwt_claims", &model.JWTClaims{
			QuotaLimit:    claims.QuotaLimit,
			WindowSeconds: claims.WindowSeconds,
		})

		authFailLimiter.Reset(key)
		return c.Next()
	}
}

func RequireAdmin() fiber.Handler {
	return func(c fiber.Ctx) error {

		user, ok := c.Locals("auth_user").(model.AuthUser)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
		}

		if user.Role != "admin" {
			return fiber.NewError(fiber.StatusForbidden, "admin only")
		}

		return c.Next()
	}
}
