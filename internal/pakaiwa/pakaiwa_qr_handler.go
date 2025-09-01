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
 * @author KAnggara75 on Mon 01/09/25 10.59
 * @project PakaiWA pakaiwa
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/pakaiwa
 */

package pakaiwa

import (
	"context"
	"github.com/mdp/qrterminal/v3"
	"github.com/sirupsen/logrus"
	"go.mau.fi/whatsmeow"
	"os"
)

func StartQRHandler(ctx context.Context, state *AppState, qrChan <-chan whatsmeow.QRChannelItem, log *logrus.Logger) {
	if qrChan == nil {
		log.Warn("[WA] QR channel nil → bukan mode pairing. Menunggu status dari event lain.")
		return
	}

	state.SetConnected(false)
	state.SetQR("")

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
					log.Warn("[WA] QR channel closed")
					return
				}

				switch evt.Event {
				case "code":
					state.SetQR(evt.Code)
					qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
					log.Info("[WA] QR diperbarui — silakan scan dari WhatsApp (Linked devices)")

				case "timeout":
					state.SetQR("")
					log.Warn("[WA] QR timeout — menunggu kode baru…")

				case "success":
					log.Info("[WA] Login QR sukses ✔️")
					state.SetQR("")
					state.SetConnected(true)
					return

				case "error":
					state.SetQR("")
					log.Error("[WA] QR error — menunggu pembaruan berikutnya")

				default:
					log.Infof("[WA] Login event: %s", evt.Event)
				}
			}
		}
	}()
}
