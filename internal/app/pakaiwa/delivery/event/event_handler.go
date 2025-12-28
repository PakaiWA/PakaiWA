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
	"context"

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
	Ctx            context.Context
}

func (h *HandleEvent) Handle(e any) {
	switch v := e.(type) {
	case *events.Receipt:
		h.DeliveryStatus.ProcessDeliveryStatus(h.Ctx, v)
	case *events.Message:
		h.ReceiveMsgUC.ProcessIncomingMessage(h.Ctx, v.Message, v.Info, v.RawMessage)
	case *events.LoggedOut:
		usecase.HandleLogout(h.PakaiWA.Client)
		h.PakaiWA.SetQR("")
		h.PakaiWA.SetConnected(false)
	}
}
