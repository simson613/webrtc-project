package mongo

import (
	"context"
	"github/simson613/webrtc-project/user/dto"
)

func (m *mongoDB) CreateUser(user *dto.CreateUser) error {
	collection := m.db.Collection("users")

	_, err := collection.InsertOne(context.TODO(), &user)
	return err
}
