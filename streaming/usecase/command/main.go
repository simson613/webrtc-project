package command

import "github/simson613/webrtc-project/streaming/config"

type Command struct {
	config config.ConfigInterface
}

func InitCommand(config config.ConfigInterface) *Command {
	return &Command{
		config: config,
	}
}
