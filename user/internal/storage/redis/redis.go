package redis

import "github.com/redis/go-redis/v9"

type cache struct {
	rd *redis.Client
}

func NewCache(rd *redis.Client) *cache {
	return &cache{rd: rd}
}
