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
 * @author KAnggara75 on Sun 07/09/25 00.09
 * @project PakaiWA event
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/delivery/event
 */

package event

import (
	"github.com/PakaiWA/whatsmeow/types/events"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/state"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/usecase"
)

type HandleEvent struct {
	PakaiWA        *state.AppState
	Producer       *kafka.Producer
	ReceiveMsgUC   usecase.ReceiveMessageUsecase
	DeliveryStatus usecase.DeliveryUsecase
	Log            *logrus.Logger
}

func (h *HandleEvent) Handle(e interface{}) {
	switch v := e.(type) {
	case *events.Receipt:
		h.DeliveryStatus.ProcessDeliveryStatus(v)
	case *events.Message:
		h.ReceiveMsgUC.ProcessIncomingMessage(v.Message, v.Info, v.RawMessage)
	case *events.LoggedOut:
		reason := v.Reason
		h.PakaiWA.SetQR("")
		h.PakaiWA.SetConnected(false)
		h.Log.Infof("Logged out: %s\n", reason.String())
		if reason >= 400 && reason < 500 {
			usecase.HandleLogout(h.PakaiWA.Client)
		}

	}
}
