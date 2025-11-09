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
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/pkg/httpserver
 */

package httpserver

import (
	"github.com/gofiber/fiber/v3"

	"github.com/PakaiWA/PakaiWA/internal/pkg/config"
)

func NewFiber() *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName:            config.GetAppName(),
		ErrorHandler:       NewErrorHandler(),
		TrustProxy:         true,
		EnableIPValidation: true,
		TrustProxyConfig: fiber.TrustProxyConfig{
			Proxies: []string{
				"10.0.0.0/8",     // internal cluster network
				"172.16.0.0/12",  // Docker / Pod CIDR
				"192.168.0.0/16", // local ranges
			},
		},
	})

	return app
}
