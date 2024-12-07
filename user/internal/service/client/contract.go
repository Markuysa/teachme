package client

import (
	"context"

	"gitlab.com/coinhubs/balance/internal/domain"
)

type repository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
}
