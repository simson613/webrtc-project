package config

import (
	"os"
	"strings"
)

type KafkaInterface interface {
	Addr() []string
}

type Kafka struct {
	addr []string
}

func initKafkaConfig() *Kafka {
	addr := os.Getenv("KAFKA_ADDR")
	if addr == "" {
		addr = "localhost:29091,localhost:29092,localhost:29093"
	}
	address := strings.Split(addr, ",")

	return &Kafka{
		addr: address,
	}
}

func (kafka *Kafka) Addr() []string {
	return kafka.addr
}
