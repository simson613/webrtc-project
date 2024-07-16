package usecase

import "github/simson613/webrtc-project/streaming/config"

type Usecase struct {
	config config.ConfigInterface
}

func InitUsecase(config config.ConfigInterface) *Usecase {
	return &Usecase{
		config: config,
	}
}
