package app

import (
	"context"

	"github.com/Markuysa/pkg/postgres"
	"gitlab.com/coinhubs/balance/internal/config"
)

func Run(ctx context.Context, cfg *config.Config) error {
	pgConn, err := postgres.New(cfg.Postgres)
	if err != nil {
		return err
	}

	return nil
}
