package usecase

import (
	"fmt"
	"github/simson613/webrtc-project/auth/dto"
	"github/simson613/webrtc-project/auth/util"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

func (uc *Usecase) RessuanceLogin(
	param *dto.ReadLoginTokenParam) (*dto.LoginToken, *util.Error) {
	// read refresh token in mongo
	tokenInfo, utilErr := uc.readLoginToken(param)
	if utilErr != nil {
		return nil, utilErr
	}

	// read user info in mongo
	userParam := dto.ReadUserByIdParam{
		Id: tokenInfo.Id,
	}
	user, err := uc.mongo.ReadUserById(&userParam)
	if err != nil {
		return nil, util.ErrorHandle(uc.checkMongoError(err), err)
	}

	//create token
	tokenParam := dto.CreateLoginTokenParam{
		Key:  user.Key,
		Id:   tokenInfo.Id,
		Name: user.Name,
	}
	token, utilErr := uc.createLoginToken(&tokenParam)
	if utilErr != nil {
		return nil, utilErr
	}

	// delete before refresh token in mongo
	deleteParam := dto.DeleteTokenId{Id: param.Id}
	if utilErr := uc.DeleteRefreshToken(&deleteParam); utilErr != nil {
		return nil, utilErr
	}

	return token, nil
}

func (uc *Usecase) readLoginToken(
	param *dto.ReadLoginTokenParam) (*dto.LoginTokenInfo, *util.Error) {
	// read refresh token in mongo
	refreshToken, err := uc.mongo.ReadLoginToken(param)
	if err != nil {
		return nil, util.ErrorHandle(uc.checkMongoError(err), err)
	}

	// check expiration
	if time.Now().After(refreshToken.Expiration) {
		return nil, util.ErrorHandle(http.StatusUnauthorized, fmt.Errorf("expiration refresh token"))
	}

	// check token valid
	tokenInfo, err := uc.loginTokenValidCheck(refreshToken.Token)
	if err != nil {
		return nil, util.ErrorHandle(http.StatusUnauthorized, err)
	}

	// extraction token info
	operator := uc.extractionLoginToken(tokenInfo)

	return operator, nil
}

func (uc *Usecase) loginTokenValidCheck(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(uc.config.Auth().TokenSecret()), nil
	})
	return token, err
}

func (uc *Usecase) extractionLoginToken(
	token *jwt.Token) *dto.LoginTokenInfo {
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	userName := claims["user_name"].(string)

	return &dto.LoginTokenInfo{
		Id:   userId,
		Name: userName,
	}
}
