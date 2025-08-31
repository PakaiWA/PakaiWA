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
 * @author KAnggara75 on Sun 31/08/25 12.04
 * @project PakaiWA qr
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/handler
 */

package handler

import (
	"github.com/PakaiWA/PakaiWA/internal/model"
	"github.com/PakaiWA/PakaiWA/internal/pakaiwa"
	"github.com/gofiber/fiber/v2"
	"github.com/mdp/qrterminal/v3"
	"github.com/sirupsen/logrus"
	"go.mau.fi/whatsmeow"
	"log"
	"os"
)

type HandlerQR struct {
	State *pakaiwa.AppState
	Log   *logrus.Logger
}

func NewQRHandler(state *pakaiwa.AppState, log *logrus.Logger) *HandlerQR {
	return &HandlerQR{
		Log:   log,
		State: state,
	}
}

func (h *HandlerQR) GetQR(c *fiber.Ctx) error {
	qrResponse := &model.ResponseQR{
		QRCode:  "",
		QRImage: "",
	}

	if h.State.Client.IsConnected() {
		h.State.SetConnected(true)
		return c.JSON(qrResponse)
	}

	qrData := h.State.GetQR()
	if qrData == "" {
		qrResponse.Msg = "No QR Code"
		return c.JSON(qrResponse)
	}

	qrResponse.QRImage = "https://api.pakaiwa.my.id/v1/qr/show?qrcode=" + qrData
	qrResponse.QRCode = qrData

	return c.JSON(qrResponse)
}

func StartQRHandler(state *pakaiwa.AppState, qrChan <-chan whatsmeow.QRChannelItem) {
	if qrChan == nil {
		state.SetConnected(true)
		return
	}

	go func() {
		for evt := range qrChan {
			switch evt.Event {
			case "code":
				state.SetQR(evt.Code)
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
				log.Println("[WA] Scan QR ini dengan WhatsApp (Linked devices)")
			case "success":
				state.SetQR("")
				state.SetConnected(true)
				log.Println("[WA] Login QR sukses ✔️")
			default:
				log.Printf("[WA] Login event: %s", evt.Event)
			}
		}
	}()
}
