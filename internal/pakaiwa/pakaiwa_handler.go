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
 * @author KAnggara75 on Sun 31/08/25 12.06
 * @project PakaiWA pakaiwa
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/pakaiwa
 */

package pakaiwa

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"go.mau.fi/whatsmeow/types/events"
)

func EventHandler(e interface{}) {
	switch v := e.(type) {
	case *events.Message:
		msg := v.Message
		fmt.Println("Received a message!", v.Message.GetConversation())
		if msg.GetConversation() != "" {
			log.Debugf("Received a message! %s from %s", msg.GetConversation(), v.Info.Sender.String())
		}
	}
}
