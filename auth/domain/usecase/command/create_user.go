package command

import (
	"fmt"
	"github/simson613/webrtc-project/auth/dto"
)

func (uc *Command) CreateUser(param *dto.SubscribeCreateUser) {
	_, utilErr := uc.MongoDBTransactionHandler(param)
	if utilErr != nil {
		// fail publish
		fmt.Println(utilErr.Error.Error())
	}
}
