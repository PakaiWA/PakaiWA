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
	"github.com/KAnggara75/scc2go"
	"github.com/PakaiWA/PakaiWA/internal/configs"
	"github.com/PakaiWA/PakaiWA/internal/helpers"
	_ "github.com/jackc/pgx/v5/stdlib"
	"os"
)

func init() {
	scc2go.GetEnv(os.Getenv("SCC_URL"), os.Getenv("AUTH"))
}

func main() {
	ctx := context.Background()
	log := configs.NewLogger()
	pool := configs.NewDatabase(ctx, log)

	// ====== WhatsApp Client ======
	pwa, err := configs.NewWhatsAppClient(ctx, pool, log)
	helpers.PanicIfError(err)

	// ====== App & Routes (Fiber) ======
	app := configs.NewFiber()
	configs.Bootstrap(
		&configs.BootstrapConfig{
			PakaiWA: pwa,
			Pool:    pool,
			App:     app,
			Log:     log,
		},
	)

	go func() {
		addr := ":8080"
		if err := app.Listen(addr); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Graceful shutdown
	helpers.WaitForSignal()
	log.Println("Shutting down...")
	_ = app.Shutdown()
	pwa.Client.Disconnect()
	log.Println("Bye!")

}
