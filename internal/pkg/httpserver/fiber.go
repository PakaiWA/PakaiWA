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
 * @author KAnggara75 on Sat 06/09/25 10.59
 * @project PakaiWA httpserver
 * https://github.com/PakaiWA/PakaiWA/tree/main/pkg/httpserver
 */

package httpserver

import (
	"github.com/PakaiWA/PakaiWA/internal/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func NewFiber() *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName:      config.GetAppName(),
		Prefork:      config.GetPreFork(),
		ErrorHandler: NewErrorHandler(),
	})

	return app
}
