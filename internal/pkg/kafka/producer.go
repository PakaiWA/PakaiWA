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
 * @author KAnggara75 on Sat 06/09/25 11.04
 * @project PakaiWA kafka
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/pkg/kafka
 */

package kafka

import (
	"github.com/PakaiWA/PakaiWA/internal/pkg/config"
	"github.com/PakaiWA/PakaiWA/internal/pkg/utils"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

func NewKafkaProducer(log *logrus.Logger) *kafka.Producer {
	producer, err := kafka.NewProducer(config.GetProducerConfig())
	utils.PanicIfError(err)

	return producer
}

func SendKafkaMessage(producer *kafka.Producer, topic string, key string, value string, log *logrus.Logger) {
	if producer == nil {
		log.Warn("Kafka producer is not initialized")
		return
	}

	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          []byte(value),
	}, nil)
	if err != nil {
		log.Errorf("Failed to produce message to topic %s: %v", topic, err)
		return
	}

	e := <-producer.Events()
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		log.Errorf("Failed to deliver message to topic %s: %v", topic, m.TopicPartition.Error)
	} else {
		log.Infof("Message delivered to topic %s [%d] at offset %v", *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
}
