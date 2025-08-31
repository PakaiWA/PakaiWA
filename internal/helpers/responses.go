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
 * @author KAnggara75 on Sun 31/08/25 08.57
 * @project PakaiWA helpers
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/helpers
 */

package helpers

import (
	"github.com/PakaiWA/PakaiWA/internal/app/messages"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func NewMsgID() string {
	prefix := "pwa"
	return prefix + "-" + uuid.NewString()
}

func RespondPending(c *fiber.Ctx) error {
	id := NewMsgID()
	response := messages.MessageResponse{
		ID:      id,
		Status:  "pending",
		Message: "Message is pending and waiting to be processed.",
		MessageMeta: messages.MessageMeta{
			Location: "https://api.pakaiwa.my.id/v1/messages/" + id,
		},
	}

	return c.JSON(response)
}
