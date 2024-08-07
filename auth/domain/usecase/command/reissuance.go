package command

import (
	"fmt"
	"github/simson613/webrtc-project/auth/dto"
	"github/simson613/webrtc-project/auth/util"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

func (c *Command) RessuanceLogin(
	param *dto.ReadLoginTokenParam) (*dto.LoginToken, *util.Error) {
	// read refresh token in mongo
	tokenInfo, utilErr := c.readLoginToken(param)
	if utilErr != nil {
		return nil, utilErr
	}

	// read user info in mongo
	userParam := dto.ReadUserByIdParam{
		Id: tokenInfo.Id,
	}
	user, err := c.mongo.ReadUserById(&userParam)
	if err != nil {
		return nil, util.ErrorHandle(c.checkMongoError(err), err)
	}

	//create token
	tokenParam := dto.CreateLoginTokenParam{
		Key:  user.Key,
		Id:   tokenInfo.Id,
		Name: user.Name,
	}
	token, utilErr := c.createLoginToken(&tokenParam)
	if utilErr != nil {
		return nil, utilErr
	}

	// delete before refresh token in mongo
	deleteParam := dto.DeleteTokenId{Id: param.Id}
	if utilErr := c.DeleteRefreshToken(&deleteParam); utilErr != nil {
		return nil, utilErr
	}

	return token, nil
}

func (c *Command) readLoginToken(
	param *dto.ReadLoginTokenParam) (*dto.LoginTokenInfo, *util.Error) {
	// read refresh token in mongo
	refreshToken, err := c.mongo.ReadLoginToken(param)
	if err != nil {
		return nil, util.ErrorHandle(c.checkMongoError(err), err)
	}

	// check expiration
	if time.Now().After(refreshToken.Expiration) {
		return nil, util.ErrorHandle(http.StatusUnauthorized, fmt.Errorf("expiration refresh token"))
	}

	// check token valid
	tokenInfo, err := c.loginTokenValidCheck(refreshToken.Token)
	if err != nil {
		return nil, util.ErrorHandle(http.StatusUnauthorized, err)
	}

	// extraction token info
	user := c.extractionLoginToken(tokenInfo)

	return user, nil
}

func (c *Command) loginTokenValidCheck(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(c.config.Auth().TokenSecret()), nil
	})
	return token, err
}

func (c *Command) extractionLoginToken(
	token *jwt.Token) *dto.LoginTokenInfo {
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	userName := claims["user_name"].(string)

	return &dto.LoginTokenInfo{
		Id:   userId,
		Name: userName,
	}
}
