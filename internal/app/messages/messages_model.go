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
 * @author KAnggara75 on Sun 31/08/25 09.10
 * @project PakaiWA messages
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/messages
 */

package messages

type MessageResponse struct {
	ID          string      `json:"id"`
	Status      string      `json:"status"`
	To          string      `json:"to"`
	Message     string      `json:"message"`
	MessageMeta MessageMeta `json:"meta"`
}

type MessageMeta struct {
	Location string `json:"location"`
}
