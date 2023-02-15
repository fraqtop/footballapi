package competition

import (
	"context"
	"encoding/json"
	"github.com/fraqtop/footballcore/competition"
	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	KeyAll = "all"
)

type RepositoryCacheProxy struct {
	ctx context.Context
	redisClient redis.UniversalClient
	repository competition.ReadRepository
}

func (this RepositoryCacheProxy) All() []competition.Competition {
	bytes, err := this.redisClient.Get(this.ctx, KeyAll).Result()
	var result []competition.Competition
	if err != nil {
		result = this.repository.All()
		jsonContent, err := json.Marshal(result)
		if err != nil {
			log.Warnf("can't encode to json: %s", err)
		} else {
			this.redisClient.Set(this.ctx, KeyAll, jsonContent, time.Second * 30)
		}
	} else {
		err = json.Unmarshal([]byte(bytes), &result)
		if err != nil {
			log.Warnf("cant't decode from json: %s", err)
		}
	}

	return result
}

func NewRepositoryCacheProxy(
	ctx context.Context,
	redisClient redis.UniversalClient,
	repository competition.ReadRepository,
	) *RepositoryCacheProxy {
	return &RepositoryCacheProxy{
		redisClient: redisClient,
		ctx: ctx,
		repository: repository,
	}
}
