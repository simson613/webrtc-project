package mongo

import (
	"context"
	"github/simson613/webrtc-project/user/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *mongoDB) ReadUserInfo(
	condition *dto.ReadUserInfoParam) (*dto.ReadUserInfo, error) {
	collection := m.db.Collection("users")

	filter := bson.D{primitive.E{Key: "_id", Value: condition.Key}}
	result := dto.ReadUserInfo{}

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	return &result, err
}
