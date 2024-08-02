package usecase

import (
	"github/simson613/webrtc-project/auth/dto"
	"github/simson613/webrtc-project/auth/util"
	"net/http"
)

func (uc *Usecase) CheckLoginToken(strToken string) (*dto.LoginTokenInfo, *util.Error) {
	tokenInfo, err := uc.loginTokenValidCheck(strToken)
	if err != nil {
		return nil, util.ErrorHandle(http.StatusUnauthorized, err)
	}

	user := uc.extractionLoginToken(tokenInfo)

	return user, nil
}
