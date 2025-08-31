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
 * @project PakaiWA middleware
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/middleware
 */

package middleware

import (
	"github.com/PakaiWA/PakaiWA/internal/model"
	"github.com/PakaiWA/PakaiWA/internal/usecase"
	"github.com/PakaiWA/PakaiWA/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func NewAuth(usecase *usecase.UserUseCase, rateLimiterUtil *utils.RateLimiterUtil) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := &model.VerifyUserRequest{Token: ctx.Get("Authorization", "NOT_FOUND")}
		usecase.Log.Debugf("Authorization : %s", request.Token)

		authData, err := usecase.Verify(ctx.UserContext(), request)
		if err != nil {
			usecase.Log.Warnf("Failed find user by token : %+v", err)
			return fiber.ErrUnauthorized
		}

		if !rateLimiterUtil.IsAllowed(ctx.UserContext(), authData) {
			usecase.Log.Warnf("User is not allowed because too many request : %+v", err)
			return fiber.ErrTooManyRequests
		}

		usecase.Log.Debugf("User : %+v", authData.ID)
		ctx.Locals("auth", authData)
		return ctx.Next()
	}
}

func GetUser(ctx *fiber.Ctx) *model.Model {
	return ctx.Locals("auth").(*model.Model)
}
