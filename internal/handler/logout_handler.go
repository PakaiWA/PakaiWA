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
 * @author KAnggara75 on Mon 01/09/25 18.05
 * @project PakaiWA handler
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/handler
 */

package handler

import (
	"context"
	"go.mau.fi/whatsmeow"
)

func HandleLogout(client *whatsmeow.Client) {
	client.Disconnect()
	_ = client.Store.Delete(context.Background())
}
