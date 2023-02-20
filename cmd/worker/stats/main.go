package main

import (
	"github.com/fraqtop/footballapi/cmd/worker/common"
	"github.com/fraqtop/footballapi/cmd/worker/stats/handler"
	"github.com/fraqtop/footballapi/internal/config"
	"github.com/fraqtop/footballapi/internal/container"
	"github.com/fraqtop/footballcore/stats"
	"log"
)

const (
	topic   = "api.stats.sync"
	groupId = "stats.sync"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = container.Get().Invoke(func(repository stats.WriteRepository) {
		consumerHandler := handler.StatsConsumerHandler{
			WriteRepository: repository,
		}

		worker := common.GetConsumingWorker(consumerHandler, topic, groupId, config.GetBrokerConfig())
		worker.Start()
	})

	if err != nil {
		log.Fatal(err)
	}
}
