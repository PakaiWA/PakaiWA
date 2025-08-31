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
	"github.com/PakaiWA/PakaiWA/internal/app/webhooks"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

func EventHandler(e interface{}) {
	switch v := e.(type) {
	case *events.Message:
		webhooks.ProcessMessageEvent(v.Message, v.Info)
	case *events.Receipt:
		switch v.Type {
		case types.ReceiptTypeDelivered:
			fmt.Printf("Pesan %v delivered ke %s\n", v.MessageIDs, v.Sender)
		case types.ReceiptTypeRead:
			fmt.Printf("Pesan %v dibaca oleh %s\n", v.MessageIDs, v.Sender)
		case types.ReceiptTypePlayed:
			fmt.Printf("Voice note %v diputar oleh %s\n", v.MessageIDs, v.Sender)
		}
	}
}
