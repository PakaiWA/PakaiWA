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
 * @author KAnggara75 on Sun 31/08/25 12.06
 * @project PakaiWA pakaiwa
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/pakaiwa
 */

package handler

import (
	"github.com/PakaiWA/PakaiWA/internal/pakaiwa"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sirupsen/logrus"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

type EventHandler struct {
	PakaiWA *pakaiwa.AppState
	Log     *logrus.Logger
}

func NewEventHandler(log *logrus.Logger, state *pakaiwa.AppState) *EventHandler {
	return &EventHandler{
		PakaiWA: state,
		Log:     log,
	}
}

func (h *EventHandler) Handle(e interface{}) {
	switch v := e.(type) {
	case *events.Message:
		ProcessMessageEvent(v.Message, v.Info, h.Log)
	case *events.LoggedOut:
		reason := v.Reason
		h.PakaiWA.SetQR("")
		h.PakaiWA.SetConnected(false)
		log.Infof("Logged out: %s\n", reason.String())
		if reason >= 400 && reason < 500 {
			HandleLogout(h.PakaiWA.Client)
		}
	case *events.Receipt:
		switch v.Type {
		case types.ReceiptTypeDelivered:
			log.Infof("Pesan %v delivered ke %s\n", v.MessageIDs, v.Sender)
		case types.ReceiptTypeRead:
			log.Infof("Pesan %v dibaca oleh %s\n", v.MessageIDs, v.Sender)
		case types.ReceiptTypePlayed:
			log.Infof("Voice note %v diputar oleh %s\n", v.MessageIDs, v.Sender)
		}
	}
}
