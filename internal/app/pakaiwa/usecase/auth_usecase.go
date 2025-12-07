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
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/model"
	"github.com/PakaiWA/PakaiWA/internal/pkg/config"
	"github.com/PakaiWA/PakaiWA/internal/pkg/utils"
)

type authUsecase struct {
	Log      *logrus.Logger
	Validate *validator.Validate
}

type AuthUsecase interface {
	Login(req *model.LoginReq) (string, error)
	Register(req *model.AuthReq) (bool, error)
}

func NewAuthUsecase(log *logrus.Logger, validator *validator.Validate) AuthUsecase {
	return &authUsecase{
		Log:      log,
		Validate: validator,
	}
}

func (u *authUsecase) Login(req *model.LoginReq) (string, error) {
	if err := u.Validate.Struct(req); err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"sub":  req.Email,
		"iss":  "config.Issuer",
		"aud":  jwt.ClaimStrings{"config.Audience"},
		"iat":  time.Now().Unix(),
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

func (u *authUsecase) Register(req *model.AuthReq) (bool, error) {

	if err := u.Validate.Struct(req); err != nil {
		return false, err
	}

	err := utils.ValidateStrongPassword(req.Password)
	if err != nil {
		return false, err
	}

	return true, nil
}
