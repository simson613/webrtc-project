package producer

import (
	"fmt"
	"github/simson613/webrtc-project/user/config"
	"github/simson613/webrtc-project/user/dto"

	"github.com/IBM/sarama"
)

type ProducerInterface interface {
	CreateUser(*dto.PublishCreateUser) error
}

type producer struct {
	config config.ConfigInterface
}

func InitProducer(config config.ConfigInterface) ProducerInterface {
	return &producer{
		config: config,
	}
}

func (p *producer) getProducer() sarama.AsyncProducer {
	var config *sarama.Config = sarama.NewConfig()
	// config.Producer.Return.Successes = true
	// config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	// config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	// config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	// config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms

	producer, err := sarama.NewAsyncProducer(p.config.Kafka().Addr(), config)
	if err != nil {
		fmt.Println("failed to create Producer", err)
		return nil
	}

	go func() {
		for err := range producer.Errors() {
			fmt.Println("Failed to write access log entry:", err)
		}
	}()

	return producer
}

func (p *producer) sendMessage(producer sarama.AsyncProducer, topic string, value string) {
	kafkaMessage := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(value),
	}

	producer.Input() <- kafkaMessage

	select {
	case success := <-producer.Successes():
		fmt.Println("Message produced:", success.Offset)
	case err := <-producer.Errors():
		fmt.Println("Failed to produce message:", err)
	}
}
