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
 * @author KAnggara75 on Sat 06/09/25 17.35
 * @project PakaiWA config
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/pkg/config
 */

package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

func GetAppName() string { return viper.GetString("app.name") }

func GetJWTKey() []byte { return []byte(viper.GetString("app.jwt.sign_key")) }

func GetAdminToken() string { return viper.GetString("app.admin.token") }

func GetAllDevicesSQL() string { return viper.GetString("app.sql.getAllDevices") }

func GetDeviceByIdSQL() string { return viper.GetString("app.sql.getDeviceById") }

func GetDeleteDeviceSQL() string { return viper.GetString("app.sql.deleteDeviceById") }

func GetAddDeviceSQL() string { return viper.GetString("app.sql.addDevice") }

func GetCountDeviceSQL() string { return viper.GetString("app.sql.countDeviceById") }

func Get40Space() string {
	return strings.Repeat(" ", 40)
}

func GetPreFork() bool { return viper.GetBool("web.prefork") }

func GetLogLevel() logrus.Level {
	viper.SetDefault("log.level", "info")
	raw := strings.TrimSpace(viper.GetString("log.level"))
	if raw == "" {
		return logrus.InfoLevel
	}

	lvl, err := logrus.ParseLevel(strings.ToLower(raw))
	if err != nil {
		return logrus.InfoLevel
	}

	return lvl
}
