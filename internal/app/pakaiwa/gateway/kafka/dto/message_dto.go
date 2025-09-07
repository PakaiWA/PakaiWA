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
 * @author KAnggara75 on Sat 06/09/25 17.25
 * @project PakaiWA dto
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/gateway/kafka/dto
 */

package dto

type IncomingMessageEvent struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (c *IncomingMessageEvent) GetId() string {
	return c.ID
}
