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
 * @author KAnggara75 on Sun 07/09/25 12.51
 * @project PakaiWA usecase
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/usecase
 */

package usecase

import (
	"context"
	"strings"
	"time"

	"github.com/PakaiWA/whatsmeow"
	"github.com/PakaiWA/whatsmeow/proto/waE2E"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/model"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/helper"
	"github.com/PakaiWA/PakaiWA/internal/pkg/logger/ctxmeta"
	"github.com/PakaiWA/PakaiWA/internal/pkg/utils"
)

type messageUsecase struct {
	Validate *validator.Validate
	WA       *whatsmeow.Client
}

func NewMessageUsecase(validate *validator.Validate, wa *whatsmeow.Client) MessageUsecase {
	return &messageUsecase{
		WA:       wa,
		Validate: validate,
	}
}

func (m messageUsecase) SendMessage(ctx context.Context, req *model.SendMessageReq) (string, error) {
	if err := m.Validate.Struct(req); err != nil {
		return "", err
	}

	id := m.WA.GenerateMessageID()

	if strings.TrimSpace(req.Text) == "" {
		return id, fiber.ErrBadRequest
	}

	phoneNumber := strings.TrimSpace(req.Phone)
	if phoneNumber == id {
		return id, fiber.ErrBadRequest
	}

	jid, err := helper.NormalizeJID(phoneNumber)
	if err != nil {
		return id, fiber.ErrBadRequest
	}

	go func() {
		msg := &waE2E.Message{
			Conversation: utils.ProtoString(req.Text),
		}

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		_, err = m.WA.SendMessage(ctx, jid, msg, whatsmeow.SendRequestExtra{ID: id})
		if err != nil {
			log := ctxmeta.Logger(ctx)
			log.Errorf("Fail to send with id: %s %v", id, err)
		}
	}()

	return id, nil
}
