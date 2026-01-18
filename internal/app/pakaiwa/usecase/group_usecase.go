/*
 * Copyright (c) 2025 KAnggara
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * See <https://www.gnu.org/licenses/gpl-3.0.html>.
 *
 * @author KAnggara on Sunday 28/12/2025 23.21
 * @project PakaiWA
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/usecase
 */

package usecase

import (
	"context"

	"github.com/PakaiWA/whatsmeow"
	"github.com/PakaiWA/whatsmeow/types"
	"github.com/go-playground/validator/v10"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/apperror"
	"github.com/PakaiWA/PakaiWA/internal/pkg/logger/ctxmeta"
)

type GroupUsecase interface {
	GetGroups(ctx context.Context) ([]*types.GroupInfo, error)
}

type groupUsecase struct {
	Validate *validator.Validate
	WA       *whatsmeow.Client
}

func NewGroupUsecase(validate *validator.Validate, wa *whatsmeow.Client) GroupUsecase {
	return &groupUsecase{
		WA:       wa,
		Validate: validate,
	}
}

// GetGroups implements [GroupUsecase].
func (g *groupUsecase) GetGroups(ctx context.Context) ([]*types.GroupInfo, error) {
	log := ctxmeta.Logger(ctx)
	if !g.WA.IsConnected() {
		log.WithError(apperror.ErrWAClientNotConnected).WithField("event", "wa_disconnected").Error("precondition failed")
		return nil, apperror.ErrWAClientNotConnected
	}

	chats, err := g.WA.GetJoinedGroups(ctx)
	if err != nil {
		log.WithError(err).Error("failed to get chats from WA client")
		return nil, apperror.ErrFailedToGetGroups
	}

	return chats, nil
}
