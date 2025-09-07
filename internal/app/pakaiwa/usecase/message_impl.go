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
 * @author KAnggara75 on Sun 07/09/25 12.51
 * @project PakaiWA usecase
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/usecase
 */

package usecase

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
)

type messageUsecase struct {
	Log      *logrus.Logger
	Producer *kafka.Producer
}

func NewMessageUsecase(log *logrus.Logger, producer *kafka.Producer) MessageUsecase {
	return &messageUsecase{Log: log, Producer: producer}
}

func (uc *messageUsecase) ProcessMessageEvent(msg *waE2E.Message, info *types.MessageInfo) error {
	//TODO implement me
	panic("implement me")
}

func (uc *messageUsecase) HandleLogout(client *whatsmeow.Client) error {
	//TODO implement me
	panic("implement me")
}
