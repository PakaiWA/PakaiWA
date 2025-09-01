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
	"github.com/mdp/qrterminal/v3"
	"github.com/sirupsen/logrus"
	"go.mau.fi/whatsmeow"
	"os"
)

func StartQRHandler(state *AppState, qrChan <-chan whatsmeow.QRChannelItem, log *logrus.Logger) {
	if qrChan == nil {
		state.SetConnected(true)
		return
	}

	go func() {
		for evt := range qrChan {
			switch evt.Event {
			case "code":
				state.SetQR(evt.Code)
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
				log.Info("[WA] Scan QR ini dengan WhatsApp (Linked devices)")
			case "success":
				state.SetQR("")
				state.SetConnected(true)
				log.Info("[WA] Login QR sukses ✔️")
			default:
				log.Infof("[WA] Login event: %s", evt.Event)
			}
		}
	}()
}
