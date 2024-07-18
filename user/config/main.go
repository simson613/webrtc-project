package config

type ConfigInterface interface {
	MariaDB() MariaDBInterface
	MongoDB() MongoDBInterface
	Server() ServerInterface
	Swagger() SwaggerInterface
}

type Config struct {
	mariaDB MariaDBInterface
	mongoDB MongoDBInterface
	server  ServerInterface
	swagger SwaggerInterface
}

func InitConfig() ConfigInterface {
	return &Config{
		mariaDB: initMariaDBConfig(),
		mongoDB: initMongoDBConfig(),
		server:  initServerConfig(),
		swagger: initSwaggerConfig(),
	}
}

func (config *Config) MariaDB() MariaDBInterface {
	return config.mariaDB
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
