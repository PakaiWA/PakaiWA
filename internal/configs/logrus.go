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
 * @author KAnggara75 on Sat 30/08/25 17.52
 * @project PakaiWA configs
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/configs
 */

package configs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"sort"
	"strings"
	"time"
)

type OrderedJSONFormatter struct {
	PadLevelTo      int
	TimestampFormat string // default RFC3339Nano
	LevelKey        string // default "level"
	TimeKey         string // default "time"
	MsgKey          string // default "msg"
	TraceIDKey      string // default "trace_id"
}

func NewLogger() *logrus.Logger {
	l := logrus.New()
	l.SetLevel(GetLogLevel())
	l.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})

	l.SetFormatter(&OrderedJSONFormatter{
		PadLevelTo:      5,
		TimestampFormat: "2006-01-02T15:04:05.000000Z07:00",
		LevelKey:        "level",
		TimeKey:         "time",
		MsgKey:          "msg",
		TraceIDKey:      "trace_id",
	})

	return l
}

func (f *OrderedJSONFormatter) Format(e *logrus.Entry) ([]byte, error) {
	padTo := f.PadLevelTo
	if padTo <= 0 {
		padTo = 5
	}
	tsFmt := f.TimestampFormat
	if tsFmt == "" {
		tsFmt = time.RFC3339Nano
	}

	msgKey := keyOr(f.MsgKey, "msg")
	timeKey := keyOr(f.TimeKey, "time")
	levelKey := keyOr(f.LevelKey, "level")
	traceKey := keyOr(f.TraceIDKey, "trace_id")

	lvl := e.Level.String()
	if n := padTo - len(lvl); n > 0 {
		lvl = lvl + strings.Repeat(" ", n)
	}

	trace := ""
	if v, ok := e.Data[traceKey]; ok {
		trace = fmt.Sprint(v)
	}

	buf := &bytes.Buffer{}
	buf.Grow(256)
	buf.WriteByte('{')

	writeKV(buf, levelKey, lvl, true)
	writeKV(buf, timeKey, e.Time.Format(tsFmt), false)
	if trace != "" {
		writeKV(buf, traceKey, trace, false)
	}
	writeKV(buf, msgKey, e.Message, false)

	if len(e.Data) > 0 {
		keys := make([]string, 0, len(e.Data))
		for k := range e.Data {
			if k == traceKey {
				continue
			}
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			buf.WriteByte(',')
			writeKey(buf, k)
			buf.WriteByte(':')
			valBytes, _ := json.Marshal(e.Data[k])
			buf.Write(valBytes)
		}
	}

	buf.WriteByte('}')
	buf.WriteByte('\n')
	return buf.Bytes(), nil
}

func keyOr(v, def string) string {
	if v == "" {
		return def
	}
	return v
}

func writeKV(buf *bytes.Buffer, k, v string, first bool) {
	if !first {
		buf.WriteByte(',')
	}
	writeKey(buf, k)
	buf.WriteByte(':')
	b, _ := json.Marshal(v)
	buf.Write(b)
}

func writeKey(buf *bytes.Buffer, k string) {
	b, _ := json.Marshal(k)
	buf.Write(b)
}
