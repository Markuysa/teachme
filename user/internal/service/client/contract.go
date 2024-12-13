package client

import (
	"context"

	"gitlab.com/coinhubs/balance/internal/domain"
)

type (
	repository interface {
		CreateUser(ctx context.Context, user domain.User) (domain.User, error)
		SaveEmailConfirmationCode(ctx context.Context, key string, code int) error
	}
	mailer interface {
		Send(ctx context.Context, to, subject, body string) error
	}
)
