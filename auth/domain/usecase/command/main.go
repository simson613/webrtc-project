package command

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

type Command struct {
	config config.ConfigInterface
	mongo  mongo.MongoDBInterface
}

func InitCommand(config config.ConfigInterface,
	mongo mongo.MongoDBInterface) *Command {
	return &Command{
		config: config,
		mongo:  mongo,
	}
}

func (c *Command) MongoDBTransactionHandler(param interface{}) (interface{}, *util.Error) {
	session, err := c.mongo.StartTransaction()
	if err != nil {
		return nil, util.DefaultErrorHandle(err)
	}
	defer session.EndSession(context.Background())

	callback := func(sessionContext mongoDB.SessionContext) (interface{}, error) {

		switch param := param.(type) {
		case *dto.SubscribeCreateUser:
			return c.mongo.CreateUser(param)
		case *dto.CreateLoginRefreshToken:
			return c.mongo.CreateLoginRefreshToken(param)
		case *dto.DeleteTokenId:
			return c.mongo.DeleteLoginRefreshToken(param)
		default:
			return nil, fmt.Errorf("not found param %s", param)
		}
	}

	options := c.mongo.TransactionOption()
	result, err := session.WithTransaction(context.Background(), callback, options)
	if err != nil {
		return nil, util.ErrorHandle(c.checkMongoError(err), err)
	}
	return result, nil
}

func (c *Command) checkMongoError(err error) int {
	if err == mongoDB.ErrNoDocuments {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
