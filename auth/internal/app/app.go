package app

import (
	"context"

	"github.com/Markuysa/pkg/redis"
	"gitlab.com/coinhubs/balance/internal/config"
)

func Run(ctx context.Context, cfg *config.Config) error {
	redisConn, err := redis.New(cfg.Redis)
	if err != nil {
		return err
	}

	return nil
}
