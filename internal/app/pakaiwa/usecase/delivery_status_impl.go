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
 * @author KAnggara75 on Sun 07/09/25 17.41
 * @project PakaiWA usecase
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/usecase
 */

package usecase

import (
	"strings"
	"time"

	"github.com/PakaiWA/whatsmeow/types"
	"github.com/PakaiWA/whatsmeow/types/events"
	"github.com/sirupsen/logrus"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/model"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/entity"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/gateway/kafka"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/helper"
)

type deliveryStatusUsecase struct {
	Log      *logrus.Logger
	Producer *kafka.DeliveryStatusProducer
}

func NewDeliveryStatusUsecase(log *logrus.Logger, producer *kafka.DeliveryStatusProducer) DeliveryUsecase {
	return &deliveryStatusUsecase{
		Log:      log,
		Producer: producer,
	}
}

func (d *deliveryStatusUsecase) ProcessDeliveryStatus(event *events.Receipt) {
	status := entity.DeliveryFailed

	switch event.Type {
	case types.ReceiptTypeDelivered:
		status = entity.DeliveryDelivered
	case types.ReceiptTypeSender:
		status = entity.DeliverySent
	case types.ReceiptTypeRead:
		status = entity.DeliveryRead
	}

	for _, msgID := range event.MessageIDs {
		msgPayload := entity.DeliveryStatusPayload{
			Message:     "",
			PhoneNumber: helper.NormalizeNumber(event.Sender.String()),
			MessageType: "",
			DeviceId:    "",
		}

		deliveryPayload := entity.WebhookEntity{
			ID:          strings.ToLower(msgID),
			Status:      status,
			WebhookType: msgPayload.GetType(),
			Payload:     msgPayload,
			CreatedTime: event.Timestamp,
			ServerTime:  time.Now(),
		}

		deliveryModel := model.ToDeliveryModel(deliveryPayload)
		err := d.Producer.Send(deliveryModel)
		if err != nil {
			d.Log.Error("failed to send delivery status to kafka: ", err)
		}
	}
}
