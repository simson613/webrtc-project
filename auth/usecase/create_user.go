package usecase

import (
	"fmt"
	"github/simson613/webrtc-project/auth/dto"
)

func (uc *Usecase) CreateUser(param *dto.SubscribeCreateUser) {
	utilErr := uc.MongoDBTransactionHandler(param)
	if utilErr != nil {
		// fail publish
		fmt.Println(utilErr.Error.Error())
	}
}
