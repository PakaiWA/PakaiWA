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
 * @author KAnggara75 on Sun 07/09/25 16.47
 * @project PakaiWA entity
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/entity
 */

package entity

import "time"

type WebhookEntity struct {
	ID          string
	Status      DeliveryStatus
	WebhookType string
	Payload     WebhookPayload
	CreatedTime time.Time
	ServerTime  time.Time
}

type WebhookPayload interface {
	GetType() string
}
