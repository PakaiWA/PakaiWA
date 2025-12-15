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

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/model"
	"github.com/PakaiWA/PakaiWA/internal/pkg/config"
)

type UserRepository interface {
	CreateUser(ctx context.Context, email, hashedPassword string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}

type userRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) UserRepository {
	return &userRepository{pool: pool}
}

func (r *userRepository) CreateUser(
	ctx context.Context,
	email string,
	hashedPassword string,
) (*model.User, error) {

	const query = `
		INSERT INTO pakaiwa.users (email, password)
		VALUES ($1, $2)
		RETURNING id
	`

	u := new(model.User)

	err := r.pool.QueryRow(ctx, query, email, hashedPassword).Scan(&u.ID)
	log.Err(err).Str("trace_id", config.Get40Space()).Msg("Creating user in database")

	if err != nil {
		return nil, err
	}

	log.Info().Str("trace_id", config.Get40Space()).Msgf("User created with ID: %s", u.ID)

	return u, nil
}

func (r *userRepository) GetUserByEmail(
	ctx context.Context,
	email string,
) (*model.User, error) {

	const q = `
		SELECT id, email, password, logout_at
		FROM pakaiwa.users
		WHERE email = $1
	`

	u := new(model.User)

	err := r.pool.QueryRow(ctx, q, email).Scan(
		&u.ID,
		&u.Email,
		&u.Password,
		&u.LogoutAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return u, nil
}
