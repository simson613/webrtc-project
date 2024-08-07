package command

import (
	"github/simson613/webrtc-project/auth/dto"
	"github/simson613/webrtc-project/auth/util"
	"net/http"
)

func (c *Command) CheckLoginToken(strToken string) (*dto.LoginTokenInfo, *util.Error) {
	tokenInfo, err := c.loginTokenValidCheck(strToken)
	if err != nil {
		return nil, util.ErrorHandle(http.StatusUnauthorized, err)
	}

	user := c.extractionLoginToken(tokenInfo)

	return user, nil
}
