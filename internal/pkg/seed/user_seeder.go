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

	"github.com/PakaiWA/PakaiWA/ent"
)

func seedUsers(ctx context.Context, client *ent.Client) error {
	count, err := client.User.Query().Count(ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return nil // sudah terisi, lewati
	}

	user, _ := client.User.Create().
		SetEmail("admin@example.com").
		SetPassword("5f4dcc3b5aa765d61d8327deb882cf99").
		Save(ctx)

	// Tambahkan permission untuk user tersebut
	_, _ = client.Permission.Create().
		SetPath("/v1/messages").
		SetMethod("POST").
		SetAccess("write").
		SetUser(user).
		Save(ctx)
	return err
}
