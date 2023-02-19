package main

import (
	"github.com/fraqtop/footballapi/cmd/worker/common"
	"github.com/fraqtop/footballapi/cmd/worker/stats/handler"
	"github.com/fraqtop/footballapi/internal/config"
	"github.com/fraqtop/footballapi/internal/connection"
	"github.com/fraqtop/footballapi/internal/repository/competition"
	"github.com/fraqtop/footballapi/internal/repository/stats"
	"github.com/fraqtop/footballapi/internal/repository/team"
	"log"
)

const (
	topic = "api.stats.sync"
	groupId = "stats.sync"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	connectionEntity, err := connection.GetStorage()
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Destroy()

	teamWriteRepository := team.NewWriteRepository(connectionEntity)
	competitionWriteRepository := competition.NewWriteRepository(connectionEntity)
	statsWriteRepository := stats.NewWriteRepository(teamWriteRepository, competitionWriteRepository, connectionEntity)
	consumerHandler := handler.StatsConsumerHandler{
		WriteRepository: statsWriteRepository,
	}

	worker := common.GetConsumingWorker(consumerHandler, topic, groupId, config.GetBrokerConfig())
	worker.Start()
}
