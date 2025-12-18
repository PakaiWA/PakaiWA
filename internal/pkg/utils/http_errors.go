/*
 * Copyright (c) 2025 KAnggara
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * See <https://www.gnu.org/licenses/gpl-3.0.html>.
 *
 * @author KAnggara on Thursday 18/12/2025 18.51
 * @project PakaiWA
 * ~/work/PakaiWA/PakaiWA/internal/pkg/utils
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/pkg/utils
 */

package utils

import "github.com/gofiber/fiber/v3"

func TooManyRequests(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
		"error": fiber.Map{
			"type":     "https://api.pakaiwa.my.id/problems/rate-limit-exceeded",
			"title":    "Too Many Requests",
			"status":   429,
			"detail":   "Please slow down, you're hitting the rate limit",
			"instance": ctx.Path(),
		},
	})
}
