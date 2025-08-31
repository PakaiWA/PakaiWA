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
 * @author KAnggara75 on Sun 31/08/25 12.25
 * @project PakaiWA pakaiwa
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/pakaiwa
 */

package pakaiwa

import (
	"github.com/sirupsen/logrus"
	waLog "go.mau.fi/whatsmeow/util/log"
)

var _ waLog.Logger = (*LogrusAdapter)(nil)

type LogrusAdapter struct {
	entry *logrus.Entry
}

func NewPakaiWALog(l *logrus.Logger, pkgName string) waLog.Logger {
	return &LogrusAdapter{
		entry: logrus.NewEntry(l).WithField("pkg", pkgName),
	}
}

func (l *LogrusAdapter) Sub(module string) waLog.Logger {
	return &LogrusAdapter{entry: l.entry.WithField("module", module)}
}

func (l *LogrusAdapter) Infof(msg string, args ...interface{})  { l.entry.Infof(msg, args...) }
func (l *LogrusAdapter) Warnf(msg string, args ...interface{})  { l.entry.Warnf(msg, args...) }
func (l *LogrusAdapter) Errorf(msg string, args ...interface{}) { l.entry.Errorf(msg, args...) }

func (l *LogrusAdapter) Debugf(msg string, args ...interface{}) {
	if module, ok := l.entry.Data["module"].(string); ok {
		if module == "Recv" || module == "Send" {
			return
		}
	}
	l.entry.Debugf(msg, args...)
}
