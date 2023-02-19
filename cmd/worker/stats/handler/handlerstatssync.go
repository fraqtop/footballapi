package handler

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/fraqtop/footballcore/stats"
	"github.com/labstack/gommon/log"
)

type StatsConsumerHandler struct {
	stats.WriteRepository
}

func (this StatsConsumerHandler) Setup(session sarama.ConsumerGroupSession) error {
	log.Info("starting consumer group of stats sync")
	return nil
}

func (this StatsConsumerHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	log.Info("starting consumer group of stats sync")
	return nil
}

func (this StatsConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	var (
		statsCollection []stats.Stats
		err             error
	)

	for message := range claim.Messages() {
		log.Info("handling new message", message)
		err = json.Unmarshal(message.Value, &statsCollection)
		if err != nil {
			log.Warnf("can't sync stats: %s", err.Error())
		} else {
			err = this.WriteRepository.BatchUpdate(statsCollection)
			if err != nil {
				return err
			}
			log.Info("successful")
		}
		session.MarkMessage(message, "")
	}

	return nil
}
