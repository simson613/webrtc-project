package mongo

import (
	"context"
	"github/simson613/webrtc-project/auth/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *mongoDB) ReadUserById(
	param *dto.ReadUserByIdParam) (*dto.ReadUserById, error) {
	coll := m.db.Collection("users")

	filter := bson.D{primitive.E{Key: "user_id", Value: param.Id}}
	result := dto.ReadUserById{}

	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	return &result, err
}
