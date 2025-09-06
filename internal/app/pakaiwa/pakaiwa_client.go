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

package pakaiwa

import (
	"context"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/state"
	"github.com/PakaiWA/PakaiWA/internal/pkg/config"
	"github.com/PakaiWA/PakaiWA/internal/pkg/logger"
	pwaStore "github.com/PakaiWA/PakaiWA/internal/pkg/store"
	"github.com/PakaiWA/PakaiWA/internal/pkg/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
)

func NewWhatsAppClient(ctx context.Context, log *logrus.Logger, pool *pgxpool.Pool) (*state.AppState, error) {
	store.SetOSInfo(config.GetAppName(), [3]uint32{0, 0, 0})

	container := pwaStore.InitStoreWithPool(ctx, pool, log)
	deviceStore, err := container.GetFirstDevice(ctx) // TODO: refactor for multi client
	utils.PanicIfError(err)

	clientLog := logger.NewPakaiWALog(log, config.GetAppName())
	client := whatsmeow.NewClient(deviceStore, clientLog)

	appState := &state.AppState{Client: client}

	// Register event handler
	eh := NewEventHandler(log, appState)
	client.AddEventHandler(eh.Handle)

	if client.Store.ID == nil {
		qrChan, _ := client.GetQRChannel(ctx)
		StartQRHandler(ctx, appState, qrChan, log)
	} else {
		appState.SetConnected(true)
	}

	if err := client.Connect(); err != nil {
		return nil, err
	}

	return appState, nil
}
