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

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/apperror"
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

func (m *messageUsecase) SendMessage(ctx context.Context, req *model.SendMessageReq) (string, error) {
	log := ctxmeta.Logger(ctx)
	if !m.WA.IsConnected() {
		log.WithError(apperror.ErrWAClientNotConnected).WithField("event", "wa_disconnected").Error("precondition failed")
		return "", apperror.ErrWAClientNotConnected
	}

	if err := m.Validate.Struct(req); err != nil {
		if log != nil {
			log.WithError(err).WithField("event", "validation_failed").Warn("invalid send message request")
		}
		return "", err
	}

	id := m.WA.GenerateMessageID()
	if strings.TrimSpace(req.Text) == "" {
		return id, apperror.ErrInvalidMessage
	}

	phoneNumber := strings.TrimSpace(req.Phone)
	jid, err := helper.NormalizeJID(phoneNumber)
	if err != nil {
		return id, apperror.ErrInvalidMessage
	}

	go func(parent context.Context, msgID string) {
		ctxSend, cancel := context.WithTimeout(parent, 15*time.Second)
		defer cancel()

		msg := &waE2E.Message{Conversation: utils.ProtoString(req.Text)}

		_, err := m.WA.SendMessage(ctxSend, jid, msg, whatsmeow.SendRequestExtra{ID: msgID})

		if err != nil {
			if l := ctxmeta.Logger(ctxSend); l != nil {
				l.
					WithError(err).
					WithField("event", "send_message_failed").
					WithField("message_id", msgID).
					Error("failed to send whatsapp message")
			}
		}
	}(ctx, id)

	return id, nil
}

func (m *messageUsecase) EditMessage(ctx context.Context, req *model.SendMessageReq, msgId string) error {
	phoneNumber := strings.TrimSpace(req.Phone)
	jid, err := helper.NormalizeJID(phoneNumber)
	if err != nil {
		return apperror.ErrInvalidMessage
	}

	id := strings.ToUpper(msgId)
	if strings.TrimSpace(req.Text) == "" {
		return apperror.ErrInvalidMessage
	}

	m.WA.Log.Infof(req.Text)

	go func(parent context.Context, msgID string) {
		ctxSend, cancel := context.WithTimeout(parent, 15*time.Second)
		defer cancel()

		_, err := m.WA.SendMessage(ctxSend, jid, m.WA.BuildEdit(jid, msgID, &waE2E.Message{
			Conversation: utils.ProtoString(req.Text),
		}))

		if err != nil {
			if l := ctxmeta.Logger(ctxSend); l != nil {
				l.
					WithError(err).
					WithField("event", "send_message_failed").
					WithField("message_id", msgID).
					Error("failed to send whatsapp message")
			}
		}
	}(ctx, id)

	return nil
}
