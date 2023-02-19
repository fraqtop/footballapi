package main

import (
	"github.com/fraqtop/footballapi/cmd/worker/common"
	"github.com/fraqtop/footballapi/cmd/worker/team/handler"
	"github.com/fraqtop/footballapi/internal/config"
	"github.com/fraqtop/footballapi/internal/connection"
	"github.com/fraqtop/footballapi/internal/repository/team"
	"log"
)

const (
	topic   = "api.team.sync"
	groupId = "team.sync"
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
	consumerHandler := handler.TeamConsumerHandler{
		WriteRepository: teamWriteRepository,
	}

	worker := common.GetConsumingWorker(consumerHandler, topic, groupId, config.GetBrokerConfig())
	worker.Start()
}
