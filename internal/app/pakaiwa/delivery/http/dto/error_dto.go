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
 * @author KAnggara75 on Sun 07/09/25 21.43
 * @project PakaiWA dto
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/delivery/http/dto
 */

package dto

type ValidationErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Param string `json:"param,omitempty"`
}
