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
 * @author KAnggara75 on Sun 07/09/25 19.32
 * @project PakaiWA model
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/delivery/model
 */

package model

import (
	"encoding/json"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/encoding/protojson"
	"time"
)

type IncomingMessageModel struct {
	Message   map[string]any    `json:"message"`
	Info      types.MessageInfo `json:"info"`
	Raw       map[string]any    `json:"raw"`
	CreatedAt time.Time         `json:"created_at"`
}

func (a *IncomingMessageModel) GetId() string {
	return ""
}

func ToIncomingMessageModel(msg *waE2E.Message, info types.MessageInfo, raw *waE2E.Message) (*IncomingMessageModel, error) {
	marshaler := protojson.MarshalOptions{
		EmitUnpopulated: false,
		UseEnumNumbers:  false,
	}

	// convert msg ke JSON
	msgJSON, err := marshaler.Marshal(msg)
	if err != nil {
		return nil, err
	}

	rawJSON, err := marshaler.Marshal(raw)
	if err != nil {
		return nil, err
	}

	// decode ke map[string]any biar fleksibel
	var msgMap map[string]any
	var rawMap map[string]any
	_ = json.Unmarshal(msgJSON, &msgMap)
	_ = json.Unmarshal(rawJSON, &rawMap)

	return &IncomingMessageModel{
		Message:   msgMap,
		Info:      info,
		Raw:       rawMap,
		CreatedAt: time.Now(),
	}, nil
}
