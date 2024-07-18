package usecase

import (
	"github/simson613/webrtc-project/user/dto"
	"github/simson613/webrtc-project/user/util"
	"time"

	"github.com/google/uuid"
)

func (uc *Usecase) CreateUser(param *dto.CreateUserParam) *util.Error {
	userParam := dto.CreateUser{
		Key:       uuid.NewString(),
		Id:        param.Id,
		Name:      param.Name,
		CreatedAt: time.Now(),
	}
	if utilErr := uc.MariaDBTransactionHandler(&userParam); utilErr != nil {
		return utilErr
	}

	viewParam := dto.CreateUserInView{
		Key:  userParam.Key,
		Id:   userParam.Id,
		Name: userParam.Name,
	}
	if utilErr := uc.MongoDBTransactionHandler(&viewParam); utilErr != nil {
		//delete user
		return utilErr
	}
	return nil
}
