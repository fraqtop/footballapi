package config

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	storageConfig *StorageConfig
	serverConfig  *ServerConfig
)

func Load() error {
	return godotenv.Load()
}

func GetStorageConfig() *StorageConfig {
	if storageConfig == nil {
		storageConfig = &StorageConfig{
			host:     os.Getenv("DATABASE_HOST"),
			port:     os.Getenv("DATABASE_EXPOSE_PORT"),
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
			port: os.Getenv("SERVER_PORT"),
		}
	}

	return serverConfig
}
