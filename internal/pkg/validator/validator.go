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
 * @author KAnggara75 on Sat 06/09/25 11.09
 * @project PakaiWA validator
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/pkg/validator
 */

package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		tag := fld.Tag.Get("json")
		if tag == "" {
			return fld.Name
		}
		name := strings.SplitN(tag, ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// Register custom validator
	validate.RegisterValidation("strongPassword", ValidateStrongPassword)

	return validate
}
