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
 * @author KAnggara75 on Sun 09/11/25 17.51
 * @project PakaiWA seed
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/pkg/seed
 */

package seed

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/PakaiWA/PakaiWA/ent"
)

func RunSeeders(ctx context.Context, client *ent.Client, log *logrus.Logger) {
	_ = clearDatabase(ctx, client)
	log.Info("running database seeders...")

	if err := seedUsers(ctx, client); err != nil {
		log.WithError(err).Fatal("failed seeding users")
	}

	log.Info("database seeding completed")
}

func clearDatabase(ctx context.Context, client *ent.Client) error {
	// Urutan penghapusan penting untuk hindari foreign key constraint error
	// Jadi mulai dari tabel yang paling bergantung (child), ke tabel utama (parent)

	if _, err := client.Permission.Delete().Exec(ctx); err != nil {
		return fmt.Errorf("failed to clear permissions: %w", err)
	}

	if _, err := client.User.Delete().Exec(ctx); err != nil {
		return fmt.Errorf("failed to clear users: %w", err)
	}

	return nil
}
