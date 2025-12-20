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
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/pkg/errors"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/apperror"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/http/dto"
	"github.com/PakaiWA/PakaiWA/internal/pkg/utils"
)

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx fiber.Ctx, err error) error {

		// Default -> 500
		code := fiber.StatusInternalServerError

		// Jika error bawaan fiber
		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		if errors.Is(err, apperror.ErrInvalidCredentials) {
			return ctx.Status(fiber.StatusUnauthorized).JSON(dto.BaseResponse{
				Error: &dto.ProblemDetails{
					Type:     "https://api.pakaiwa.my.id/problems/auth/invalid-credentials",
					Title:    "Unauthorized",
					Status:   fiber.StatusUnauthorized,
					Detail:   err.Error(),
					Instance: ctx.Path(),
				},
			})
		}

		var validationError validator.ValidationErrors
		if errors.As(err, &validationError) {
			var fieldErrors []dto.ValidationError
			for _, v := range validationError {
				fieldErrors = append(fieldErrors, dto.ValidationError{
					Field: v.Field(),
					Tag:   v.Tag(),
					Param: v.Param(),
				})
			}

			return ctx.Status(fiber.StatusBadRequest).JSON(dto.BaseResponse{
				Error: &dto.ProblemDetails{
					Type:     "https://api.pakaiwa.my.id/problems/validation-error",
					Title:    "Validation Error",
					Status:   fiber.StatusBadRequest,
					Detail:   "request validation failed",
					Instance: ctx.Path(),
					Errors:   fieldErrors, // ⬅️ di SINI tempatnya
				},
			})
		}

		var ve *utils.PasswordValidationError
		if errors.As(err, &ve) {
			return ctx.Status(fiber.StatusBadRequest).JSON(dto.BaseResponse{
				Error: &dto.ProblemDetails{
					Type:     "https://api.pakaiwa.my.id/problems/auth/validation-error",
					Title:    "Validation Error",
					Status:   fiber.StatusBadRequest,
					Detail:   err.Error(),
					Instance: ctx.Path(),
					Errors:   ve.Errors,
				},
			})
		}

		if errors.Is(err, apperror.ErrUsernameExists) {
			return ctx.Status(fiber.StatusBadRequest).JSON(dto.BaseResponse{
				Error: &dto.ProblemDetails{
					Type:     "https://api.pakaiwa.my.id/problems/auth/validation-error",
					Title:    "Validation Error",
					Status:   fiber.StatusBadRequest,
					Detail:   err.Error(),
					Instance: ctx.Path(),
				},
			})
		}

		// 4. Error umum -> 500 atau kode fiber
		return ctx.Status(code).JSON(dto.BaseResponse{
			Error: &dto.ProblemDetails{
				Title:    http.StatusText(code),
				Status:   code,
				Detail:   err.Error(),
				Instance: ctx.Path(),
			},
		})
	}
}
