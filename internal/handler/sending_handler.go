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
 * @author KAnggara75 on Sun 31/08/25 18.41
 * @project PakaiWA webhooks
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/webhooks
 */

package handler

import (
	"context"
	"github.com/PakaiWA/PakaiWA/internal/helpers"
	"github.com/PakaiWA/PakaiWA/internal/model"
	"github.com/PakaiWA/PakaiWA/internal/pakaiwa"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"strings"
	"time"
)

type MessageHandler struct {
	State *pakaiwa.AppState
	Log   *logrus.Logger
}

func NewMessageHandler(state *pakaiwa.AppState, log *logrus.Logger) *MessageHandler {
	return &MessageHandler{
		Log:   log,
		State: state,
	}
}

func (mh *MessageHandler) SendMsg(c *fiber.Ctx) error {
	if !mh.State.Client.IsConnected() {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "whatsapp_not_connected",
		})
	}

	req := model.SendMessageReq{}

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
		Conversation: helpers.ProtoString(req.Text),
	}

	response, err := mh.State.Client.SendMessage(ctx, jid, msg)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, "sending Fail: "+err.Error())
	}
	return helpers.RespondPending(c, response.ID)
}
