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
	"errors"
	"strings"

	"github.com/gofiber/fiber/v3"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/model"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/helper"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/usecase"
	"github.com/PakaiWA/PakaiWA/internal/pkg/utils"
)

type MessageHandler struct {
	UseCase usecase.MessageUsecase
}

func NewMessageHandler(useCase usecase.MessageUsecase) *MessageHandler {
	return &MessageHandler{
		UseCase: useCase,
	}
}

func (h *MessageHandler) SendMsg(c fiber.Ctx) error {
	request := new(model.SendMessageReq)
	if err := c.Bind().Body(request); err != nil {
		utils.LogValidationErrors(c.Context(), err, "error parsing request body", c.Path())
		return fiber.ErrBadRequest
	}

	id, err := h.UseCase.SendMessage(c.Context(), request)
	if err != nil {
		utils.LogValidationErrors(c.Context(), err, "validation failed in SendMessage", c.Path())
		return err
	}

	return helper.RespondPending(c, "Message is pending and waiting to be processed.", id)
}

func (h *MessageHandler) EditMsg(c fiber.Ctx) error {
	request := new(model.SendMessageReq)
	if err := c.Bind().Body(request); err != nil {
		utils.LogValidationErrors(c.Context(), err, "error parsing request body", c.Path())
		return fiber.ErrBadRequest
	}

	msgId := strings.TrimPrefix(c.Params("msgId"), "pwa-")
	if msgId == "" {
		utils.LogValidationErrors(c.Context(), errors.New("msgId is required"), "validation failed in EditMsg", c.Path())
		return fiber.ErrBadRequest
	}

	if err := h.UseCase.EditMessage(c.Context(), request, msgId); err != nil {
		utils.LogValidationErrors(c.Context(), err, "validation failed in EditMessage", c.Path())
		return fiber.ErrBadRequest
	}

	message := "Request accepted. Edit processing is asynchronous. Updates are applied only if the request is evaluated within the 15-minute edit window."

	return helper.RespondPending(c, message, msgId)
}

func (h *MessageHandler) DeleteMsg(c fiber.Ctx) error {
	msgId := strings.TrimPrefix(c.Params("msgId"), "pwa-")
	if msgId == "" {
		utils.LogValidationErrors(c.Context(), errors.New("msgId is required"), "validation failed in DeleteMsg", c.Path())
		return fiber.ErrBadRequest
	}

	chatId := c.Query("chatId")
	if err := h.UseCase.DeleteMessage(c.Context(), chatId, msgId); err != nil {
		utils.LogValidationErrors(c.Context(), err, "validation failed in DeleteMessage", c.Path())
		return fiber.ErrBadRequest
	}

	message := "Request accepted. Deletion processing is asynchronous. The message will be deleted shortly."

	return helper.RespondPending(c, message, msgId)
}
