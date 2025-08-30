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
 * @author KAnggara75 on Sat 30/08/25 18.00
 * @project PakaiWA configs
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/configs
 */

package configs

import (
	"context"
	"github.com/PakaiWA/PakaiWA/internal/helpers"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var (
	pool   *pgxpool.Pool
	onceDb sync.Once
)

func NewDatabase(ctx context.Context, log *logrus.Logger) *pgxpool.Pool {
	traceID := Get40Space()
	log.WithField("trace_id", traceID).Info("Connecting to database...")

	onceDb.Do(func() {
		cfg, err := pgxpool.ParseConfig(GetDBConn())
		helpers.PanicIfError(err)

		cfg.MinConns = GetDBMinConn()
		cfg.MaxConns = GetDBMaxConn()
		cfg.MaxConnIdleTime = DetMaxConnIdleTime()
		cfg.HealthCheckPeriod = GetDBHealthCheckPeriod()
		cfg.ConnConfig.ConnectTimeout = GetConnectTimeout()

		start := time.Now()
		pool, err = pgxpool.NewWithConfig(ctx, cfg)
		helpers.PanicIfError(err)
		log.WithField("trace_id", traceID).Debugf("pgxpool took %s", time.Since(start))

		ctx, cancel := context.WithTimeout(ctx, time.Minute)
		defer cancel()
		log.WithField("trace_id", traceID).Info("Pinging database...")
		if err := pool.Ping(ctx); err != nil {
			log.WithField("trace_id", traceID).Errorf("Ping timeout: %v", err)
		}
		log.WithField("trace_id", traceID).Info("Pinging done...")
	})

	if pool == nil {
		log.WithField("trace_id", traceID).Error("Database pool is nil")
	} else {
		log.WithField("trace_id", traceID).Info("Connected to database...")
	}

	return pool
}
