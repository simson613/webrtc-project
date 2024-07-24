package mongo

import (
	"context"
	"github/simson613/webrtc-project/auth/dto"
)

func (m *mongoDB) CreateUser(user *dto.SubscribeCreateUser) error {
	coll := m.db.Collection("users")

	_, err := coll.InsertOne(context.TODO(), &user)
	return err
}
