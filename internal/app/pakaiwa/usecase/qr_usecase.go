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
 * @author KAnggara75 on Sun 31/08/25 20.50
 * @project PakaiWA qr
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/usecase
 */

package usecase

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type UseCaseQR struct {
	Log      *logrus.Logger
	Validate *validator.Validate
}

func NewQRUsecase(logger *logrus.Logger, validate *validator.Validate) *UseCaseQR {
	return &UseCaseQR{
		Log:      logger,
		Validate: validate,
	}
}
