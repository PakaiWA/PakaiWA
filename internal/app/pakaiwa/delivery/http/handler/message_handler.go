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
 * @author KAnggara75 on Sun 07/09/25 08.00
 * @project PakaiWA handler
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/delivery/http/handler
 */

package handler

import (
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/model"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/helper"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type MessageHandler struct {
	UseCase usecase.MessageUsecase
	Log     *logrus.Logger
}

func NewMessageHandler(useCase usecase.MessageUsecase, log *logrus.Logger) *MessageHandler {
	return &MessageHandler{
		Log:     log,
		UseCase: useCase,
	}
}

func (h *MessageHandler) SendMsg(c *fiber.Ctx) error {
	request := new(model.SendMessageReq)
	if err := c.BodyParser(request); err != nil {
		h.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	id, err := h.UseCase.SendMessage(request)
	if err != nil {
		h.Log.WithError(err).Error("error creating contact")
		return err
	}

	return helper.RespondPending(c, id)
}
