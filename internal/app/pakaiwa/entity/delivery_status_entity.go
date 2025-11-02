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
 * @author KAnggara75 on Sun 07/09/25 16.56
 * @project PakaiWA entity
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/entity
 */

package entity

type DeliveryStatus string

const (
	DeliverySent      DeliveryStatus = "sent"
	DeliveryDelivered DeliveryStatus = "delivered"
	DeliveryRead      DeliveryStatus = "read"
	DeliveryFailed    DeliveryStatus = "failed"
)

type DeliveryStatusPayload struct {
	ID          string
	Status      DeliveryStatus
	Message     string
	PhoneNumber string
	DeviceId    string
	MessageType string
}

func (p DeliveryStatusPayload) GetType() string {
	return "delivery_status"
}
