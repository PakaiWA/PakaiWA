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
	"github.com/PakaiWA/PakaiWA/internal/handler"
	"github.com/PakaiWA/PakaiWA/internal/pakaiwa"
	"github.com/PakaiWA/PakaiWA/internal/repository"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/KAnggara75/scc2go"
	"github.com/PakaiWA/PakaiWA/internal/configs"
	"github.com/PakaiWA/PakaiWA/internal/helpers"
	"go.mau.fi/whatsmeow"
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
	container := repository.InitStoreWithPool(ctx, pool, log)
	deviceStore, err := container.GetFirstDevice(ctx)
	helpers.PanicIfError(err)

	clientLog := pakaiwa.NewPakaiWALog(log, "PakaiWA")
	client := whatsmeow.NewClient(deviceStore, clientLog)

	state := &pakaiwa.AppState{Client: client}

	// Event Handler
	client.AddEventHandler(handler.EventHandler)

	// QR Channel, This must be called *before* Connect().
	var qrChan <-chan whatsmeow.QRChannelItem
	if client.Store.ID == nil {
		qrChan, _ = client.GetQRChannel(ctx)
	}
	// Start connection
	helpers.PanicIfError(client.Connect())

	// QR Handler
	pakaiwa.StartQRHandler(state, qrChan)

	// ====== App & Routes (Fiber) ======
	app := configs.NewFiber()
	configs.Bootstrap(
		&configs.BootstrapConfig{
			PakaiWA: state,
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
	state.Client.Disconnect()
	log.Println("Bye!")

	//
	//// Healthcheck
	//app.Get("/healthz", func(c *fiber.Ctx) error {
	//	return c.JSON(fiber.Map{
	//		"ok":        true,
	//		"connected": state.Client.IsConnected(),
	//	})
	//})
	//
	//app.Post("/v1/messages", func(c *fiber.Ctx) error {
	//	if !state.Client.IsConnected() {
	//		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
	//			"error": "whatsapp_not_connected",
	//		})
	//	}
	//
	//	var req struct {
	//		JID   string `json:"jid"`
	//		Phone string `json:"phone_number"`
	//		Text  string `json:"message"`
	//	}
	//
	//	if err := c.BodyParser(&req); err != nil {
	//		return fiber.NewError(fiber.StatusBadRequest, "invalid json: "+err.Error())
	//	}
	//	if strings.TrimSpace(req.Text) == "" {
	//		return fiber.NewError(fiber.StatusBadRequest, "`text` wajib diisi")
	//	}
	//
	//	phoneNumber := strings.TrimSpace(req.Phone)
	//	if phoneNumber == "" {
	//		return fiber.NewError(fiber.StatusBadRequest, "harus menyertakan `jid` atau `phone`")
	//	}
	//
	//	jid, err := helpers.NormalizeJID(phoneNumber)
	//	if err != nil {
	//		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	//	}
	//
	//	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//	defer cancel()
	//
	//	msg := &waE2E.Message{
	//		Conversation: helpers.ProtoString(req.Text),
	//	}
	//
	//	response, err := state.Client.SendMessage(ctx, jid, msg)
	//	if err != nil {
	//		return fiber.NewError(fiber.StatusBadGateway, "gagal mengirim: "+err.Error())
	//	}
	//	return helpers.RespondPending(c, response.ID)
	//})

}
