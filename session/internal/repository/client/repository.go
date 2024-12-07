package client

import (
	"context"
	rd "session/internal/storage/redis"
	v1 "session/pkg/api/grpc/v1"

	"github.com/Markuysa/pkg/tracer"
	"github.com/redis/go-redis/v9"
)

type repo struct {
	storage storage
}

func New(client *redis.Client) *repo {
	return &repo{
		storage: rd.NewStorage(client),
	}
}

func (r *repo) SetClientSession(ctx context.Context, session *v1.Session) error {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	return nil
}

func (r *repo) FindSessionByAccessToken(ctx context.Context, accessToken string) (*v1.Session, error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	return nil, nil
}
