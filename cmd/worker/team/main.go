package main

import (
	"github.com/fraqtop/footballapi/cmd/worker/common"
	"github.com/fraqtop/footballapi/cmd/worker/team/handler"
	"github.com/fraqtop/footballapi/internal/config"
	"github.com/fraqtop/footballapi/internal/container"
	"github.com/fraqtop/footballcore/team"
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

	err = container.Get().Invoke(func(repository team.WriteRepository) {
		consumerHandler := handler.TeamConsumerHandler{
			WriteRepository: repository,
		}

		worker := common.GetConsumingWorker(consumerHandler, topic, groupId, config.GetBrokerConfig())
		worker.Start()
	})

	if err != nil {
		log.Fatal(err)
	}
}
