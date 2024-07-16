package config

type ConfigInterface interface {
	Server() ServerInterface
}

type Config struct {
	server ServerInterface
}

func InitConfig() ConfigInterface {
	return &Config{
		server: initServerConfig(),
	}
}
func (config *Config) Server() ServerInterface {
	return config.server
}
