package main

import (
	"github.com/fraqtop/footballapi/cmd/worker/common"
	"github.com/fraqtop/footballapi/cmd/worker/competition/handler"
	"github.com/fraqtop/footballapi/internal/config"
	"github.com/fraqtop/footballapi/internal/connection"
	"github.com/fraqtop/footballapi/internal/repository/competition"
	"log"
)

const (
	topic = "api.competition.sync"
	groupId = "competition.sync"
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

	competitionWriteRepository := competition.NewWriteRepository(connectionEntity)
	consumerHandler := handler.CompetitionConsumerHandler{
		WriteRepository: competitionWriteRepository,
	}

	worker := common.GetConsumingWorker(consumerHandler, topic, groupId, config.GetBrokerConfig())
	worker.Start()
}
