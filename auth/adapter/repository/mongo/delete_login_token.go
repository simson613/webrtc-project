package mongo

import (
	"context"
	"github/simson613/webrtc-project/auth/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *mongoDB) DeleteLoginRefreshToken(
	refreshToken *dto.DeleteTokenId) (interface{}, error) {
	coll := m.db.Collection("user_refresh")

	filter := bson.D{primitive.E{Key: "_id", Value: refreshToken.Id}}
	return coll.DeleteOne(context.TODO(), filter)
}
