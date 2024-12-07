package client

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/coinhubs/balance/internal/domain"
	"gitlab.com/coinhubs/balance/internal/storage/postgres"
)

type repository struct {
	pgConn postgres.Conn
}

func New(pool *pgxpool.Pool) *repository {
	return &repository{
		pgConn: postgres.NewConn(pool),
	}
}

func (r *repository) CreateUser(ctx context.Context, user domain.User) (created domain.User, err error) {
	pgUser, err := r.pgConn.Queries(ctx).CreateUser(ctx, postgres.CreateUserParams{
		Login: user.Login,
		Email: user.Email,
		CreatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		return created, fmt.Errorf("failed to create user: %w", err)
	}

	return userFromRepository(pgUser), nil
}
