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
	"github.com/sirupsen/logrus"
	"github.com/skip2/go-qrcode"
	"net/url"
)

type QRHandler struct {
	State *pakaiwa.AppState
	Log   *logrus.Logger
}

func NewQRHandler(state *pakaiwa.AppState, log *logrus.Logger) *QRHandler {
	return &QRHandler{
		Log:   log,
		State: state,
	}
}

func (h *QRHandler) GetQR(c *fiber.Ctx) error {
	qrResponse := &model.ResponseQR{}

	if h.State.GetConnected() {
		h.Log.Info("client connected")
		return c.JSON(qrResponse)
	}

	h.Log.Info("client not connected yet")
	qrData := h.State.GetQR()
	if qrData == "" {
		qrResponse.Msg = "No QR Code"
		return c.JSON(qrResponse)
	}
	qrData = url.QueryEscape(qrData)

	qrResponse.QRImage = c.BaseURL() + "/v1/qr/show?qrcode=" + qrData
	qrResponse.QRCode = qrData
	return c.JSON(qrResponse)
}

func (h *QRHandler) ShowQR(c *fiber.Ctx) error {
	qrData := c.Query("qrcode", "")
	h.Log.Infof("show qrcode: %s", qrData)
	if qrData == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "qrcode query parameter is required",
		})
	}

	c.Set("Content-Type", "image/png")
	png, err := qrcode.Encode(qrData, qrcode.Highest, 512)

	if err != nil {
		h.Log.Error("failed to generate QR code image:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to generate QR code image",
		})
	}

	return c.Send(png)
}
