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
 * @author KAnggara on Saturday 03/01/2026 00.34
 * @project PakaiWA
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/delivery/model
 */

package model

type GroupModel struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
