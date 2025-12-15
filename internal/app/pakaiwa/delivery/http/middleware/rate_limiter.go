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
 * @author KAnggara75 on Sat 08/11/25 17.51
 * @project PakaiWA middleware
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/delivery/http/middleware
 */

package middleware

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v3"
)

type RateLimiter struct {
	mu       sync.Mutex
	requests map[string][]time.Time
	limit    int
	window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
}

func (r *RateLimiter) isAllowed(key string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	windowStart := now.Add(-r.window)
	reqs := r.requests[key]

	validReqs := reqs[:0]
	for _, t := range reqs {
		if t.After(windowStart) {
			validReqs = append(validReqs, t)
		}
	}

	r.requests[key] = validReqs

	if len(validReqs) >= r.limit {
		return false
	}

	r.requests[key] = append(r.requests[key], now)
	return true
}

func RateLimitMiddleware(limit int, window time.Duration) fiber.Handler {
	rl := NewRateLimiter(limit, window)

	return func(c fiber.Ctx) error {
		ip := c.IP()

		if !rl.isAllowed(ip) {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"success": false,
				"error": fiber.Map{
					"title":  "Too Many Requests",
					"status": 429,
					"detail": "Please slow down, you're hitting the rate limit",
				},
			})
		}

		return c.Next()
	}
}

func (r *RateLimiter) Reset(key string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.requests, key)
}
