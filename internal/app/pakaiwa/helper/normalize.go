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
 * @author KAnggara75 on Sat 06/09/25 11.38
 * @project PakaiWA helper
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/helper
 */

package helper

import (
	"fmt"
	"go.mau.fi/whatsmeow/types"
	"strings"
)

func NormalizeJID(s string) (types.JID, error) {
	s = strings.TrimSpace(s)
	if strings.Contains(s, "@") {
		j, err := types.ParseJID(s)
		if err != nil {
			return types.JID{}, err
		}
		return j, nil
	}
	if isAllDigits(s) {
		return types.JID{User: s, Server: types.DefaultUserServer}, nil
	}
	return types.JID{}, fmt.Errorf("invalid phone_number %s", s)
}

func isAllDigits(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return s != ""
}

func NormalizeNumber(jid string) string {
	if i := strings.Index(jid, ":"); i != -1 {
		jid = jid[:i]
	}
	if i := strings.Index(jid, "@"); i != -1 {
		jid = jid[:i]
	}
	return jid
}
