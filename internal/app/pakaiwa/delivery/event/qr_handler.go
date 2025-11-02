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
 * @author KAnggara75 on Sun 07/09/25 00.12
 * @project PakaiWA event
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/delivery/event
 */

package event

import (
	"context"
	"os"

	"github.com/PakaiWA/whatsmeow"
	"github.com/mdp/qrterminal/v3"
	"github.com/sirupsen/logrus"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/state"
)

func StartQRHandler(ctx context.Context, state *state.AppState, qrChan <-chan whatsmeow.QRChannelItem, log *logrus.Logger) {
	if qrChan == nil {
		log.Warn("[WA] QR channel nil → not in pairing mode")
		return
	}

	state.SetQR("")
	state.SetConnected(false)

	go func() {
		defer func() {
			state.SetQR("")
		}()

		for {
			select {
			case <-ctx.Done():
				log.Info("[WA] QR listener: context canceled")
				return

			case evt, ok := <-qrChan:
				if !ok {
					state.Client.Disconnect()
					log.Warn("[WA] QR channel closed")
					return
				}

				switch evt.Event {
				case "code":
					state.SetQR(evt.Code)
					qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
					log.Info("[WA] QR updated — scan please")

				case "timeout":
					state.SetQR("")
					log.Warn("[WA] QR timeout — Waiting for next QR")

				case "success":
					log.Info("[WA] Login QR Success ✔️")
					state.SetQR("")
					state.SetConnected(true)
					return

				case "error":
					state.SetQR("")
					log.Error("[WA] QR error — Waiting for next QR")

				default:
					log.Infof("[WA] Login event: %s", evt.Event)
				}
			}
		}
	}()
}
