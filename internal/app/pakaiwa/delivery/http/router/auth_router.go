/*
 * Copyright (c) 2026 KAnggara
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * See <https://www.gnu.org/licenses/gpl-3.0.html>.
 *
 * @author KAnggara on Thursday 01/01/2026 21.22
 * @project PakaiWA
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/delivery/http/router
 */

package router

import (
	"time"

	"github.com/gofiber/fiber/v3"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/http/middleware"
)

func RegisterAuthGroup(app *fiber.App) fiber.Router {
	return app.Group(
		"",
		middleware.RateLimitMiddleware(1000, time.Minute),
		middleware.AuthMiddleware(middleware.NewRateLimiter(5, time.Minute)),
	)
}
