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
 * @author KAnggara75 on Sun 07/09/25 08.04
 * @project PakaiWA usecase
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/usecase
 */

package usecase

import (
	"context"
	"fmt"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/helper"
	"github.com/PakaiWA/PakaiWA/internal/pkg/httpclient"
	"github.com/sirupsen/logrus"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"net/http"
	"time"
)

var (
	webhookURL = "https://pakaiwa.requestcatcher.com/"
	httpClient = &http.Client{Timeout: 5 * time.Second}
)

type MessageEventUseCase struct {
}

func ProcessMessageEvent(msg *waE2E.Message, info types.MessageInfo, log *logrus.Logger) {
	text, msgType, _ := helper.ExtractMessageTextAndType(msg)
	//payload := model.MessageEventPayload{
	//	Event:       "message",
	//	From:        helper.NormalizeNumber(info.Sender.String()),
	//	Chat:        info.Chat.String(),
	//	PushName:    info.PushName,
	//	Timestamp:   info.Timestamp.Unix(),
	//	Text:        text,
	//	MessageType: msgType,
	//	Raw:         raw,
	//}
	//
	fmt.Printf("TYPE %s", msgType)
	fmt.Printf("text %s", text)

	//if msgType != "unknown" {
	postJSON(context.Background(), webhookURL, msg, log)
	postJSON(context.Background(), webhookURL, info, log)
	//}

}

func postJSON(ctx context.Context, url string, v any, log *logrus.Logger) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	resp, err := httpclient.Post(ctx, url, v)
	if err != nil {
		log.WithError(err).Error("send webhook request")
		return
	}
	_ = resp.Body.Close()

	if resp.StatusCode >= 300 {
		log.WithField("status", resp.Status).Warn("webhook non-2xx")
	} else {
		log.Info("webhook sent successfully")
	}
}
