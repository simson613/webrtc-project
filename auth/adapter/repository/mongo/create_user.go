package mongo

import (
	"context"
	"github/simson613/webrtc-project/auth/dto"
)

func (m *mongoDB) CreateUser(user *dto.SubscribeCreateUser) error {
	collection := m.db.Collection("users")

	_, err := collection.InsertOne(context.TODO(), &user)
	return err
}
