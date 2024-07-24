package mongo

import (
	"context"
	"github/simson613/webrtc-project/auth/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *mongoDB) ReadLoginToken(
	param *dto.ReadLoginTokenParam) (*dto.ReadLoginToken, error) {
	coll := m.db.Collection("user_refresh")

	filter := bson.D{primitive.E{Key: "_id", Value: param.Id}}
	result := dto.ReadLoginToken{}

	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	return &result, err
}
