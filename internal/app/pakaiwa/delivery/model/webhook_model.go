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
 * @author KAnggara75 on Sun 31/08/25 14.08
 * @project PakaiWA webhooks
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/webhooks
 */
package model

import "time"

type IncomingMessageWebhook struct {
	ID          string                 `json:"id"`
	WebhookType string                 `json:"webhook_type"`
	Payload     IncomingMessagePayload `json:"payload"`
	CreatedAt   time.Time              `json:"created_at"`
	ServerTime  time.Time              `json:"server_time"`
}

type IncomingMessagePayload struct {
	ID               string              `json:"id"`
	DeviceID         string              `json:"device_id"`
	Sender           string              `json:"sender"`
	FromMe           bool                `json:"from_me"`
	MessageType      string              `json:"message_type"`
	Text             string              `json:"text"`
	Caption          *string             `json:"caption"`
	IsGroupMessage   bool                `json:"is_group_message"`
	GroupID          *string             `json:"group_id"`
	Timestamp        int64               `json:"timestamp"`
	Contact          ContactInfo         `json:"contact"`
	Location         LocationInfo        `json:"location"`
	Meta             MediaMeta           `json:"meta"`
	InteractiveReply InteractiveResponse `json:"interactive_response"`
}

type ContactInfo struct {
	Name  *string `json:"name"`
	VCard *string `json:"vcard"`
}

type LocationInfo struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      *string `json:"name"`
	Address   *string `json:"address"`
}

type MediaMeta struct {
	FileName  *string  `json:"file_name"`
	Seconds   int      `json:"seconds"`
	Size      int64    `json:"size"`
	MimeType  *string  `json:"mime_type"`
	Dimension MediaDim `json:"dimension"`
}

type MediaDim struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

type InteractiveResponse struct {
	ID    *string `json:"id"`
	Label *string `json:"label"`
}

type MessageEventPayload struct {
	Event       string  `json:"event"`
	From        string  `json:"from"`
	Chat        string  `json:"chat"`
	GroupId     string  `json:"chat_name,omitempty"`
	PushName    string  `json:"push_name,omitempty"`
	Timestamp   int64   `json:"timestamp"`
	Text        string  `json:"text,omitempty"`
	MessageType string  `json:"message_type"`
	Raw         *RawMsg `json:"raw,omitempty"`
}

type RawMsg struct {
	Conversation        string `json:"conversation,omitempty"`
	ExtendedText        string `json:"extended_text,omitempty"`
	ImageCaption        string `json:"image_caption,omitempty"`
	DocumentCaption     string `json:"document_caption,omitempty"`
	StickerEmoji        string `json:"sticker_emoji,omitempty"`
	AudioPTT            bool   `json:"audio_ptt,omitempty"`
	ButtonsMessageText  string `json:"buttons_message_text,omitempty"`
	TemplateMessageText string `json:"template_message_text,omitempty"`
}
