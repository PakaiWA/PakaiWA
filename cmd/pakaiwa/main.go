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
	_ "github.com/mattn/go-sqlite3"

	"fmt"
	"github.com/KAnggara75/scc2go"
	"github.com/PakaiWA/PakaiWA/internal/configs"
	"github.com/PakaiWA/PakaiWA/internal/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/mdp/qrterminal/v3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

func init() {
	scc2go.GetEnv(os.Getenv("SCC_URL"), os.Getenv("AUTH"))
}

type AppState struct {
	Client    *whatsmeow.Client
	QRMu      sync.RWMutex
	LastQR    string
	Connected bool
}

func eventHandler(e interface{}) {
	switch v := e.(type) {
	case *events.Message:
		msg := v.Message
		fmt.Println("Received a message!", v.Message.GetConversation())
		if msg.GetConversation() != "" {
			log.Printf("[INCOMING] from %s: %s", v.Info.Sender.String(), msg.GetConversation())
		}
	}
}

func main() {
	ctx := context.Background()

	dbLog := waLog.Stdout("Database", "DEBUG", true)
	container, err := sqlstore.New(ctx, "sqlite3", "file:store.db?_foreign_keys=on", dbLog)
	helpers.PanicIfError(err)

	deviceStore, err := container.GetFirstDevice(ctx)
	helpers.PanicIfError(err)

	clientLog := waLog.Stdout("Client", "DEBUG", true)
	client := whatsmeow.NewClient(deviceStore, clientLog)

	state := &AppState{Client: client}
	client.AddEventHandler(eventHandler)

	// Siapkan QR channel hanya kalau BELUM pernah login
	var qrChan <-chan whatsmeow.QRChannelItem
	if client.Store.ID == nil {
		qrChan, _ = client.GetQRChannel(ctx)
	}

	// Connect ke WA
	helpers.PanicIfError(client.Connect())

	// Jika login via QR (pertama kali), handle event QR di satu goroutine saja
	if qrChan != nil {
		go func() {
			for evt := range qrChan {
				switch evt.Event {
				case "code":
					state.setQR(evt.Code)
					qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
					log.Println("[WA] Scan QR ini dengan WhatsApp (Linked devices)")
				case "success":
					// Login sukses -> sudah terhubung
					state.setQR("")
					state.setConnected(true)
					log.Println("[WA] Login QR sukses ✔️")
				default:
					log.Printf("[WA] Login event: %s", evt.Event)
				}
			}
		}()
	} else {
		// Sudah punya sesi: biasanya langsung connected setelah Connect()
		// Set optimistic true; untuk verifikasi runtime gunakan IsConnected() saat kirim.
		state.setConnected(true)
	}

	// ====== App & Routes (Fiber) ======
	appLog := configs.NewLogger()
	app := configs.NewFiber()
	db := configs.NewDatabase(ctx, appLog)

	configs.Bootstrap(
		&configs.BootstrapConfig{
			Pool: db,
			App:  app,
			Log:  appLog,
		},
	)

	// Status & QR
	app.Get("/qr", func(c *fiber.Ctx) error {
		if state.Client.IsConnected() {
			state.setConnected(true) // sinkronkan flag internal
			return c.JSON(fiber.Map{"status": "connected"})
		}

		qr := state.getQR()
		if qr == "" {
			return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
				"status": "waiting",
				"note":   "QR belum tersedia, refresh sebentar lagi.",
			})
		}
		return c.JSON(fiber.Map{
			"status": "scan_me",
			"qr":     qr,
			"note":   "Scan dengan WhatsApp (Linked devices).",
		})
	})

	// Healthcheck
	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"ok":        true,
			"connected": state.Client.IsConnected(),
		})
	})

	app.Post("/v1/messages", func(c *fiber.Ctx) error {
		if !state.Client.IsConnected() {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error": "whatsapp_not_connected",
			})
		}

		var req struct {
			JID   string `json:"jid"`
			Phone string `json:"phone_number"`
			Text  string `json:"message"`
		}

		if err := c.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid json: "+err.Error())
		}
		if strings.TrimSpace(req.Text) == "" {
			return fiber.NewError(fiber.StatusBadRequest, "`text` wajib diisi")
		}

		phoneNumber := strings.TrimSpace(req.Phone)
		if phoneNumber == "" {
			return fiber.NewError(fiber.StatusBadRequest, "harus menyertakan `jid` atau `phone`")
		}

		jid, err := helpers.NormalizeJID(phoneNumber)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		msg := &waE2E.Message{
			Conversation: protoString(req.Text),
		}

		if _, err := state.Client.SendMessage(ctx, jid, msg); err != nil {
			return fiber.NewError(fiber.StatusBadGateway, "gagal mengirim: "+err.Error())
		}

		return c.JSON(fiber.Map{
			"status":  "pending",
			"to":      jid.String(),
			"message": "Message is pending and waiting to be processed.",
			"id":      "pwa-532cd8656da54257975644d6319",
			"meta": fiber.Map{
				"location": "https://api.pakaiwa.my.id/v1/messages/pwa-532cd8656da54257975644d6319",
			},
		})
	})

	go func() {
		addr := ":8080"
		if err := app.Listen(addr); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Graceful shutdown
	waitForSignal()
	log.Println("Shutting down...")
	_ = app.Shutdown()
	state.Client.Disconnect()
	log.Println("Bye!")
}

/* ================= Helpers ================= */

func waitForSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
}

func protoString(s string) *string { return &s }

func (a *AppState) setQR(code string) {
	a.QRMu.Lock()
	a.LastQR = code
	a.QRMu.Unlock()
}
func (a *AppState) getQR() string {
	a.QRMu.RLock()
	defer a.QRMu.RUnlock()
	return a.LastQR
}
func (a *AppState) setConnected(v bool) {
	a.QRMu.Lock()
	a.Connected = v
	a.QRMu.Unlock()
}
func (a *AppState) getConnected() bool {
	a.QRMu.RLock()
	defer a.QRMu.RUnlock()
	return a.Connected
}
