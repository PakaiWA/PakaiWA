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
 * @author KAnggara75 on Sat 06/09/25 11.23
 * @project PakaiWA pakaiwa
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa
 */

package pakaiwa

import (
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/handler"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/state"
	"github.com/sirupsen/logrus"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

type EventHandler struct {
	PakaiWA *state.AppState
	Log     *logrus.Logger
}

func NewEventHandler(log *logrus.Logger, state *state.AppState) *EventHandler {
	return &EventHandler{
		PakaiWA: state,
		Log:     log,
	}
}

func (h *EventHandler) Handle(e interface{}) {
	switch v := e.(type) {
	case *events.Message:
		handler.ProcessMessageEvent(v.Message, v.Info, h.Log)
	case *events.LoggedOut:
		reason := v.Reason
		h.PakaiWA.SetQR("")
		h.PakaiWA.SetConnected(false)
		h.Log.Infof("Logged out: %s\n", reason.String())
		if reason >= 400 && reason < 500 {
			handler.HandleLogout(h.PakaiWA.Client)
		}
	case *events.Receipt:
		switch v.Type {
		case types.ReceiptTypeDelivered:
			h.Log.Infof("Pesan %v delivered ke %s\n", v.MessageIDs, v.Sender)
		case types.ReceiptTypeRead:
			h.Log.Infof("Pesan %v dibaca oleh %s\n", v.MessageIDs, v.Sender)
		case types.ReceiptTypePlayed:
			h.Log.Infof("Voice note %v diputar oleh %s\n", v.MessageIDs, v.Sender)
		}
	}
}
