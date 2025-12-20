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
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/apperror"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/model"
	"github.com/PakaiWA/PakaiWA/internal/pkg/config"
)

func NewKafkaProducer(log *logrus.Logger) *kafka.Producer {
	producer, err := kafka.NewProducer(config.GetProducerConfig())
	apperror.PanicIfError(err)

	return producer
}

type Producer[T model.Event] struct {
	Producer *kafka.Producer
	Topic    string
	Log      *logrus.Logger
}

func (p *Producer[T]) GetTopic() *string {
	return &p.Topic
}

func (p *Producer[T]) Send(event T) error {
	value, err := json.Marshal(event)
	if err != nil {
		p.Log.WithError(err).Error("failed to marshal event")
		return err
	}

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     p.GetTopic(),
			Partition: kafka.PartitionAny,
		},
		Value: value,
		Key:   []byte(event.GetId()),
	}

	err = p.Producer.Produce(message, nil)
	if err != nil {
		p.Log.WithError(err).Error("failed to produce message")
		return err
	}

	return nil
}
