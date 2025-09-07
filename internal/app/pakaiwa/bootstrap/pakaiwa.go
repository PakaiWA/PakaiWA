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
 * @author KAnggara75 on Sun 07/09/25 00.22
 * @project PakaiWA bootstrap
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/bootstrap
 */

package bootstrap

import (
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/http/handler"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/http/router"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/state"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/usecase"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type AppContext struct {
	Log      *logrus.Logger
	Pool     *pgxpool.Pool
	Fiber    *fiber.App
	PakaiWA  *state.AppState
	Producer *kafka.Producer
	Validate *validator.Validate
}

func InitApp(b *AppContext) {
	qrHandler := handler.NewQRHandler(b.PakaiWA, b.Log)

	// Message
	msgUsecase := usecase.NewMessageUsecase(b.Log, b.Validate, b.PakaiWA.Client)
	msgHandler := handler.NewMessageHandler(msgUsecase, b.Log)

	routeConfig := router.RouteConfig{
		Fiber:          b.Fiber,
		MessageHandler: msgHandler,
		QRHandler:      qrHandler,
	}
	routeConfig.Setup()
}
