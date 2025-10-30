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
 * @author KAnggara75 on Sat 06/09/25 10.56
 * @project PakaiWA config
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/pkg/config
 */

package config

import (
	"time"

	"github.com/spf13/viper"
)

func GetDBConn() string {
	return viper.GetString("db.pakaiwa.host")
}

func GetDBMinConn() int32 {
	minConn := viper.GetInt32("db.pakaiwa.MinConns")
	if minConn <= 0 {
		minConn = 1
	}
	return minConn
}

func GetDBMaxConn() int32 {
	maxConn := viper.GetInt32("db.pakaiwa.MaxConns")
	if maxConn <= 0 {
		maxConn = 10
	}
	return maxConn
}

func GetDBHealthCheckPeriod() time.Duration {
	if viper.IsSet("db.pakaiwa.HealthCheckPeriod") {
		val := viper.GetDuration("db.pakaiwa.HealthCheckPeriod")
		if val > 0 {
			return val
		}
	}
	return 1 * time.Minute
}

func GetConnectTimeout() time.Duration {
	if viper.IsSet("db.pakaiwa.connectTimeout") {
		val := viper.GetDuration("db.pakaiwa.connectTimeout")
		if val > 0 {
			return val
		}
	}
	return 30 * time.Second
}

func DetMaxConnIdleTime() time.Duration {
	if viper.IsSet("db.pakaiwa.maxConnIdleTime") {
		val := viper.GetDuration("db.pakaiwa.maxConnIdleTime")
		if val > 0 {
			return val
		}
	}
	return 30 * time.Second
}
