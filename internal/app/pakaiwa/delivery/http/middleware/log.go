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
 * @author KAnggara75 on Sat 08/11/25 21.53
 * @project PakaiWA middleware
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/delivery/http/middleware
 */

package middleware

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func FiberLogger(log *logrus.Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		traceID := c.Get("X-Request-ID", uuid.New().String())
		c.Locals("trace_id", traceID)
		start := time.Now()

		err := c.Next()

		latency := time.Since(start)
		status := c.Response().StatusCode()
		method := c.Method()
		path := c.Path()
		ip := c.IP()

		entry := log.WithFields(logrus.Fields{
			"trace_id": traceID,
			"status":   status,
			"method":   method,
			"path":     path,
			"ip":       ip,
			"latency":  latency.String(),
		})

		if err != nil {
			entry.WithError(err).Error("request failed")
			return err
		}

		switch {
		case status >= 500:
			entry.Error("server error")
		case status >= 400:
			entry.Warn("client error")
		default:
			entry.Info("request handled")
		}

		return nil
	}
}
