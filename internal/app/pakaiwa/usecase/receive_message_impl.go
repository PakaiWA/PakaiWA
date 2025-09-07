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
 * @author KAnggara75 on Sun 07/09/25 15.28
 * @project PakaiWA usecase
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/usecase
 */

package usecase

import (
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/model"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/gateway/kafka"
	"github.com/sirupsen/logrus"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
)

type receiveMessageUsecase struct {
	Log      *logrus.Logger
	Producer *kafka.IncomingMessageProducer
}

func NewReceiveMessageUsecase(log *logrus.Logger, producer *kafka.IncomingMessageProducer) ReceiveMessageUsecase {
	return &receiveMessageUsecase{Log: log, Producer: producer}
}

func (uc *receiveMessageUsecase) ProcessIncomingMessage(msg *waE2E.Message, info types.MessageInfo, rawMsg *waE2E.Message) {

	incomingMsgModel, err := model.ToIncomingMessageModel(msg, info, rawMsg)
	if err != nil {
		uc.Log.Error(err)
	}

	err = uc.Producer.Send(incomingMsgModel)
	if err != nil {
		uc.Log.Error(err)
	}
}
