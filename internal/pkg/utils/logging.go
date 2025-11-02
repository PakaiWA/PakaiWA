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
 * @author KAnggara75 on Sun 07/09/25 22.21
 * @project PakaiWA utils
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/pkg/utils
 */

package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func LogValidationErrors(log *logrus.Logger, msg string, err error) {
	var validationError validator.ValidationErrors
	if errors.As(err, &validationError) {
		for _, v := range validationError {
			log.WithFields(logrus.Fields{
				"field": v.Field(),
				"tag":   v.Tag(),
				"param": v.Param(),
			}).Error(msg)
		}
	} else {
		log.WithError(err).Error(msg)
	}
}
