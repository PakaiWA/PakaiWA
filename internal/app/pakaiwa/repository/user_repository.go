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

	"github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/model"
	"github.com/PakaiWA/PakaiWA/internal/pkg/logger/ctxmeta"
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
	log := ctxmeta.Logger(ctx)

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		if log != nil {
			log.WithError(err).Error("failed to begin transaction")
		}
		return nil, err
	}
	defer tx.Rollback(ctx)

	u := new(model.User)

	err = tx.QueryRow(ctx, `
		INSERT INTO pakaiwa.users (email, password)
		VALUES ($1, $2)
		RETURNING id
	`, email, hashedPassword).Scan(&u.ID)

	if err != nil {
		if log != nil {
			log.WithError(err).Error("failed to insert user")
		}
		return nil, err
	}

	_, err = tx.Exec(ctx, `INSERT INTO pakaiwa.user_quotas (user_id) VALUES ($1)`, u.ID)

	if err != nil {
		if log != nil {
			log.WithError(err).WithField("user_id", u.ID).Error("failed to insert default user quota")
		}
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		if log != nil {
			log.WithError(err).WithField("user_id", u.ID).Error("failed to commit transaction")
		}
		return nil, err
	}

	if log != nil {
		log.WithField("user_id", u.ID).Info("user created with default quota")
	}

	return u, nil
}

func (r *userRepository) GetUserByEmail(
	ctx context.Context,
	email string,
) (*model.User, error) {
	log := ctxmeta.Logger(ctx)
	log.Infof("GetUserByEmail")

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
