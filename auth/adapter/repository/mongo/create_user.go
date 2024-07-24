package mongo

import (
	"context"
	"github/simson613/webrtc-project/auth/dto"
)

func (m *mongoDB) CreateUser(
	user *dto.SubscribeCreateUser) (interface{}, error) {
	coll := m.db.Collection("users")

	return coll.InsertOne(context.TODO(), &user)
}
