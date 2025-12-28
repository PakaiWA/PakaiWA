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
	"context"
	"encoding/json"
	"time"

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

func (p *Producer[T]) Send(ctx context.Context, event T) error {
	value, err := json.Marshal(event)
	if err != nil {
		p.Log.WithError(err).Error("failed to marshal event")
		return err
	}

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &p.Topic,
			Partition: kafka.PartitionAny,
		},
		Key:   []byte(event.GetId()),
		Value: value,
	}

	for {
		err = p.Producer.Produce(msg, nil)
		if err == nil {
			return nil
		}

		// Queue penuh → tunggu sebentar
		if err.(kafka.Error).Code() == kafka.ErrQueueFull {
			select {
			case <-time.After(50 * time.Millisecond):
				continue
			case <-ctx.Done():
				return ctx.Err()
			}
		}

		p.Log.WithError(err).Error("failed to produce kafka message")
		return err
	}
}

func StartProducerPollLoop(
	ctx context.Context,
	producer *kafka.Producer,
	log *logrus.Logger,
) {
	go func() {
		log.Info("Kafka producer poll loop started")

		for {
			select {
			case <-ctx.Done():
				log.Info("Kafka producer poll loop stopping")
				return

			case ev := <-producer.Events():
				switch e := ev.(type) {

				case *kafka.Message:
					if e.TopicPartition.Error != nil {
						log.WithFields(logrus.Fields{
							"topic":     *e.TopicPartition.Topic,
							"partition": e.TopicPartition.Partition,
							"offset":    e.TopicPartition.Offset,
							"module":    "Kafka",
						}).WithError(e.TopicPartition.Error).
							Error("Kafka delivery failed")
					} else {
						log.WithFields(logrus.Fields{
							"topic":     *e.TopicPartition.Topic,
							"partition": e.TopicPartition.Partition,
							"offset":    e.TopicPartition.Offset,
							"module":    "Kafka",
						}).Debug("Kafka message delivered")
					}

				case kafka.Error:
					log.WithError(e).Error("Kafka error")

				default:
					// abaikan event lain
				}
			}
		}
	}()
}
