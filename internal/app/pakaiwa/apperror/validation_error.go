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
 * @author KAnggara75 on Sun 07/09/25 21.55
 * @project PakaiWA error
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/error
 */

package apperror

import (
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/http/dto"
)

func MapValidationErrors(err error) []dto.ValidationError {
	var errs validator.ValidationErrors
	if !errors.As(err, &errs) {
		return nil
	}

	var details []dto.ValidationError
	for _, e := range errs {
		details = append(details, dto.ValidationError{
			Field: e.Field(),
			Tag:   e.Tag(),
			Param: e.Param(),
		})
	}
	return details
}
