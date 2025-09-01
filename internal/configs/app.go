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
 * @author KAnggara75 on Sat 30/08/25 13.07
 * @project PakaiWA configs
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/configs
 */

package configs

import (
	"github.com/PakaiWA/PakaiWA/internal/handler"
	"github.com/PakaiWA/PakaiWA/internal/pakaiwa"
	"github.com/PakaiWA/PakaiWA/internal/router"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type BootstrapConfig struct {
	PakaiWA *pakaiwa.AppState
	Pool    *pgxpool.Pool
	App     *fiber.App
	Log     *logrus.Logger
}

func Bootstrap(config *BootstrapConfig) {

	msgHandler := handler.NewMessageHandler(config.PakaiWA, config.Log)

	routeConfig := router.RouteConfig{
		App:            config.App,
		MessageHandler: msgHandler,
	}
	routeConfig.Setup()
}
