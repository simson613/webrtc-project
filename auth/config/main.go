package config

type ConfigInterface interface {
	Kafka() KafkaInterface
	MongoDB() MongoDBInterface
	Server() ServerInterface
	Swagger() SwaggerInterface
}

type Config struct {
	kafka   KafkaInterface
	mongoDB MongoDBInterface
	server  ServerInterface
	swagger SwaggerInterface
}

func InitConfig() ConfigInterface {
	return &Config{
		kafka:   initKafkaConfig(),
		mongoDB: initMongoDBConfig(),
		server:  initServerConfig(),
		swagger: initSwaggerConfig(),
	}
}

func (config *Config) Kafka() KafkaInterface {
	return config.kafka
}

func (config *Config) MongoDB() MongoDBInterface {
	return config.mongoDB
}

func (config *Config) Server() ServerInterface {
	return config.server
}

func (config *Config) Swagger() SwaggerInterface {
	return config.swagger
}
