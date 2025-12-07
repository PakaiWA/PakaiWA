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
 * @author KAnggara75 on Tue 18/11/25 06.03
 * @project PakaiWA handler
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/delivery/http/handler
 */

package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/model"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/usecase"
	"github.com/PakaiWA/PakaiWA/internal/pkg/utils"
)

type AuthHandler struct {
	UseCase usecase.AuthUsecase
	Log     *logrus.Logger
}

func NewAuthHandler(uc usecase.AuthUsecase, log *logrus.Logger) *AuthHandler {
	return &AuthHandler{
		Log:     log,
		UseCase: uc,
	}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	request := new(model.LoginReq)

	if err := c.Bind().Body(request); err != nil {
		utils.LogValidationErrors(h.Log, err, "error parsing request body")
		return fiber.ErrBadRequest
	}

	token, err := h.UseCase.Login(request)
	if err != nil {
		utils.LogValidationErrors(h.Log, err, "validation failed in Login", c.Path())
		return err
	}

	response := model.JwtResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   3600,
	}

	return c.JSON(response)
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	request := new(model.AuthReq)
	if err := c.Bind().Body(request); err != nil {
		utils.LogValidationErrors(h.Log, err, "error parsing request body", c.Path())
		return fiber.ErrBadRequest
	}

	isSuccess, err := h.UseCase.Register(request)
	if err != nil {
		utils.LogValidationErrors(h.Log, err, "validation failed in Register", c.Path())
		return err
	}

	if isSuccess {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"status":  "success",
			"message": "user registered",
		})
	} else {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  "failed",
			"message": "user already exists",
		})
	}
}
