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
 * @author KAnggara75 on Tue 18/11/25 06.05
 * @project PakaiWA usecase
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/usecase
 */

package usecase

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/apperror"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/http/dto"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/repository"
	"github.com/PakaiWA/PakaiWA/internal/pkg/config"
	"github.com/PakaiWA/PakaiWA/internal/pkg/security/password"
	"github.com/PakaiWA/PakaiWA/internal/pkg/utils"
)

type authUsecase struct {
	Log        *logrus.Logger
	Repository repository.UserRepository
	Validate   *validator.Validate
}

type AuthUsecase interface {
	Login(ctx context.Context, req *dto.LoginReq, iss string) (string, error)
	Register(ctx context.Context, req *dto.AuthReq) (bool, error)
}

func NewAuthUsecase(log *logrus.Logger, repo repository.UserRepository, validator *validator.Validate) AuthUsecase {
	return &authUsecase{
		Log:        log,
		Repository: repo,
		Validate:   validator,
	}
}

func (u *authUsecase) Login(ctx context.Context, req *dto.LoginReq, iss string) (string, error) {
	if err := u.Validate.Struct(req); err != nil {
		return "", err
	}

	user, err := u.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return "", err // error sistem, JANGAN ditelan
	}

	if user == nil || !password.Compare(user.Password, req.Password) {
		return "", apperror.ErrInvalidCredentials
	}

	now := time.Now().Unix()

	claims := jwt.MapClaims{
		"sub":  user.ID,
		"iss":  iss,
		"iat":  now,
		"nbf":  now,
		"exp":  time.Now().Add(time.Hour * 24 * 7).Unix(),
		"jti":  uuid.NewString(),
		"role": "user",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token
	signedToken, err := token.SignedString([]byte(config.GetJWTKey()))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (u *authUsecase) Register(ctx context.Context, req *dto.AuthReq) (bool, error) {

	if err := u.Validate.Struct(req); err != nil {
		return false, err
	}

	err := utils.ValidateStrongPassword(req.Password)
	if err != nil {
		return false, err
	}

	user, err := u.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return false, err // error sistem, JANGAN ditelan
	}

	if user != nil {
		return false, apperror.ErrUsernameExists
	}

	// Hash password
	hashed, err := password.Hash(req.Password)
	if err != nil {
		return false, err
	}

	// Create user
	u.Repository.CreateUser(ctx, req.Email, hashed)

	return true, nil
}
