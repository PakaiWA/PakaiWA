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
 * @author KAnggara on Thursday 01/01/2026 21.21
 * @project PakaiWA
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/delivery/http/router
 */

package router

import (
	"time"

	"github.com/gofiber/fiber/v3"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/http/dto"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/http/middleware"
	"github.com/PakaiWA/PakaiWA/internal/pkg/config"
	"github.com/PakaiWA/PakaiWA/internal/pkg/metrics"
)

func RegisterPublicRoutes(app *fiber.App) {
	app.Get("/",
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

	app.Get("/metrics",
		middleware.RateLimitMiddleware(30, time.Minute),
		metrics.PrometheusHandler(),
	)
}
