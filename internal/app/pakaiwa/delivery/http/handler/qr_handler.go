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
 * @author KAnggara75 on Sun 07/09/25 08.02
 * @project PakaiWA handler
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/delivery/http/handler
 */

package handler

import (
	"github.com/gofiber/fiber/v3"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/state"
)

type QRHandler struct {
	State *state.AppState
}

func NewQRHandler(state *state.AppState) *QRHandler {
	return &QRHandler{
		State: state,
	}
}

func (h *QRHandler) GetQR(c fiber.Ctx) error {
	//		log := h.Log
	//		qrResponse := &model.ResponseQR{}
	//
	//		if h.State.GetConnected() {
	//			log.Info("client connected")
	//			qrResponse.Msg = "Already Connected"
	//			qrResponse.Status = "connected"
	//			qrResponse.Device.JID = helpers.NormalizeNumber(h.State.Client.Store.ID.String())
	//			qrResponse.Device.PushName = h.State.Client.Store.PushName
	//			return c.Status(fiber.StatusAccepted).JSON(qrResponse)
	//		}
	//
	//		log.Info("client not connected yet")
	//		qrData := h.State.GetQR()
	//		if qrData == "" {
	//			qrResponse.Msg = "No QR Code"
	//			return c.Status(fiber.StatusServiceUnavailable).JSON(qrResponse)
	//		} else {
	//			qrData = url.QueryEscape(qrData)
	//			qrResponse.QRImage = c.BaseURL() + "/v1/qr/show?qrcode=" + qrData
	//			qrResponse.QRCode = qrData
	//			return c.JSON(qrResponse)
	//		}
	//	}
	//
	//	func (h *QRHandler) ShowQR(c fiber.Ctx) error {
	//		log := h.Log
	//		qrData := c.Query("qrcode", c.Query("qrCode", ""))
	//		log.Infof("show qrcode: %s", qrData)
	//		if qrData == "" {
	//			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//				"error": "qrcode query parameter is required",
	//			})
	//		}
	//
	//		c.Set("Content-Type", "image/png")
	//		png, err := qrcode.Encode(qrData, qrcode.Highest, 512)
	//
	//		if err != nil {
	//			log.Error("failed to generate QR code image:", err)
	//			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//				"error": "failed to generate QR code image",
	//			})
	//		}
	//
	//		return c.Send(png)
	return c.JSON(fiber.Map{
		"message": "Not Implemented",
	})
}
