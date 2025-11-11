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
 * @author KAnggara75 on Sat 06/09/25 09.06
 * @project PakaiWA router
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/delivery/http/router
 */

package router

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/http/dto"
	"github.com/PakaiWA/PakaiWA/internal/pkg/metrics"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/http/middleware"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/model"

	"github.com/PakaiWA/PakaiWA/internal/pkg/config"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/http/handler"
)

type RouteConfig struct {
	Fiber          *fiber.App
	QRHandler      *handler.QRHandler
	MessageHandler *handler.MessageHandler
}

func (c *RouteConfig) Setup() {
	c.NoLimitRoute()
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) NoLimitRoute() {
	c.Fiber.Get("/", middleware.RateLimitMiddleware(3, time.Minute*1), func(ctx fiber.Ctx) error {
		baseUrl := ctx.BaseURL()
		res := dto.VersionRes{
			Message:   baseUrl + " - Unofficial WhatsApp Restful API Gateway",
			Version:   config.GetAppVersion(),
			Stability: config.GetAppDesc(),
		}
		return ctx.JSON(res)
	})

	c.Fiber.Get("/metrics",
		middleware.RateLimitMiddleware(3, time.Minute*1),
		metrics.PrometheusHandler(),
	)
}

func (c *RouteConfig) SetupGuestRoute() {
	c.Fiber.Post("/auth/login",
		middleware.RateLimitMiddleware(10, time.Minute*1),
		GenerateJWT(),
	)

	c.Fiber.Post("/logout",
		middleware.RateLimitMiddleware(3, time.Minute*1),
		metrics.PrometheusHandler(),
	)
}

func (c *RouteConfig) SetupAuthRoute() {
	//c.Fiber.Use(middleware.AuthMiddleware())
	//c.Fiber.Use(middleware.AuthMiddleware()) // Quota Middleware
	auth := c.Fiber.Group("/v1", middleware.RateLimitMiddleware(9999, time.Minute*1))
	auth.Post("/messages", c.MessageHandler.SendMsg)
}

func GenerateJWT() fiber.Handler {
	return func(c fiber.Ctx) error {
		claims := jwt.MapClaims{
			"sub":  "userID",
			"role": "role",
			"exp":  time.Now().Add(1 * time.Hour).Unix(),
			"iat":  time.Now().Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		result, _ := token.SignedString([]byte(config.GetJWTKey()))

		response := model.SendMessageResponse{
			Message: result,
		}
		c.Status(200)
		return c.JSON(response)
	}
}
