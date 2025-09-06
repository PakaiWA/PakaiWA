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
 * @author KAnggara75 on Sun 31/08/25 06.34
 * @project PakaiWA qr
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/model
 */

package model

type ResponseQR struct {
	QRCode  string           `json:"qr_code,omitempty"`
	QRImage string           `json:"image_url,omitempty"`
	Status  string           `json:"status,omitempty"`
	Msg     string           `json:"message,omitempty"`
	Device  ResponseQRDevice `json:"device,omitempty"`
}

type ResponseQRDevice struct {
	JID      string `json:"jid,omitempty"`
	PushName string `json:"pushName,omitempty"`
}
