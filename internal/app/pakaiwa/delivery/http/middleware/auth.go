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
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"

	"github.com/PakaiWA/PakaiWA/internal/pkg/config"
)

func AuthMiddleware(log *logrus.Logger, authFailLimiter *RateLimiter) fiber.Handler {
	return func(c fiber.Ctx) error {
		ip := c.IP()
		authHeader := c.Get("Authorization")

		fail := func(msg string) error {
			key := "auth_fail:" + ip

			if !authFailLimiter.isAllowed(key) {
				log.WithField("ip", ip).
					Warn("auth failure rate limit exceeded")

				return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
					"error": fiber.Map{
						"title":  "Too Many Requests",
						"status": 429,
						"detail": "Please slow down, you're hitting the rate limit",
					},
				})
			}

			log.WithField("ip", ip).Warn(msg)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": msg,
			})
		}

		if authHeader == "" {
			return fail("missing Authorization header")
		}

		parts := strings.Fields(authHeader)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return fail("invalid Authorization header format")
		}

		secretKey := config.GetJWTKey()
		if secretKey == "" {
			log.Fatal("JWT_SECRET is not set")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "server misconfiguration",
			})
		}

		token, err := jwt.Parse(parts[1], func(t *jwt.Token) (any, error) {
			if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, fiber.ErrUnauthorized
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			return fail("invalid or expired token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return fail("invalid token claims")
		}

		if jti, ok := claims["jti"].(string); ok {
			c.Locals("jti", jti)
		}

		// Optional: exp sudah divalidasi jwt.Parse,
		// tapi ini defensif dan eksplisit (boleh dipertahankan)
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				return fail("token expired")
			}
		}

		// Auth sukses → reset counter (opsional tapi direkomendasikan)
		authFailLimiter.Reset("auth_fail:" + ip)

		c.Locals("user", claims)

		log.WithFields(logrus.Fields{
			"sub":  claims["sub"],
			"role": claims["role"],
			"path": c.Path(),
			"ip":   ip,
		}).Info("authenticated request")

		return c.Next()
	}
}
