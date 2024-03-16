// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: user_management.write.sql

package xsqlc

import (
	"context"
	"time"
)

const saveActivationToken = `-- name: SaveActivationToken :one
INSERT INTO activation_tokens (id, token, expired_at, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (id) DO UPDATE SET 
    token = EXCLUDED.token,
    expired_at = EXCLUDED.expired_at,
    created_at = EXCLUDED.created_at,
    updated_at = EXCLUDED.updated_at
RETURNING id
`

type SaveActivationTokenParams struct {
	ID        string
	Token     string
	ExpiredAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) SaveActivationToken(ctx context.Context, arg SaveActivationTokenParams) (string, error) {
	row := q.db.QueryRowContext(ctx, saveActivationToken,
		arg.ID,
		arg.Token,
		arg.ExpiredAt,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}

const saveRelUserToActivationToken = `-- name: SaveRelUserToActivationToken :one
INSERT INTO rel_user_to_activation_tokens (user_id, activation_token_id)
VALUES ($1, $2)
ON CONFLICT (user_id, activation_token_id) DO NOTHING
RETURNING user_id, activation_token_id
`

type SaveRelUserToActivationTokenParams struct {
	UserID            string
	ActivationTokenID string
}

type SaveRelUserToActivationTokenRow struct {
	UserID            string
	ActivationTokenID string
}

func (q *Queries) SaveRelUserToActivationToken(ctx context.Context, arg SaveRelUserToActivationTokenParams) (SaveRelUserToActivationTokenRow, error) {
	row := q.db.QueryRowContext(ctx, saveRelUserToActivationToken, arg.UserID, arg.ActivationTokenID)
	var i SaveRelUserToActivationTokenRow
	err := row.Scan(&i.UserID, &i.ActivationTokenID)
	return i, err
}

const saveUser = `-- name: SaveUser :one
INSERT INTO users (id, "name", email, "password", "role", active, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT (id) DO UPDATE SET 
    "name" = EXCLUDED."name",
    email = EXCLUDED.email,
    "password" = EXCLUDED."password",
    "role" = EXCLUDED."role",
    active = EXCLUDED.active,
    created_at = EXCLUDED.created_at,
    updated_at = EXCLUDED.updated_at
RETURNING id
`

type SaveUserParams struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Role      string
	Active    int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) SaveUser(ctx context.Context, arg SaveUserParams) (string, error) {
	row := q.db.QueryRowContext(ctx, saveUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.Role,
		arg.Active,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}
