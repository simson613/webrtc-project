package command

import (
	"github/simson613/webrtc-project/user/dto"
	"github/simson613/webrtc-project/user/util"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (c *Command) CreateUser(param *dto.CreateUserParam) *util.Error {
	userParam := dto.CreateUser{
		Key:       uuid.NewString(),
		Id:        param.Id,
		Name:      param.Name,
		CreatedAt: time.Now(),
	}
	if utilErr := c.MariaDBTransactionHandler(&userParam); utilErr != nil {
		return utilErr
	}

	viewParam := dto.CreateUserInView{
		Key:  userParam.Key,
		Id:   userParam.Id,
		Name: userParam.Name,
	}
	if utilErr := c.MongoDBTransactionHandler(&viewParam); utilErr != nil {
		//delete user in maria
		return utilErr
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	if err != nil {
		return util.DefaultErrorHandle(err)
	}
	publishParam := dto.PublishCreateUser{
		Key:      userParam.Key,
		Id:       param.Id,
		Name:     param.Name,
		Password: string(hashedPassword),
	}
	if err := c.producer.CreateUser(&publishParam); err != nil {
		//delete user in maria and mongo
		return util.DefaultErrorHandle(err)
	}

	return nil
}
