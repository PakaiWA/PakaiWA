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
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v3"

	"github.com/PakaiWA/whatsmeow/types"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/model"
)

func RespondPending(c fiber.Ctx, msg, msgId string) error {
	id := "pwa-" + strings.ToLower(msgId)
	location := fmt.Sprintf("%s/v1/messages/%s", c.BaseURL(), id)
	response := model.SendMessageResponse{
		ID:      id,
		Status:  "pending",
		Message: msg,
		MessageMeta: model.MessageMeta{
			Location: location,
		},
	}
	return c.JSON(response)
}

func ResponseGroupList(c fiber.Ctx, groups []*types.GroupInfo) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Groups retrieved successfully",
		"data":    groups,
	})
}
