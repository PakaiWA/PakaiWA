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
 * @author KAnggara75 on Sun 09/11/25 16.20
 * @project PakaiWA repository
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/repository
 */

package repository

import (
	"context"

	"github.com/PakaiWA/PakaiWA/ent"
	"github.com/PakaiWA/PakaiWA/ent/user"
)

type UserRepository interface {
	CreateUser(ctx context.Context, email, hashedPassword string) (*ent.User, error)
	GetUserByEmail(ctx context.Context, email string) (*ent.User, error)
}

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(c *ent.Client) UserRepository {
	return &userRepository{client: c}
}

func (r *userRepository) CreateUser(ctx context.Context, email, hashedPassword string) (*ent.User, error) {
	return r.client.User.
		Create().
		SetEmail(email).
		SetPassword(hashedPassword).
		Save(ctx)
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	return r.client.User.
		Query().
		Where(user.EmailEQ(email)).
		Only(ctx)
}
