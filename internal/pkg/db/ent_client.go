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
 * @author KAnggara75 on Sun 09/11/25 16.40
 * @project PakaiWA db
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/pkg/db
 */

package db

import (
	"context"
	"database/sql"

	"entgo.io/ent/dialect"
	entSql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	logrus "github.com/sirupsen/logrus"

	"github.com/PakaiWA/PakaiWA/ent"
	"github.com/PakaiWA/PakaiWA/internal/pkg/config"
)

func NewEntClient(ctx context.Context, log *logrus.Logger) *ent.Client {
	db, err := sql.Open("pgx", config.GetDBConn())
	if err != nil {
		log.Fatalf("failed to open pgx driver: %v", err)
	}

	drv := entSql.OpenDB(dialect.Postgres, db)

	client := ent.NewClient(ent.Driver(drv))

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
