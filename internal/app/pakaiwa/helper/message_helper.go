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
 * @author KAnggara75 on Sun 07/09/25 08.34
 * @project PakaiWA helper
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/helper
 */

package helper

import (
	"github.com/PakaiWA/whatsmeow/proto/waE2E"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/model"
)

func ExtractMessageTextAndType(m *waE2E.Message) (text string, msgType string, raw *model.RawMsg) {
	if m == nil {
		return "", "unknown", &model.RawMsg{}
	}
	r := &model.RawMsg{}

	switch {
	case m.GetConversation() != "":
		text = m.GetConversation()
		msgType = "text"
		r.Conversation = text

	case m.ReactionMessage != nil:
		text = m.GetReactionMessage().GetText()
		msgType = "reaction"
		r.ExtendedText = text

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
