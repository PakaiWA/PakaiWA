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
 * @project PakaiWA routers
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/routers
 */

package routers

import (
	"github.com/PakaiWA/PakaiWA/internal/app/devices"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App              *fiber.App
	DeviceController *devices.DeviceController
	AuthMiddleware   fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/devices", c.DeviceController.List)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
}
