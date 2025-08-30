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
 * @author KAnggara75 on Sat 30/08/25 12.35
 * @project PakaiWA devices
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/devices
 */

package devices

import (
	"github.com/PakaiWA/PakaiWA/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type DeviceController struct {
	Log     *logrus.Logger
	UseCase *DeviceUseCase
}

func NewDeviceController(useCase *DeviceUseCase, log *logrus.Logger) *DeviceController {
	return &DeviceController{
		Log:     log,
		UseCase: useCase,
	}
}

func (c *DeviceController) List(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	contactId := ctx.Params("contactId")

	request := &model.ListAddressRequest{
		UserId:    auth.ID,
		ContactId: contactId,
	}

	responses, err := c.UseCase.List(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to list addresses")
		return err
	}

}
