package usecase

import (
	"github/simson613/webrtc-project/user/dto"
	"github/simson613/webrtc-project/user/util"
)

func (uc *Usecase) ReadUserInfo(param *dto.ReadUserInfoParam) (
	*dto.ReadUserInfo, *util.Error) {
	userInfo, err := uc.mongo.ReadUserInfo(param)
	if err != nil {
		return nil, util.ErrorHandle(uc.checkMongoError(err), err)
	}
	return userInfo, nil
}
