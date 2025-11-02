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
 * @author KAnggara75 on Sat 06/09/25 14.05
 * @project PakaiWA store
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/pkg/store
 */

package store

import (
	"context"

	"github.com/PakaiWA/whatsmeow/store/sqlstore"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/sirupsen/logrus"

	"github.com/PakaiWA/PakaiWA/internal/pkg/logger"
)

func InitStoreWithPool(ctx context.Context, pool *pgxpool.Pool, log *logrus.Logger) *sqlstore.Container {
	db := stdlib.OpenDBFromPool(pool)
	dbLog := logger.NewPakaiWALog(log, "PakaiWA_DB")
	container := sqlstore.NewWithDB(db, "postgres", dbLog)

	if err := container.Upgrade(ctx); err != nil {
		panic(err)
	}
	return container
}
