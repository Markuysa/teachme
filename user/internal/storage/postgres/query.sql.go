// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
insert into "user" (id, login, email, created_at)
values ($1, $2, $3, $4)
    returning id, login, email, created_at
`

type CreateUserParams struct {
	ID        pgtype.UUID
	Login     string
	Email     string
	CreatedAt pgtype.Timestamp
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.Login,
		arg.Email,
		arg.CreatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Login,
		&i.Email,
		&i.CreatedAt,
	)
	return i, err
}
