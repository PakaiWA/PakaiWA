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
	"encoding/json"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/gateway/kafka"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sirupsen/logrus"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/encoding/protojson"
)

type receiveMessageUsecase struct {
	Log      *logrus.Logger
	Producer *kafka.IncomingMessageProducer
}

func NewReceiveMessageUsecase(log *logrus.Logger, producer *kafka.IncomingMessageProducer) ReceiveMessageUsecase {
	return &receiveMessageUsecase{Log: log, Producer: producer}
}

func (uc *receiveMessageUsecase) ProcessMessage(msg *waE2E.Message, info types.MessageInfo, rawMsg *waE2E.Message) {

	marshaler := protojson.MarshalOptions{
		Multiline:       false,
		EmitUnpopulated: false,
		UseEnumNumbers:  false,
	}

	msgJSON, err := marshaler.Marshal(msg)
	if err != nil {
		uc.Log.Errorf("failed to marshal msg to JSON: %v", err)
	} else {
		uc.Log.Infof("Message JSON: %s", msgJSON)
	}

	rawJSON, _ := marshaler.Marshal(rawMsg)
	uc.Log.Infof("Raw Message JSON: %s", rawJSON)

	infoJSON, err := json.Marshal(info)
	if err != nil {
		uc.Log.Errorf("failed to marshal MessageInfo: %v", err)
	} else {
		uc.Log.Infof("Info JSON: %s", infoJSON)
	}

	//uc.Producer.Send(msg)

	log.Infof("{\"msg\": %s,\"info\": %s}", rawJSON, infoJSON)

}
