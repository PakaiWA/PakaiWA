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
	"go.mau.fi/whatsmeow/types"
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

	// Jika belum pernah login, siapkan QR channel
	var qrChan <-chan whatsmeow.QRChannelItem
	if client.Store.ID == nil {
		qrChan, _ = client.GetQRChannel(ctx)
		err = client.Connect()
		helpers.PanicIfError(err)

		for evt := range qrChan {
			if evt.Event == "code" {
				state.setQR(evt.Code)
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		err = client.Connect()
		helpers.PanicIfError(err)
	}

	if qrChan != nil {
		go func() {
			for evt := range qrChan {
				switch evt.Event {
				case "code":
					state.setQR(evt.Code)
					log.Printf("Scan QR (GET /qr untuk melihat): %s", evt.Code)
				default:
					log.Printf("Login event: %s", evt.Event)
				}
			}
		}()
	} else {
		state.setConnected(true)
	}

	log := configs.NewLogger()
	app := configs.NewFiber()

	db := configs.NewDatabase(ctx, log)

	configs.Bootstrap(
		&configs.BootstrapConfig{
			Pool: db,
			App:  app,
			Log:  log,
		},
	)

	app.Get("/qr", func(c *fiber.Ctx) error {
		if state.getConnected() {
			return c.JSON(fiber.Map{
				"status": "connected",
			})
		}
		qr := state.getQR()
		if qr == "" {
			return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
				"status": "waiting",
				"note":   "QR belum tersedia, tunggu sebentar lalu refresh.",
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
			"connected": state.getConnected(),
		})
	})
	// POST /send -> kirim pesan teks
	// Body: { "jid": "6281234567890", "text": "halo" }
	//        atau { "phone": "6281234567890", "text": "halo" }
	app.Post("/send", func(c *fiber.Ctx) error {
		if !state.getConnected() {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error": "whatsapp_not_connected",
			})
		}

		var req struct {
			JID   string `json:"jid"`
			Phone string `json:"phone"`
			Text  string `json:"text"`
		}
		if err := c.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid json: "+err.Error())
		}
		if strings.TrimSpace(req.Text) == "" {
			return fiber.NewError(fiber.StatusBadRequest, "`text` wajib diisi")
		}

		jidStr := strings.TrimSpace(req.JID)
		if jidStr == "" {
			jidStr = strings.TrimSpace(req.Phone)
		}
		if jidStr == "" {
			return fiber.NewError(fiber.StatusBadRequest, "harus menyertakan `jid` atau `phone`")
		}

		jid, err := normalizeJID(jidStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "jid/phone tidak valid: "+err.Error())
		}

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		msg := &waE2E.Message{
			Conversation: protoString(req.Text),
		}

		_, err = state.Client.SendMessage(ctx, jid, msg)
		if err != nil {
			return fiber.NewError(fiber.StatusBadGateway, "gagal mengirim: "+err.Error())
		}
		return c.JSON(fiber.Map{
			"status": "sent",
			"to":     jid.String(),
			"text":   req.Text,
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
	//client.Disconnect()
	log.Println("Bye!")
}

func waitForSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
}

func protoString(s string) *string {
	return &s
}

func (a *AppState) setQR(code string) {
	a.QRMu.Lock()
	defer a.QRMu.Unlock()
	a.LastQR = code
}

func (a *AppState) getQR() string {
	a.QRMu.RLock()
	defer a.QRMu.RUnlock()
	return a.LastQR
}

func (a *AppState) setConnected(v bool) {
	a.QRMu.Lock()
	defer a.QRMu.Unlock()
	a.Connected = v
}

func (a *AppState) getConnected() bool {
	a.QRMu.RLock()
	defer a.QRMu.RUnlock()
	return a.Connected
}

func normalizeJID(s string) (types.JID, error) {
	s = strings.TrimSpace(s)
	// Sudah lengkap?
	if strings.Contains(s, "@") {
		j, err := types.ParseJID(s)
		if err != nil {
			return types.JID{}, err
		}
		return j, nil
	}
	// Angka saja -> asumsikan personal chat
	if isAllDigits(s) {
		return types.JID{User: s, Server: types.DefaultUserServer}, nil // s.whatsapp.net
	}
	return types.JID{}, fmt.Errorf("format tidak dikenali")
}

func isAllDigits(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return s != ""
}
