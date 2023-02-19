package handler

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/fraqtop/footballcore/competition"
	"github.com/labstack/gommon/log"
)

type CompetitionConsumerHandler struct {
	competition.WriteRepository
}

func (this CompetitionConsumerHandler) Setup(session sarama.ConsumerGroupSession) error {
	log.Info("starting consumer group of competitions sync")
	return nil
}

func (this CompetitionConsumerHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	log.Info("exiting consumer group of competitions sync")
	return nil
}

func (this CompetitionConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	var (
		competitions []competition.Competition
		err error
	)

	for message := range claim.Messages() {
		log.Info("handling new message", message)
		err = json.Unmarshal(message.Value, &competitions)
		if err != nil {
			log.Warnf("cant't update competitions, %s", err.Error())
		} else {
			for _, entity := range competitions {
				err = this.WriteRepository.Save(entity)
				if err != nil {
					return err
				}
			}
			log.Info("successful")
		}
		session.MarkMessage(message, "")
	}

	return nil
}
