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
 * @author KAnggara75 on Sat 30/08/25 12.51
 * @project PakaiWA auth
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/auth
 */

package model

import "github.com/golang-jwt/jwt/v5"

type Model struct {
	ID string
}

type JWTClaims struct {
	jwt.RegisteredClaims
	Role          string `json:"role"`
	QuotaLimit    int64  `json:"quota_limit"`
	WindowSeconds int64  `json:"window_seconds"`
}

type AuthUser struct {
	Sub  string
	JTI  string
	Role string
}
