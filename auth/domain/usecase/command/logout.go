package command

import (
	"github/simson613/webrtc-project/auth/dto"
	"github/simson613/webrtc-project/auth/util"
)

func (uc *Command) DeleteRefreshToken(
	param *dto.DeleteTokenId) *util.Error {
	// delete refresh token in mongodb
	_, utilErr := uc.MongoDBTransactionHandler(param)
	return utilErr
}
