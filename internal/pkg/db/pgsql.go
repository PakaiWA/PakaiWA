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
 * @author KAnggara75 on Sat 06/09/25 10.57
 * @project PakaiWA db
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/pkg/db
 */

package db

import (
	"context"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"

	"github.com/PakaiWA/PakaiWA/internal/pkg/config"
	"github.com/PakaiWA/PakaiWA/internal/pkg/utils"
)

var (
	pool   *pgxpool.Pool
	onceDb sync.Once
)

func NewDatabase(ctx context.Context, log *logrus.Logger) *pgxpool.Pool {
	traceID := config.Get40Space()
	log.WithField("trace_id", traceID).Info("Connecting to database...")

	onceDb.Do(func() {
		cfg, err := pgxpool.ParseConfig(config.GetDBConn())
		utils.PanicIfError(err)

		cfg.MinConns = config.GetDBMinConn()
		cfg.MaxConns = config.GetDBMaxConn()
		cfg.MaxConnIdleTime = config.DetMaxConnIdleTime()
		cfg.HealthCheckPeriod = config.GetDBHealthCheckPeriod()
		cfg.ConnConfig.ConnectTimeout = config.GetConnectTimeout()

		start := time.Now()
		pool, err = pgxpool.NewWithConfig(ctx, cfg)
		utils.PanicIfError(err)
		log.WithField("trace_id", traceID).Debugf("pgxpool took %s", time.Since(start))

		pingCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
		log.WithField("trace_id", traceID).Info("Pinging database...")
		if err := pool.Ping(pingCtx); err != nil {
			log.WithField("trace_id", traceID).WithError(err).Fatal("database ping failed")
		}
		log.WithField("trace_id", traceID).Infof("Pinging done in %s", time.Since(start))
	})

	if pool == nil {
		log.WithField("trace_id", traceID).Fatal("database pool is nil")
	} else {
		log.WithField("trace_id", traceID).Info("Connected to database...")
	}

	return pool
}
