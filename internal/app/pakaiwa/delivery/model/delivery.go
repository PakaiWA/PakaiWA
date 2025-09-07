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
 * @author KAnggara75 on Sun 07/09/25 18.05
 * @project PakaiWA model
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/delivery/model
 */

package model

import (
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/entity"
	"time"
)

type DeliveryModel struct {
	Id          string                `json:"id"`
	WebhookType string                `json:"webhook_type"`
	Status      entity.DeliveryStatus `json:"status"`
	Message     string                `json:"message"`
	Payload     DeliveryPayload       `json:"payload"`
	CreatedAt   time.Time             `json:"created_at"`
	ServerTime  time.Time             `json:"server_time"`
}

func (a *DeliveryModel) GetId() string {
	return a.Id
}

type DeliveryPayload struct {
	Message     string `json:"message"`
	PhoneNumber string `json:"phone_number"`
	DeviceId    string `json:"device_id"`
	MessageType string `json:"message_type"`
}

func ToDeliveryModel(ent entity.WebhookEntity) *DeliveryModel {
	var payload DeliveryPayload
	if p, ok := ent.Payload.(entity.DeliveryStatusPayload); ok {
		payload = DeliveryPayload{
			Message:     p.Message,
			PhoneNumber: p.PhoneNumber,
			DeviceId:    p.DeviceId,
			MessageType: p.MessageType,
		}
	}

	return &DeliveryModel{
		Id:          ent.ID,
		WebhookType: ent.WebhookType,
		Status:      ent.Status,
		Message:     "Message has been processed",
		Payload:     payload,
		CreatedAt:   ent.CreatedTime,
		ServerTime:  ent.ServerTime,
	}

}
