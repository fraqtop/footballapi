package handler

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/fraqtop/footballcore/team"
	"github.com/labstack/gommon/log"
)

type TeamConsumerHandler struct {
	team.WriteRepository
}

func (this TeamConsumerHandler) Setup(session sarama.ConsumerGroupSession) error {
	log.Info("starting consumer group of team sync")
	return nil
}

func (this TeamConsumerHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	log.Info("exiting consumer group of team sync")
	return nil
}

func (this TeamConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	var (
		teams []team.Team
		err   error
	)

	for message := range claim.Messages() {
		log.Info("handling new message", message.Value)
		err = json.Unmarshal(message.Value, &teams)
		if err != nil {
			log.Warnf("can't sync teams: %s", err.Error())
		} else {
			err = this.WriteRepository.BatchUpdate(teams)
			if err != nil {
				return err
			}
			log.Info("successful")
		}
		session.MarkMessage(message, "")
	}

	return nil
}
