package usecase

import (
	"context"
	"fmt"
	"github/simson613/webrtc-project/auth/adapter/repository/mongo"
	"github/simson613/webrtc-project/auth/config"
	"github/simson613/webrtc-project/auth/dto"
	"github/simson613/webrtc-project/auth/util"
	"net/http"

	mongoDB "go.mongodb.org/mongo-driver/mongo"
)

type Usecase struct {
	config config.ConfigInterface
	mongo  mongo.MongoDBInterface
}

func InitUsecase(config config.ConfigInterface,
	mongo mongo.MongoDBInterface) *Usecase {
	return &Usecase{
		config: config,
		mongo:  mongo,
	}
}

func (uc *Usecase) MongoDBTransactionHandler(param interface{}) (interface{}, *util.Error) {
	session, err := uc.mongo.StartTransaction()
	if err != nil {
		return nil, util.DefaultErrorHandle(err)
	}
	defer session.EndSession(context.Background())

	callback := func(sessionContext mongoDB.SessionContext) (interface{}, error) {

		switch param := param.(type) {
		case *dto.SubscribeCreateUser:
			return nil, uc.mongo.CreateUser(param)
		case *dto.LoginRefreshToken:
			return uc.mongo.CreateLoginRefreshToken(param)
		default:
			return nil, fmt.Errorf("not found param %s", param)
		}
	}

	options := uc.mongo.TransactionOption()
	result, err := session.WithTransaction(context.Background(), callback, options)
	if err != nil {
		return nil, util.ErrorHandle(uc.checkMongoError(err), err)
	}
	return result, nil
}

func (uc *Usecase) checkMongoError(err error) int {
	if err == mongoDB.ErrNoDocuments {
		return http.StatusNoContent
	}
	return http.StatusInternalServerError
}
