// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package postgres

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID        pgtype.UUID
	Name      string
	Email     string
	CreatedAt pgtype.Timestamp
}
