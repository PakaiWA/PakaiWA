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
 * @author KAnggara75 on Sun 07/09/25 00.14
 * @project PakaiWA bootstrap
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/bootstrap
 */

package bootstrap

import (
	"context"

	"github.com/PakaiWA/whatsmeow"
	"github.com/PakaiWA/whatsmeow/store"
	confluent "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/apperror"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/event"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/gateway/kafka"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/state"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/usecase"
	"github.com/PakaiWA/PakaiWA/internal/pkg/config"
	"github.com/PakaiWA/PakaiWA/internal/pkg/logger"
	pwaStore "github.com/PakaiWA/PakaiWA/internal/pkg/store"
)

type PwaContext struct {
	Log      *logrus.Logger
	Pool     *pgxpool.Pool
	Producer *confluent.Producer
}

func InitWhatsapp(ctx context.Context, b *PwaContext) (*state.AppState, error) {
	log := b.Log
	pool := b.Pool
	producer := b.Producer

	store.SetOSInfo(config.GetAppName(), [3]uint32{0, 0, 0})

	container := pwaStore.InitStoreWithPool(ctx, pool, log)
	deviceStore, err := container.GetFirstDevice(ctx) // TODO: refactor for multi client
	apperror.PanicIfError(err)

	clientLog := logger.NewPakaiWALog(log, "debug", config.GetAppName())
	client := whatsmeow.NewClient(deviceStore, clientLog)

	appState := &state.AppState{Client: client}

	// INJECT dependencies
	incomingMsgProducer := kafka.NewIncomingMessageProducer(producer, log)
	receiveMsgUC := usecase.NewReceiveMessageUsecase(log, incomingMsgProducer)

	deliveryStatusProducer := kafka.NewDeliveryStatusProducer(producer, log)
	deliveryUC := usecase.NewDeliveryStatusUsecase(log, deliveryStatusProducer)

	eventHandler := event.HandleEvent{
		Ctx:            ctx,
		PakaiWA:        appState,
		Producer:       producer,
		ReceiveMsgUC:   receiveMsgUC,
		DeliveryStatus: deliveryUC,
	}

	// Register event handler
	client.AddEventHandler(eventHandler.Handle)

	if client.Store.ID == nil {
		qrChan, _ := client.GetQRChannel(ctx)
		event.StartQRHandler(ctx, appState, qrChan, log)
	} else {
		appState.SetConnected(true)
	}

	if err := client.Connect(); err != nil {
		return nil, err
	}

	return appState, nil
}
