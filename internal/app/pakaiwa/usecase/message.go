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
 * @author KAnggara75 on Sun 31/08/25 05.54
 * @project PakaiWA messages
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/messages
 */

package usecase

import (
	"context"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/model"
)

type MessageUsecase interface {
	SendMessage(ctx context.Context, req *model.SendMessageReq) (string, error)
	EditMessage(ctx context.Context, req *model.SendMessageReq, msgId string) error
}
