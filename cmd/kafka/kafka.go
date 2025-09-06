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
 * @author KAnggara75 on Fri 05/09/25 09.28
 * @project PakaiWA kafka
 * https://github.com/PakaiWA/PakaiWA/tree/main/cmd/kafka
 */

package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/KAnggara75/scc2go"
	"github.com/PakaiWA/PakaiWA/internal/pkg/kafka"
	"github.com/PakaiWA/PakaiWA/internal/pkg/logger"
	"github.com/PakaiWA/PakaiWA/internal/pkg/utils"
	confluent "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	scc2go.GetEnv(os.Getenv("SCC_URL"), os.Getenv("AUTH"))
}

func main() {
	log := logger.NewLogger()
	consumer := kafka.NewKafkaConsumer(log)
	defer func(consumer *confluent.Consumer) {
		err := consumer.Close()
		if err != nil {
			log.Errorf("Failed to close consumer: %v", err)
		}
	}(consumer)

	err := consumer.Subscribe("pakaiwa-incoming-message", nil)
	utils.PanicIfError(err)

	topic := "pakaiwa-incoming-message"
	if err := consumer.SubscribeTopics([]string{topic}, nil); err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}

	producer := kafka.NewKafkaProducer(log)
	defer producer.Close()

	// Context + signal handling
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 🔹 Goroutine khusus untuk delivery report
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *confluent.Message:
				if ev.TopicPartition.Error != nil {
					log.Errorf("fail: %v", ev.TopicPartition.Error)
				} else {
					log.Infof("[%d] OK %v",
						ev.TopicPartition.Partition,
						ev.TopicPartition.Offset)
				}
			}
		}
	}()

	// 🔹 Goroutine producer
	// Channel untuk job
	jobs := make(chan int, 1000)

	// Buat worker pool (misal 100 worker)
	workerCount := 1000
	for w := 0; w < workerCount; w++ {
		go func(id int) {
			for i := range jobs {
				value := fmt.Sprintf("Hello %d => %d", i, time.Now().Unix())

				err := producer.Produce(&confluent.Message{
					TopicPartition: confluent.TopicPartition{Topic: &topic, Partition: confluent.PartitionAny},
					Key:            nil,
					Value:          []byte(value),
				}, nil)

				if err != nil {
					log.Errorf("[Worker %d] Failed to produce: %v", id, err)
				}
			}
		}(w)
	}

	// Kirim 1 juta pesan ke channel
	for i := 0; i < 5_000_000; i++ {
		jobs <- i
	}
	close(jobs)

	// 🔹 Goroutine consumer
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				var kafkaErr confluent.Error
				msg, err := consumer.ReadMessage(time.Second)
				if err == nil {
					log.Infof("Consumed message on %s: %s", msg.TopicPartition, string(msg.Value))
				} else if errors.As(err, &kafkaErr) && !kafkaErr.IsTimeout() {
					log.Errorf("Consumer error: %v", err)
				}
			}
		}
	}()

	// Tunggu signal
	<-ctx.Done()
	log.Info("Shutting down...")

	// Flush semua pesan producer sebelum exit
	producer.Flush(5000) // tunggu max 5 detik

}
