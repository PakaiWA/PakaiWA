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
 * @author KAnggara75 on Sat 30/08/25 08.47
 * @project PakaiWA router
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/routers
 */

package router

import (
	"github.com/PakaiWA/PakaiWA/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	Fiber          *fiber.App
	QRHandler      *handler.QRHandler
	MessageHandler *handler.MessageHandler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.Fiber.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message":   "API PakaiWA.my.id - Unofficial WhatsApp API Gateway",
			"version":   "0.0.1",
			"stability": "Developer-Preview",
		})
	})

	c.Fiber.Get("/v1/qr", c.QRHandler.GetQR)
	c.Fiber.Get("/v1/qr/show", c.QRHandler.ShowQR)
	c.Fiber.Post("/v1/messages", c.MessageHandler.SendMsg)
}

func (c *RouteConfig) SetupAuthRoute() {
	//c.App.Use(c.AuthMiddleware)
}
