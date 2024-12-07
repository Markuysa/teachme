package client

import (
	"context"

	v1 "session/pkg/api/grpc/v1"
)

type storage interface {
	SetClientSession(ctx context.Context, session *v1.Session) error
	FindSessionByAccessToken(ctx context.Context, accessToken string) (*v1.Session, error)
}
