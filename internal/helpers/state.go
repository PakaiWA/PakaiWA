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
 * @author KAnggara75 on Sun 31/08/25 09.06
 * @project PakaiWA helpers
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/helpers
 */

package helpers

import (
	"go.mau.fi/whatsmeow"
	"sync"
)

type AppState struct {
	Client    *whatsmeow.Client
	QRMu      sync.RWMutex
	LastQR    string
	Connected bool
}

func (a *AppState) SetQR(code string) {
	a.QRMu.Lock()
	a.LastQR = code
	a.QRMu.Unlock()
}

func (a *AppState) GetQR() string {
	a.QRMu.RLock()
	defer a.QRMu.RUnlock()
	return a.LastQR
}

func (a *AppState) SetConnected(v bool) {
	a.QRMu.Lock()
	a.Connected = v
	a.QRMu.Unlock()
}

func (a *AppState) GetConnected() bool {
	a.QRMu.RLock()
	defer a.QRMu.RUnlock()
	return a.Connected
}
