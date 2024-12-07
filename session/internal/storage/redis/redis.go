package redis

import (
	"context"
	"fmt"
	"time"

	v1 "session/pkg/api/grpc/v1"

	"github.com/Markuysa/pkg/tracer"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
)

const (
	sessionKeyfmt = "session:%s:%s"
)

type cache struct {
	rd *redis.Client
}

func NewStorage(rd *redis.Client) *cache {
	return &cache{rd: rd}
}

func (c *cache) SetClientSession(ctx context.Context, session *v1.Session) error {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	data, err := proto.Marshal(session)
	if err != nil {
		return err
	}

	key := fmt.Sprintf(sessionKeyfmt, session.GetAccessToken(), session.GetUserId())

	return c.rd.Set(
		ctx,
		key,
		data,
		time.Duration(session.ExpireIn),
	).Err()
}

func (c *cache) FindSessionByAccessToken(ctx context.Context, accessToken string) (*v1.Session, error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	result, err := c.rd.Eval(
		ctx,
		fetchUserSessionScript,
		[]string{},
		accessToken,
	).Result()
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, fmt.Errorf("session not found")
	}

	session := &v1.Session{}
	if err := proto.Unmarshal([]byte(result.(string)), session); err != nil {
		return nil, err
	}

	return session, nil
}
