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
 * @project PakaiWA users
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/users
 */

package users

import (
	"context"
	"github.com/PakaiWA/PakaiWA/internal/app/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserUseCase struct {
	Log *logrus.Logger
}

func NewUserUseCase(logger *logrus.Logger) *UserUseCase {
	return &UserUseCase{
		Log: logger,
	}
}

func (c *UserUseCase) Verify(ctx context.Context, request *VerifyUserRequest) (*auth.Model, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	user := new(entity.User)
	if err := c.UserRepository.FindByToken(tx, user, request.Token); err != nil {
		c.Log.Warnf("Failed find user by token : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return &model.Auth{ID: user.ID}, nil
}
