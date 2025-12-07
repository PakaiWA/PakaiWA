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
 * @author KAnggara75 on Tue 18/11/25 06.04
 * @project PakaiWA usecase
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/usecase
 */

package usecase

import (
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/model"
)

type AuthUsecase interface {
	Login(req *model.LoginReq) (string, error)
	Register(req *model.AuthReq) (bool, error)
}
