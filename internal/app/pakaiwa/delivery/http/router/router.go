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

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/http/dto"
	"github.com/PakaiWA/PakaiWA/internal/pkg/metrics"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/http/middleware"
	"github.com/PakaiWA/PakaiWA/internal/pkg/config"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/http/handler"
)

type RouteConfig struct {
	Fiber          *fiber.App
	QRHandler      *handler.QRHandler
	AuthHandler    *handler.AuthHandler
	MessageHandler *handler.MessageHandler
}

func (c *RouteConfig) Setup() {
	c.setupPublicRoutes()
	c.setupGuestRoutes()
	c.setupAuthRoutes()
}

func (c *RouteConfig) setupPublicRoutes() {

	c.Fiber.Get("/",
		middleware.RateLimitMiddleware(3, time.Minute),
		func(ctx fiber.Ctx) error {
			baseUrl := ctx.BaseURL()
			res := dto.VersionRes{
				Message:   baseUrl + " - Unofficial WhatsApp Restful API Gateway",
				Version:   config.GetAppVersion(),
				Stability: config.GetAppDesc(),
			}
			return ctx.JSON(res)
		},
	)

	c.Fiber.Get("/metrics",
		middleware.RateLimitMiddleware(30, time.Minute),
		metrics.PrometheusHandler(),
	)
}

func (c *RouteConfig) setupGuestRoutes() {
	c.Fiber.Post("/auth/login",
		middleware.RateLimitMiddleware(10, time.Minute),
		c.AuthHandler.Login,
	)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.Fiber.Use(middleware.AuthMiddleware(middleware.NewRateLimiter(5, time.Minute))) // Auth Middleware
	c.Fiber.Post("/register", middleware.RateLimitMiddleware(5, time.Minute*1), c.AuthHandler.Register)
	// Grouped Auth Routes V1
	v1 := c.Fiber.Group("/v1", middleware.RateLimitMiddleware(9999, time.Minute*1))
	v1.Post("/messages", c.MessageHandler.SendMsg)
}

func (c *RouteConfig) setupAuthRoutes() {
	// =====================
	// Base authenticated routes
	// =====================
	auth := c.Fiber.Group(
		"/",
		middleware.RateLimitMiddleware(1000, time.Minute),
		middleware.AuthMiddleware(middleware.NewRateLimiter(5, time.Minute)),
	)

	// logout (authenticated user)
	auth.Post("/logout", middleware.RateLimitMiddleware(5, time.Minute), c.AuthHandler.Register)

	// =====================
	// Admin-only routes
	// =====================
	admin := auth.Group("/", middleware.RequireAdmin())
	admin.Post("/register", middleware.RateLimitMiddleware(5, time.Minute), c.AuthHandler.Register)

	// =====================
	// User API v1
	// =====================
	v1 := auth.Group("/v1") //  middleware.QuotaMiddleware(c.Redis, 100),
	v1.Post("/messages", c.MessageHandler.SendMsg)
}
