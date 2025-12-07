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
 * @author KAnggara75 on Sat 06/09/25 10.19
 * @project PakaiWA dto
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/delivery/http/dto
 */

package dto

type BaseResponse struct {
	Success bool            `json:"success"`
	Data    any             `json:"data,omitempty"`
	Error   *ProblemDetails `json:"error,omitempty"`
	Meta    *Meta           `json:"meta,omitempty"`
}

type Meta struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}
