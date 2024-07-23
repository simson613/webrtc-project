package consumer

import (
	"github/simson613/webrtc-project/auth/config"
	"github/simson613/webrtc-project/auth/usecase"

	"github.com/IBM/sarama"
)

type Consumer struct {
	config config.ConfigInterface
	uc     *usecase.Usecase
	client *sarama.Client
}

func InitConsumer(config config.ConfigInterface,
	uc *usecase.Usecase) *Consumer {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.ChannelBufferSize = 1000000
	kafkaConfig.Consumer.Return.Errors = true
	kafkaConfig.Consumer.Offsets.Initial = sarama.OffsetNewest
	kafkaConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange

	kafkaClient, err := sarama.NewClient(config.Kafka().Addr(), kafkaConfig)
	if err != nil {
		panic(err)
	}

	return &Consumer{
		config: config,
		uc:     uc,
		client: &kafkaClient,
	}
}

func (c *Consumer) Listener() {
	c.CreateUser()
}
