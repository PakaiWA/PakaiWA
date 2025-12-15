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

func AuthMiddleware(log *logrus.Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			log.Warn("missing Authorization header")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing Authorization header",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			log.Warn("invalid Authorization header format")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid Authorization header format",
			})
		}

		tokenStr := parts[1]
		secretKey := config.GetJWTKey()
		if secretKey == "" {
			log.Fatal("JWT_SECRET is not set in environment")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "server misconfiguration",
			})
		}

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Warn("unexpected signing method")
				return nil, fiber.NewError(fiber.StatusUnauthorized, "invalid token")
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			log.WithError(err).Warn("invalid JWT token")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid or expired token",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Warn("invalid JWT claims type")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token claims",
			})
		}

		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				log.Warn("token expired")
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "token expired",
				})
			}
		}

		c.Locals("user", claims)

		log.WithFields(logrus.Fields{
			"sub":  claims["sub"],
			"role": claims["role"],
			"path": c.Path(),
		}).Info("authenticated request")

		return c.Next()
	}
}
