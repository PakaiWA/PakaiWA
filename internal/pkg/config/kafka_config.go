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
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/spf13/viper"
)

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
