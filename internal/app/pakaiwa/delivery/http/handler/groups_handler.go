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
 * @author KAnggara on Sunday 28/12/2025 23.20
 * @project PakaiWA
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/delivery/http/handler
 */

package handler

import "github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/usecase"

type GroupHandler struct {
	UseCase usecase.GroupUsecase
}

func NewGroupHandler(useCase usecase.GroupUsecase) *GroupHandler {
	return &GroupHandler{
		UseCase: useCase,
	}
}
