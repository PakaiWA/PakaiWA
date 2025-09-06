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
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
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

func GetJWTKey() []byte { return []byte(viper.GetString("app.jwt.sign_key")) }

func GetAdminToken() string { return viper.GetString("app.admin.token") }

func GetAllDevicesSQL() string { return viper.GetString("app.sql.getAllDevices") }

func GetDeviceByIdSQL() string { return viper.GetString("app.sql.getDeviceById") }

func GetDeleteDeviceSQL() string { return viper.GetString("app.sql.deleteDeviceById") }

func GetAddDeviceSQL() string { return viper.GetString("app.sql.addDevice") }

func GetCountDeviceSQL() string { return viper.GetString("app.sql.countDeviceById") }

func GetRedisHost() string { return viper.GetString("redis.host") }

func GetRedisPassword() string { return viper.GetString("redis.password") }

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

func Get40Space() string {
	return strings.Repeat(" ", 40)
}

func GetAppName() string { return viper.GetString("app.name") }

func GetPreFork() bool { return viper.GetBool("web.prefork") }

func GetBaseKafkaConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": viper.GetString("kafka.bootstrap.servers"),
		"security.protocol": viper.GetString("kafka.security.protocol"),
		"ssl.ca.location":   viper.GetString("kafka.ssl.ca.location"),
		"sasl.mechanism":    viper.GetString("kafka.sasl.mechanism"),
		"sasl.username":     viper.GetString("kafka.sasl.username"),
		"sasl.password":     viper.GetString("kafka.sasl.password"),
	}
}

func GetConsumerConfig() *kafka.ConfigMap {
	cfg := GetBaseKafkaConfig()
	groupId := viper.GetString("kafka.group.id")
	if groupId == "" {
		groupId = "pakaiwa-group"
	}

	(*cfg)["group.id"] = groupId
	(*cfg)["auto.offset.reset"] = viper.GetString("kafka.auto.offset.reset")
	return cfg
}

func GetProducerConfig() *kafka.ConfigMap {
	cfg := GetBaseKafkaConfig()
	(*cfg)["linger.ms"] = 5                           // tunggu sebentar supaya batch lebih besar
	(*cfg)["batch.num.messages"] = 10000              // kirim batch besar
	(*cfg)["queue.buffering.max.kbytes"] = 1048576    // default 1GB, bisa dinaikkan
	(*cfg)["queue.buffering.max.messages"] = 10000000 // default 100000

	return cfg
}
