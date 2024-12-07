package config

import (
	"github.com/Markuysa/pkg/logger"
	"github.com/Markuysa/pkg/postgres"
)

type Config struct {
	Postgres postgres.PgxPoolCfg
	Logger   logger.Config
}
