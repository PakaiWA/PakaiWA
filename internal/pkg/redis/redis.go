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
 * @author KAnggara75 on Sat 06/09/25 11.07
 * @project PakaiWA redis
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/pkg/redis
 */

package redis

import (
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/PakaiWA/PakaiWA/internal/pkg/config"
)

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         config.GetRedisHost(),
		Password:     config.GetRedisPassword(),
		DB:           0,
		DialTimeout:  15 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})
}
