package container

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/fraqtop/footballapi/internal/config"
	"github.com/fraqtop/footballapi/internal/connection"
	"github.com/fraqtop/footballapi/internal/output"
	"github.com/fraqtop/footballapi/internal/repository/competition"
	"github.com/fraqtop/footballapi/internal/repository/stats"
	"github.com/fraqtop/footballapi/internal/repository/team"
	competitioncore "github.com/fraqtop/footballcore/competition"
	"github.com/redis/go-redis/v9"
	"go.uber.org/dig"
	"log"
	"time"
)

var containerInstance *dig.Container

func Init() {
	containerInstance = dig.New()
	registerDependencies(containerInstance)
}

func Get() *dig.Container {
	if containerInstance == nil {
		Init()
	}

	return containerInstance
}

func registerDependencies(container *dig.Container) {
	_ = container.Provide(config.GetBrokerConfig)
	_ = container.Provide(config.GetCacheConfig)
	_ = container.Provide(config.GetServerConfig)
	_ = container.Provide(config.GetStorageConfig)
	_ = container.Provide(provideConnection)
	_ = container.Provide(provideCacheClient)
	_ = container.Provide(provideCompetitionReadRepository)
	_ = container.Provide(provideCompetitionListFormatter)
	_ = container.Provide(competition.NewWriteRepository)
	_ = container.Provide(team.NewWriteRepository)
	_ = container.Provide(stats.NewWriteRepository)
}

func provideConnection() *sql.DB {
	storageConnection, err := connection.GetStorage()
	if err != nil {
		log.Fatal(err)
	}

	return storageConnection
}

func provideCacheClient(cacheConfig *config.CacheConfig) redis.UniversalClient {
	return connection.GetRedisClient(cacheConfig)
}

func provideCompetitionReadRepository(redisClient redis.UniversalClient, dbConnection *sql.DB) competitioncore.ReadRepository {
	return competition.NewRepositoryCacheProxy(
		context.Background(),
		redisClient,
		competition.NewReadRepository(dbConnection),
		time.Second*30,
	)
}

func provideCompetitionListFormatter() *output.CompetitionListFormatter {
	return &output.CompetitionListFormatter{}
}
