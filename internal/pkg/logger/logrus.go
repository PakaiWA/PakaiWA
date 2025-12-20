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
 * @author KAnggara75 on Sat 06/09/25 11.06
 * @project PakaiWA logger
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/pkg/logger
 */

package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/PakaiWA/PakaiWA/internal/pkg/config"
)

type OrderedJSONFormatter struct {
	PadLevelTo      int
	TimestampFormat string // default RFC3339Nano
	LevelKey        string // default "level"
	TimeKey         string // default "time"
	MsgKey          string // default "msg"
	TraceIDKey      string // default "trace_id"
	EscapeHTML      bool
}

func NewLogger() *logrus.Logger {
	l := logrus.New()
	l.SetLevel(config.GetLogLevel())
	l.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})

	l.SetFormatter(&OrderedJSONFormatter{
		PadLevelTo:      5,
		TimestampFormat: time.RFC3339Nano,
		LevelKey:        "level",
		TimeKey:         "time",
		MsgKey:          "msg",
		TraceIDKey:      "trace_id",
		EscapeHTML:      false,
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

	lvl := normalizeLevel(e.Level)
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

	writeKV(buf, levelKey, lvl, true, f.EscapeHTML)
	writeKV(buf, timeKey, e.Time.Format(tsFmt), false, f.EscapeHTML)
	if trace != "" {
		writeKV(buf, traceKey, trace, false, f.EscapeHTML)
	}
	writeKV(buf, msgKey, e.Message, false, f.EscapeHTML)

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
			writeKey(buf, k, f.EscapeHTML)
			buf.WriteByte(':')

			// marshal value dengan SetEscapeHTML(f.EscapeHTML)
			var vb bytes.Buffer
			enc := json.NewEncoder(&vb)
			enc.SetEscapeHTML(f.EscapeHTML)
			_ = enc.Encode(e.Data[k])
			val := vb.Bytes()
			if n := len(val); n > 0 && val[n-1] == '\n' {
				val = val[:n-1]
			}
			buf.Write(val)
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

func writeKey(buf *bytes.Buffer, k string, escapeHTML bool) {
	writeJSONString(buf, k, escapeHTML)
}

func writeKV(buf *bytes.Buffer, k, v string, first bool, escapeHTML bool) {
	if !first {
		buf.WriteByte(',')
	}
	writeKey(buf, k, escapeHTML)
	buf.WriteByte(':')
	writeJSONString(buf, v, escapeHTML)
}

func writeJSONString(buf *bytes.Buffer, s string, escapeHTML bool) {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(escapeHTML)
	if err := enc.Encode(s); err != nil {
		b.WriteString(`""`)
	}
	out := b.Bytes()
	if n := len(out); n > 0 && out[n-1] == '\n' {
		out = out[:len(out)-1]
	}
	buf.Write(out)
}

func normalizeLevel(level logrus.Level) string {
	switch level {
	case logrus.WarnLevel:
		return "warn"
	case logrus.ErrorLevel:
		return "error"
	case logrus.FatalLevel:
		return "fatal"
	case logrus.PanicLevel:
		return "panic"
	case logrus.InfoLevel:
		return "info"
	case logrus.DebugLevel:
		return "debug"
	case logrus.TraceLevel:
		return "trace"
	default:
		return level.String()
	}
}
