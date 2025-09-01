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
 * @author KAnggara75 on Mon 01/09/25 21.11
 * @project PakaiWA app
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app
 */

package app

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
	Fiber   *fiber.App
	Log     *logrus.Logger
}

func Bootstrap(b *BootstrapConfig) {
	qrHandler := handler.NewQRHandler(b.PakaiWA, b.Log)
	msgHandler := handler.NewMessageHandler(b.PakaiWA, b.Log)

	routeConfig := router.RouteConfig{
		Fiber:          b.Fiber,
		MessageHandler: msgHandler,
		QRHandler:      qrHandler,
	}
	routeConfig.Setup()
}
