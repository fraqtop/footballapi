package connection

import (
	"github.com/fraqtop/footballapi/internal/config"
	"github.com/redis/go-redis/v9"
)

var cacheConnection *redis.Client

func GetRedisClient(config *config.CacheConfig) *redis.Client {
	if cacheConnection == nil {
		cacheConnection = redis.NewClient(&redis.Options{
			Addr:     config.Host(),
			Password: config.Password(),
			DB:       0,
		})
	}

	return cacheConnection
}
