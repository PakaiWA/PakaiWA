/*
 * Copyright (c) 2025 KAnggara
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * See <https://www.gnu.org/licenses/gpl-3.0.html>.
 *
 * @author KAnggara on Saturday 13/12/2025 12.20
 * @project PakaiWA
 * ~/work/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/model
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/delivery/model
 */

package model

import (
	"time"
)

type User struct {
	ID        string
	Email     string
	Role      string
	Password  string
	LogoutAt  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
