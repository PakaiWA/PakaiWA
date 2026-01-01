/*
 * Copyright (c) 2026 KAnggara
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * See <https://www.gnu.org/licenses/gpl-3.0.html>.
 *
 * @author KAnggara on Thursday 01/01/2026 21.14
 * @project PakaiWA
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/bootstrap
 */

package bootstrap

import (
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/http/handler"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/http/middleware"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/http/router"
	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/usecase"
)

func initMessageModule(b *AppContext) {
	// Message
	msgUsecase := usecase.NewMessageUsecase(b.Validate, b.PakaiWA.Client)
	msgHandler := handler.NewMessageHandler(msgUsecase)

	// authenticated
	auth := router.RegisterAuthGroup(b.Fiber)

	// =====================
	// User API v1
	// =====================
	v1 := auth.Group("/v1", middleware.QuotaMiddleware(b.Redis))
	v1.Post("/messages", msgHandler.SendMsg)
	v1.Patch("/messages/:msgId", msgHandler.EditMsg)
	v1.Delete("/messages/:msgId", msgHandler.DeleteMsg)
	v1.Get("/groups", msgHandler.SendMsg)
}
