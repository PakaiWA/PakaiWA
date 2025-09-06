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
 * @author KAnggara75 on Sat 30/08/25 12.48
 * @project PakaiWA usecase
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/usecase
 */

package usecase

import (
	"github.com/sirupsen/logrus"
)

type UserUseCase struct {
	Log *logrus.Logger
}

func NewUserUseCase(logger *logrus.Logger) *UserUseCase {
	return &UserUseCase{
		Log: logger,
	}
}
