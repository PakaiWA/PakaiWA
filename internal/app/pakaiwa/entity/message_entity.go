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
 * @author KAnggara75 on Sat 06/09/25 14.13
 * @project PakaiWA entity
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/entity
 */

package entity

import "time"

type MessageEntity struct {
	Id       string
	Status   string
	PhoneNo  string
	Message  string
	DeviceId string
	Type     string
	Caption  string
	IsGroup  bool
	SendAt   time.Time
}

type BatchMessageEntity struct {
	DeviceId string
	Messages []MessageEntity
}
