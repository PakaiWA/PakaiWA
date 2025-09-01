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
 * @author KAnggara75 on Mon 01/09/25 20.59
 * @project PakaiWA configs
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/configs
 */

package configs

import (
	"context"
	"github.com/PakaiWA/PakaiWA/internal/handler"
	"github.com/PakaiWA/PakaiWA/internal/helpers"
	"github.com/PakaiWA/PakaiWA/internal/pakaiwa"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
)

func NewWhatsAppClient(ctx context.Context, pool *pgxpool.Pool, log *logrus.Logger) (*pakaiwa.AppState, error) {
	store.SetOSInfo(GetAppName(), [3]uint32{0, 0, 0})

	container := pakaiwa.InitStoreWithPool(ctx, pool, log)
	deviceStore, err := container.GetFirstDevice(ctx) // TODO: refactor for multi client
	helpers.PanicIfError(err)

	clientLog := pakaiwa.NewPakaiWALog(log, GetAppName())
	client := whatsmeow.NewClient(deviceStore, clientLog)

	state := &pakaiwa.AppState{Client: client}

	// Register event handler
	eh := handler.NewEventHandler(log, state)
	client.AddEventHandler(eh.Handle)

	if client.Store.ID == nil {
		qrChan, _ := client.GetQRChannel(ctx)
		pakaiwa.StartQRHandler(ctx, state, qrChan, log)
	} else {
		state.SetConnected(true)
	}

	if err := client.Connect(); err != nil {
		return nil, err
	}

	return state, nil
}
