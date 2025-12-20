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
 * @author KAnggara75 on Fri 08/08/25 08.29
 * @project PakaiWA pakaiwa
 * https://github.com/PakaiWA/PakaiWA/tree/main/cmd/pakaiwa
 */

package main

import (
	"context"
	"os"

	"github.com/KAnggara75/scc2go"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/apperror"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/http/middleware"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/bootstrap"
	"github.com/PakaiWA/PakaiWA/internal/pkg/db"
	"github.com/PakaiWA/PakaiWA/internal/pkg/httpserver"
	"github.com/PakaiWA/PakaiWA/internal/pkg/kafka"
	"github.com/PakaiWA/PakaiWA/internal/pkg/logger"
	"github.com/PakaiWA/PakaiWA/internal/pkg/redis"
	"github.com/PakaiWA/PakaiWA/internal/pkg/utils"
	"github.com/PakaiWA/PakaiWA/internal/pkg/validator"
)

func init() {
	scc2go.GetEnv(os.Getenv("SCC_URL"), os.Getenv("AUTH"))
}

func main() {
	ctx := context.Background()

	log := logger.NewLogger()
	pool := db.NewDatabase(ctx, log)

	validate := validator.NewValidator()
	redis := redis.NewRedisClient()

	// ====== Kafka Producer ======
	producer := kafka.NewKafkaProducer(log)

	// ====== WhatsApp Client ======
	pwa, err := bootstrap.InitWhatsapp(&bootstrap.PwaContext{
		Log:      log,
		Pool:     pool,
		Producer: producer,
	})
	apperror.PanicIfError(err)

	// ====== App & Routes (Fiber) ======
	fiber := httpserver.NewFiber()
	fiber.Use(middleware.FiberLogger(log))
	bootstrap.InitApp(&bootstrap.AppContext{
		Log:      log,
		Pool:     pool,
		Redis:    redis,
		Fiber:    fiber,
		PakaiWA:  pwa,
		Producer: producer,
		Validate: validate,
	})

	go func() {
		addr := ":8080"
		if err := fiber.Listen(addr); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Graceful shutdown
	utils.WaitForSignal()
	log.Println("Shutting down...")
	_ = fiber.Shutdown()
	producer.Close()
	pwa.Client.Disconnect()
	log.Println("Bye!")
}
