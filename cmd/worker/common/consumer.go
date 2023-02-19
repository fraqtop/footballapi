package common

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/fraqtop/footballapi/internal/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type ConsumingWorker struct {
	handler sarama.ConsumerGroupHandler
	topic string
	groupId string
	brokerConfig *config.BrokerConfig
}

var worker *ConsumingWorker

func GetConsumingWorker(
	handler sarama.ConsumerGroupHandler,
	topic,
	groupId string,
	brokerConfig *config.BrokerConfig,
	) *ConsumingWorker {
	if worker == nil {
		worker = &ConsumingWorker{
			handler: handler,
			topic: topic,
			groupId: groupId,
			brokerConfig: brokerConfig,
		}
	}

	return worker
}

func (this ConsumingWorker) Start() {
	interruptContext, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	consumerConfig := sarama.NewConfig()
	consumerConfig.Consumer.Offsets.Initial = sarama.OffsetNewest
	consumerGroup, err := sarama.NewConsumerGroup([]string{this.brokerConfig.Host()}, this.groupId, consumerConfig)
	if err != nil {
		log.Fatal(err)
	}
	err = consumerGroup.Consume(interruptContext, []string{this.topic}, this.handler)
	if err != nil {
		log.Fatal(err)
	}
	<-interruptContext.Done()
}
