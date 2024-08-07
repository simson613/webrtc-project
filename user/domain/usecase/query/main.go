package query

import (
	"github/simson613/webrtc-project/user/adapter/repository/mongo"
	"github/simson613/webrtc-project/user/config"
	"net/http"

	mongoDB "go.mongodb.org/mongo-driver/mongo"
)

type Query struct {
	config config.ConfigInterface
	mongo  mongo.MongoDBInterface
}

func InitQuery(config config.ConfigInterface,
	mongo mongo.MongoDBInterface) *Query {
	return &Query{
		config: config,
		mongo:  mongo,
	}
}

func (uc *Query) checkMongoError(err error) int {
	if err == mongoDB.ErrNoDocuments {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
