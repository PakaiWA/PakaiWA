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
 * @author KAnggara75 on Fri 08/08/25 08.32
 * @project PakaiWA config
 * https://github.com/PakaiWA/PakaiWA/tree/main/configs
 */

package config

import (
	"github.com/spf13/viper"
	"strings"
	"time"
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

func GetJWTKey() []byte { return []byte(viper.GetString("app.jwt.sign_key")) }

func GetAdminToken() string { return viper.GetString("app.admin.token") }

func GetAllDevicesSQL() string { return viper.GetString("app.sql.getAllDevices") }

func GetDeviceByIdSQL() string { return viper.GetString("app.sql.getDeviceById") }

func GetDeleteDeviceSQL() string { return viper.GetString("app.sql.deleteDeviceById") }

func GetAddDeviceSQL() string { return viper.GetString("app.sql.addDevice") }

func GetCountDeviceSQL() string { return viper.GetString("app.sql.countDeviceById") }

func GetRedisHost() string { return viper.GetString("redis.host") }

func GetRedisPassword() string { return viper.GetString("redis.password") }

func GetLogLevel() string { return viper.GetString("log.level") }

func Get40Space() string {
	return strings.Repeat(" ", 40)
}
