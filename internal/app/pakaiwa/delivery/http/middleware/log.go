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

	"github.com/PakaiWA/PakaiWA/internal/pkg/logger/ctxmeta"
)

func FiberLogger(base *logrus.Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		// 1. Trace ID
		traceID := c.Get("X-Request-ID")
		if traceID == "" {
			traceID = uuid.NewString()
		}

		c.Locals("trace_id", traceID)
		c.Set("X-Request-ID", traceID)

		start := time.Now()

		// 2. Request-scoped logger
		entry := base.WithFields(logrus.Fields{
			"trace_id": traceID,
			"method":   c.Method(),
			"path":     c.Path(),
			"ip":       c.IP(),
		})

		// 3. Inject context SEBELUM c.Next()
		ctx := c.Context()
		ctx = ctxmeta.WithTraceID(ctx, traceID)
		ctx = ctxmeta.WithLogger(ctx, entry)
		c.SetContext(ctx)

		entry.WithField("event", "request_start").Info("request started")

		// 4. Continue chain
		err := c.Next()

		// 5. Post-request logging
		status := c.Response().StatusCode()
		latency := time.Since(start)

		fields := logrus.Fields{
			"status":  status,
			"latency": latency.String(),
		}

		// Optional: jti
		if jti, ok := c.Locals("jti").(string); ok && jti != "" {
			fields["jti"] = jti
		}

		logEntry := entry.WithFields(fields)

		switch {
		case err != nil:
			logEntry.WithField("event", "request_end").Error("request failed")
		case status >= fiber.StatusInternalServerError:
			logEntry.WithField("event", "request_end").Error("server error")
		case status >= fiber.StatusBadRequest:
			logEntry.WithField("event", "request_end").WithField("category", "client_error").Info("request handled")
		default:
			logEntry.WithField("event", "request_end").Info("request handled")
		}

		return err
	}
}
