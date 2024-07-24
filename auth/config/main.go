package config

type ConfigInterface interface {
	Auth() AuthInterface
	Cookie() CookieInterface
	Kafka() KafkaInterface
	MongoDB() MongoDBInterface
	Server() ServerInterface
	Swagger() SwaggerInterface
}

type Config struct {
	auth    AuthInterface
	cookie  CookieInterface
	kafka   KafkaInterface
	mongoDB MongoDBInterface
	server  ServerInterface
	swagger SwaggerInterface
}

func InitConfig() ConfigInterface {
	return &Config{
		auth:    initAuthConfig(),
		cookie:  initCookieConfig(),
		kafka:   initKafkaConfig(),
		mongoDB: initMongoDBConfig(),
		server:  initServerConfig(),
		swagger: initSwaggerConfig(),
	}
}

func (config *Config) Auth() AuthInterface {
	return config.auth
}

func (config *Config) Cookie() CookieInterface {
	return config.cookie
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
