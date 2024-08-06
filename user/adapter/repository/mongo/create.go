package mongo

import (
	"context"
	"github/simson613/webrtc-project/user/dto"
)

func (m *mongoDB) CreateUser(user *dto.CreateUserInView) error {
	coll := m.db.Collection("users")

	_, err := coll.InsertOne(context.TODO(), &user)
	return err
}
