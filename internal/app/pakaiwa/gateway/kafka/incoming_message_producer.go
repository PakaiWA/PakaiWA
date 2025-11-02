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
 * @author KAnggara75 on Sat 06/09/25 17.22
 * @project PakaiWA kafka
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/gateway/kafka
 */

package kafka

import (
	confluent "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/model"
	"github.com/PakaiWA/PakaiWA/internal/pkg/config"
	"github.com/PakaiWA/PakaiWA/internal/pkg/kafka"
)

type IncomingMessageProducer struct {
	kafka.Producer[*model.IncomingMessageModel]
}

func NewIncomingMessageProducer(producer *confluent.Producer, log *logrus.Logger) *IncomingMessageProducer {
	return &IncomingMessageProducer{
		Producer: kafka.Producer[*model.IncomingMessageModel]{
			Producer: producer,
			Topic:    config.GetIncomingMessageTopic(),
			Log:      log,
		},
	}
}
