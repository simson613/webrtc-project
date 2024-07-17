package config

type ConfigInterface interface {
	MariaDB() MariaDBInterface
	MongoDB() MongoDBInterface
	Server() ServerInterface
}

type Config struct {
	mariaDB MariaDBInterface
	mongoDB MongoDBInterface
	server  ServerInterface
}

func InitConfig() ConfigInterface {
	return &Config{
		mariaDB: initMariaDBConfig(),
		mongoDB: initMongoDBConfig(),
		server:  initServerConfig(),
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
