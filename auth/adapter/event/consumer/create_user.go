package consumer

import (
	"encoding/json"
	"fmt"
	"github/simson613/webrtc-project/auth/dto"

	"github.com/IBM/sarama"
)

func (c *Consumer) CreateUser() {
	topic := "create-user"
	consumer, err := sarama.NewConsumerFromClient(*c.client)
	if err != nil {
		fmt.Printf("consumer err: %v", err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Println(err)
	}

	for _, partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			fmt.Println(err)
		}
		defer func() {
			if err := pc.Close(); err != nil {
				fmt.Println(err)
			}
		}()

		go func() {
			for {
				select {
				case err := <-pc.Errors():
					fmt.Printf("err %v \n", err)
				case message := <-pc.Messages():
					msg := dto.SubscribeCreateUser{}
					if err := json.Unmarshal(message.Value, &msg); err != nil {
						fmt.Println(err)
					}
					c.uc.CreateUser(&msg)
					// fmt.Printf("msg %v \n %s \n", msg, string(message.Value))
				}
			}
		}()
	}
}
