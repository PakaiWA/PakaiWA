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
 * @author KAnggara75 on Sat 30/08/25 12.50
 * @project PakaiWA utils
 * https://github.com/PakaiWA/PakaiWA/tree/main/pkg/utils
 */

package utils

import (
	"context"
	"fmt"
	"github.com/PakaiWA/PakaiWA/internal/model"
	"github.com/redis/go-redis/v9"
	"time"
)

type RateLimiterUtil struct {
	Redis      *redis.Client
	MaxRequest int64
	Duration   time.Duration
}

func NewRateLimiterUtil(redis *redis.Client) *RateLimiterUtil {
	return &RateLimiterUtil{
		Redis:      redis,
		MaxRequest: 1,
		Duration:   time.Second * 1,
	}
}

func (u RateLimiterUtil) IsAllowed(ctx context.Context, auth *model.Model) bool {
	key := auth.ID

	increment, err := u.Redis.Incr(ctx, key).Result()
	if err != nil {
		fmt.Println("Error incrementing:", err)
		return false
	}

	if increment == 1 {
		err := u.Redis.Expire(ctx, key, u.Duration).Err()
		if err != nil {
			fmt.Println("Error setting expiration:", err)
			return false
		}
	}

	return increment <= u.MaxRequest
}
