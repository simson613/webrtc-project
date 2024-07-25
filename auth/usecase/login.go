package usecase

import (
	"github/simson613/webrtc-project/auth/dto"
	"github/simson613/webrtc-project/auth/util"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (uc *Usecase) Login(
	param *dto.LoginParam) (*dto.LoginToken, *util.Error) {
	// 1. read user by id
	userParam := dto.ReadUserByIdParam{
		Id: param.Id,
	}
	user, err := uc.mongo.ReadUserById(&userParam)
	if err != nil {
		return nil, util.ErrorHandle(uc.checkMongoError(err), err)
	}

	// 2. compare password
	if err := uc.compareHashedPassword(user.Password, param.Password); err != nil {
		return nil, util.ErrorHandle(http.StatusNotFound, err)
	}

	// 3. create token
	createTokenParam := dto.CreateLoginTokenParam{
		Key:  user.Key,
		Id:   param.Id,
		Name: user.Name,
	}

	return uc.createLoginToken(&createTokenParam)
}

func (uc *Usecase) compareHashedPassword(pwd1 string, pwd2 string) error {
	return bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2))
}
