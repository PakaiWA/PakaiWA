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
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/spf13/viper"
)

var (
	once          sync.Once
	localCertPath string
)

func downloadCertOnce(certURL string) string {
	once.Do(func() {
		dir := "/tmp/kafka"
		if err := os.MkdirAll(dir, 0755); err != nil {
			panic("Failed to create dir for Kafka cert: " + err.Error())
		}

		localCertPath = filepath.Join(dir, "kafka.cert")

		// Skip download jika file sudah ada
		if _, err := os.Stat(localCertPath); os.IsNotExist(err) {
			resp, err := http.Get(certURL)
			if err != nil {
				panic("Failed to download Kafka cert: " + err.Error())
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					panic("Failed to close Kafka cert response body: " + err.Error())
				}
			}(resp.Body)

			out, err := os.Create(localCertPath)
			if err != nil {
				panic("Failed to create Kafka cert file: " + err.Error())
			}
			defer func(out *os.File) {
				err := out.Close()
				if err != nil {
					panic("Failed to close Kafka cert file: " + err.Error())
				}
			}(out)

			if _, err := io.Copy(out, resp.Body); err != nil {
				panic("Failed to write Kafka cert file: " + err.Error())
			}
		}
	})
	return localCertPath
}

func GetBaseKafkaConfig() *kafka.ConfigMap {
	certLocation := viper.GetString("kafka.ssl.ca.location")

	// cek apakah certLocation adalah URL
	if u, err := url.Parse(certLocation); err == nil && (strings.HasPrefix(u.Scheme, "http")) {
		certLocation = downloadCertOnce(certLocation)
	} else {
		// pastikan file ada
		if _, err := os.Stat(certLocation); os.IsNotExist(err) {
			panic("Kafka cert file not found: " + certLocation)
		}
	}

	return &kafka.ConfigMap{
		"bootstrap.servers": viper.GetString("kafka.bootstrap.servers"),
		"security.protocol": viper.GetString("kafka.security.protocol"),
		"ssl.ca.location":   certLocation,
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
