/*
 * Copyright (c) 2026 KAnggara
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * See <https://www.gnu.org/licenses/gpl-3.0.html>.
 *
 * @author KAnggara on Saturday 03/01/2026 01.36
 * @project PakaiWA
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/helper
 */

package helper

import (
	"context"

	"github.com/PakaiWA/PakaiWA/internal/pkg/logger/ctxmeta"
)

func IsMessageError(err error, ctxSend context.Context, msgID string) bool {
	if err != nil {
		if l := ctxmeta.Logger(ctxSend); l != nil {
			l.
				WithError(err).
				WithField("event", "send_message_failed").
				WithField("message_id", msgID).
				Error("failed to send whatsapp message")
		}
		return true
	}
	return false
}
