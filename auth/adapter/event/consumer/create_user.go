package consumer

import (
	"encoding/json"
	"fmt"
	"github/simson613/webrtc-project/auth/dto"
	"log"

	"github.com/IBM/sarama"
)

func (c *Consumer) CreateUser() {
	topic := "create-user"

	// consumerGroup, err := sarama.NewConsumerGroupFromClient(topic, *c.client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// consumerGroup.

	consumer, err := sarama.NewConsumerFromClient(*c.client)
	if err != nil {
		log.Fatal(err)
	}

	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		log.Fatal(err)
	}

	for _, partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			for {
				select {
				case err := <-pc.Errors():
					if err != nil {
						fmt.Printf("err %v \n", err)
					}
				case message := <-pc.Messages():
					if message != nil {
						fmt.Printf("Get Message %s\n", string(message.Value))
						msg := dto.SubscribeCreateUser{}
						if err := json.Unmarshal(message.Value, &msg); err != nil {
							fmt.Println(err)
						}
						c.command.CreateUser(&msg)
					}
				}
			}
		}()
	}
}
