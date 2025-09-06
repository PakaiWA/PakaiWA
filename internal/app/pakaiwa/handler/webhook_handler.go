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
 * @author KAnggara75 on Sun 31/08/25 13.46
 * @project PakaiWA webhooks
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/webhooks
 */

package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/model"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/helper"
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

func ProcessMessageEvent(msg *waE2E.Message, info types.MessageInfo, log *logrus.Logger) {
	text, msgType, raw := extractMessageTextAndType(msg)
	payload := model.MessageEventPayload{
		Event:       "message",
		From:        helper.NormalizeNumber(info.Sender.String()),
		Chat:        info.Chat.String(),
		PushName:    info.PushName,
		Timestamp:   info.Timestamp.Unix(),
		Text:        text,
		MessageType: msgType,
		Raw:         raw,
	}

	if msgType != "unknown" {
		go postJSON(context.Background(), webhookURL, payload, log)
	}

}

func postJSON(ctx context.Context, url string, v any, log *logrus.Logger) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	b, err := json.Marshal(v)
	if err != nil {
		log.WithError(err).Error("marshal webhook payload")
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		log.WithError(err).Error("create webhook request")
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
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

func extractMessageTextAndType(m *waE2E.Message) (text string, msgType string, raw *model.RawMsg) {
	if m == nil {
		return "", "unknown", &model.RawMsg{}
	}
	r := &model.RawMsg{}

	switch {
	case m.GetConversation() != "":
		text = m.GetConversation()
		msgType = "text"
		r.Conversation = text

	case m.ExtendedTextMessage != nil && m.ExtendedTextMessage.Text != nil:
		text = m.GetExtendedTextMessage().GetText()
		msgType = "extended_text"
		r.ExtendedText = text

	case m.ImageMessage != nil:
		text = m.GetImageMessage().GetCaption()
		msgType = "image"
		r.ImageCaption = text

	case m.DocumentMessage != nil:
		text = m.GetDocumentMessage().GetCaption()
		msgType = "document"
		r.DocumentCaption = text

	case m.StickerMessage != nil:
		text = m.StickerMessage.String()
		msgType = "sticker"
		r.StickerEmoji = text

	case m.AudioMessage != nil:
		text = m.AudioMessage.String()
		if m.AudioMessage.PTT != nil && m.AudioMessage.GetPTT() {
			msgType = "audio_ptt"
			r.AudioPTT = true
		} else {
			msgType = "audio"
		}

	case m.ButtonsMessage != nil:
		if m.GetButtonsMessage().ContentText != nil {
			text = m.GetButtonsMessage().GetContentText()
			r.ButtonsMessageText = text
		}
		msgType = "buttons_message"

	case m.TemplateMessage != nil && m.GetTemplateMessage().HydratedTemplate != nil:
		ht := m.GetTemplateMessage().GetHydratedTemplate()
		if ht.GetHydratedContentText() != "" {
			text = ht.GetHydratedContentText()
			r.TemplateMessageText = text
		}
		msgType = "template_message"

	default:
		msgType = "unknown"
	}

	return text, msgType, r
}
