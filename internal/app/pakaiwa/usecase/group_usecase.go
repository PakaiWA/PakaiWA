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
 * @author KAnggara on Sunday 28/12/2025 23.21
 * @project PakaiWA
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/usecase
 */

package usecase

import (
	"github.com/PakaiWA/whatsmeow"
	"github.com/go-playground/validator/v10"
)

type GroupUsecase interface {
}

type groupUsecase struct {
	Validate *validator.Validate
	WA       *whatsmeow.Client
}

func NewGroupUsecase(validate *validator.Validate, wa *whatsmeow.Client) GroupUsecase {
	return &groupUsecase{
		WA:       wa,
		Validate: validate,
	}
}
