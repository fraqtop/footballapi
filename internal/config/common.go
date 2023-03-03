package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var (
	storageConfig *StorageConfig
	serverConfig  *ServerConfig
	cacheConfig   *CacheConfig
	brokerConfig  *BrokerConfig
)

func Load() error {
	return godotenv.Load()
}

func GetStorageConfig() *StorageConfig {
	if storageConfig == nil {
		storageConfig = &StorageConfig{
			host:     os.Getenv("DATABASE_HOST"),
			port:     os.Getenv("DATABASE_INNER_PORT"),
			user:     os.Getenv("POSTGRES_USER"),
			password: os.Getenv("POSTGRES_PASSWORD"),
			name:     os.Getenv("POSTGRES_DB"),
		}
	}

	return storageConfig
}

func GetServerConfig() *ServerConfig {
	if serverConfig == nil {
		serverConfig = &ServerConfig{
			port:  os.Getenv("SERVER_INNER_PORT"),
			debug: os.Getenv("SERVER_DEBUG") == "1",
		}
	}

	return serverConfig
}

func GetCacheConfig() *CacheConfig {
	if cacheConfig == nil {
		cacheConfig = &CacheConfig{
			host:     fmt.Sprintf("%s:%s", os.Getenv("CACHE_HOST"), os.Getenv("CACHE_PORT")),
			password: os.Getenv("CACHE_PASSWORD"),
			port:     os.Getenv("CACHE_PORT"),
		}
	}

	return cacheConfig
}

func GetBrokerConfig() *BrokerConfig {
	if brokerConfig == nil {
		brokerConfig = &BrokerConfig{
			host: os.Getenv("BROKER_HOST"),
		}
	}

	return brokerConfig
}
