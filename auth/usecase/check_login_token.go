package usecase

import (
	"github/simson613/webrtc-project/auth/util"
	"net/http"
)

func (uc *Usecase) CheckLoginToken(strToken string) *util.Error {
	_, err := uc.loginTokenValidCheck(strToken)
	if err != nil {
		return util.ErrorHandle(http.StatusUnauthorized, err)
	}
	return nil
}
