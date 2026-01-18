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
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/apperror"
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
		return apperror.BadRequest(c, err.Error())
	}

	msgId := strings.TrimPrefix(c.Params("msgId"), "pwa-")
	if msgId == "" {
		utils.LogValidationErrors(c.Context(), errors.New("msgId is required"), "validation failed in EditMsg", c.Path())
		return apperror.BadRequest(c, "msgId is required")
	}

	if err := h.UseCase.EditMessage(c.Context(), request, msgId); err != nil {
		utils.LogValidationErrors(c.Context(), err, "validation failed in EditMessage", c.Path())
		return apperror.BadRequest(c, err.Error())
	}

	message := "Request accepted. Edit processing is asynchronous. Updates are applied only if the request is evaluated within the 15-minute edit window."

	return helper.RespondPending(c, message, msgId)
}

func (h *MessageHandler) DeleteMsg(c fiber.Ctx) error {
	ctx := c.Context()

	// =====================
	// Validate msgId
	// =====================
	rawMsgId := c.Params("msgId")
	if rawMsgId == "" {
		utils.LogValidationErrors(ctx, errors.New("msgId is required"), "DeleteMsg", c.Path())
		return apperror.BadRequest(c, "msgId is required")
	}

	if !strings.HasPrefix(rawMsgId, "pwa-") {
		utils.LogValidationErrors(ctx, errors.New("invalid msgId format"), "DeleteMsg", c.Path())
		return apperror.BadRequest(c, "invalid msgId format")
	}

	msgId := strings.TrimPrefix(rawMsgId, "pwa-")

	// =====================
	// Validate chatId
	// =====================
	chatId := c.Query("chatId")
	if chatId == "" {
		utils.LogValidationErrors(ctx, errors.New("chatId is required"), "DeleteMsg", c.Path())
		return apperror.BadRequest(c, "chatId is required")
	}

	// =====================
	// Parse isGroup (STRICT)
	// =====================
	isGroupStr := c.Query("isGroup", "false")
	isGroup, err := strconv.ParseBool(isGroupStr)
	if err != nil {
		utils.LogValidationErrors(ctx, errors.New("isGroup must be boolean"), "DeleteMsg", c.Path())
		return apperror.BadRequest(c, "isGroup must be boolean")
	}

	// =====================
	// Call use case
	// =====================
	if err := h.UseCase.DeleteMessage(ctx, chatId, msgId, isGroup); err != nil {
		// utils.LogDomainError(ctx, err, "DeleteMessage failed", c.Path())
		return apperror.Internal(c, err.Error())
	}

	// =====================
	// Async accepted response
	// =====================
	return helper.RespondPending(
		c,
		"Request accepted. Deletion processing is asynchronous.",
		msgId,
	)
}
