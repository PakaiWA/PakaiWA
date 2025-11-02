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
 * @author KAnggara75 on Sat 06/09/25 17.34
 * @project PakaiWA config
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/pkg/config
 */

package config

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/spf13/viper"
)

var (
	once          sync.Once
	localCertPath string
)

func downloadCertOnce(url string) string {
	once.Do(func() {
		dir := "/tmp/kafka"
		if err := os.MkdirAll(dir, 0755); err != nil {
			panic("Failed to create dir for Kafka cert: " + err.Error())
		}

		localCertPath = filepath.Join(dir, "kafka.cert")

		// Skip download jika file sudah ada
		if _, err := os.Stat(localCertPath); os.IsNotExist(err) {
			resp, err := http.Get(url)
			if err != nil {
				panic("Failed to download Kafka cert: " + err.Error())
			}
			defer resp.Body.Close()

			out, err := os.Create(localCertPath)
			if err != nil {
				panic("Failed to create Kafka cert file: " + err.Error())
			}
			defer out.Close()

			if _, err := io.Copy(out, resp.Body); err != nil {
				panic("Failed to write Kafka cert file: " + err.Error())
			}
		}
	})
	return localCertPath
}

func GetBaseKafkaConfig() *kafka.ConfigMap {
	certURL := viper.GetString("kafka.ssl.ca.location") // ini URL dari config
	localCert := downloadCertOnce(certURL)

	return &kafka.ConfigMap{
		"bootstrap.servers": viper.GetString("kafka.bootstrap.servers"),
		"security.protocol": viper.GetString("kafka.security.protocol"),
		"ssl.ca.location":   localCert,
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

func GetIncomingMessageTopic() string {
	topic := viper.GetString("kafka.out.topic.incoming_msg")
	if topic == "" {
		topic = "pakaiwa-incoming-message"
	}
	return topic
}

func GetDeliveryStatusTopic() string {
	topic := viper.GetString("kafka.out.topic.delivery_status")
	if topic == "" {
		topic = "pakaiwa-delivery-status"
	}
	return topic
}
