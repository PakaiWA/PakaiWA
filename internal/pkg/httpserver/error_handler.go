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
 * @author KAnggara75 on Sat 06/09/25 11.02
 * @project PakaiWA httpserver
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/pkg/httpserver
 */

package httpserver

import (
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/http/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func NewErrorHandler() fiber.ErrorHandler {

	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		return ctx.Status(code).JSON(dto.BaseResponse{
			Success: false,
			Error: &dto.ProblemDetails{
				Type:     "about:blank",
				Title:    "Internal Server Error",
				Status:   code,
				Detail:   err.Error(),
				Instance: ctx.Path(),
			},
		})
	}
}
