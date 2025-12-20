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
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetAppName() string { return viper.GetString("app.name") }

func GetAppVersion() string {
	if v := viper.GetString("app.version"); v != "" {
		return v
	}
	return "v0.0.0-alpha"
}

func GetAppDesc() string {
	if v := viper.GetString("app.description"); v != "" {
		return v
	}
	return "Developer Preview"
}

func GetAppProfile() string {
	profile := "production"
	if v := viper.GetString("app.profile"); v != "" {
		return v
	}
	return profile
}

func GetJWTKey() string { return viper.GetString("app.jwt.sign_key") }

func GetAdminToken() string { return viper.GetString("app.admin.token") }

func GetAllDevicesSQL() string { return viper.GetString("app.sql.getAllDevices") }

func GetDeviceByIdSQL() string { return viper.GetString("app.sql.getDeviceById") }

func GetDeleteDeviceSQL() string { return viper.GetString("app.sql.deleteDeviceById") }

func GetAddDeviceSQL() string { return viper.GetString("app.sql.addDevice") }

func GetCountDeviceSQL() string { return viper.GetString("app.sql.countDeviceById") }

func Get40Space() string {
	return strings.Repeat(" ", 40)
}

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

func GetDefaultQuotaLimit() int64 {
	viper.SetDefault("app.quota.default_limit", 6)
	return viper.GetInt64("app.quota.default_limit")
}
