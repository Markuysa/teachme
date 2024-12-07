package config

import (
	"github.com/Markuysa/pkg/logger"
	"github.com/Markuysa/pkg/redis"
	"github.com/Markuysa/pkg/srv/grpc"
)

type Config struct {
	Redis  redis.Config
	Logger logger.Config
	GRPC   grpc.Config
}
