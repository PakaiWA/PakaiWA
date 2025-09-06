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
 * @author KAnggara75 on Sat 06/09/25 11.39
 * @project PakaiWA helper
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/helper
 */

package helper

import (
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/model"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func RespondPending(c *fiber.Ctx, msgId string) error {
	id := "pwa-" + strings.ToLower(msgId)
	response := model.SendMessageResponse{
		ID:      id,
		Status:  "pending",
		Message: "Message is pending and waiting to be processed.",
		MessageMeta: model.MessageMeta{
			Location: "https://api.pakaiwa.my.id/v1/messages/" + id,
		},
	}
	return c.JSON(response)
}
