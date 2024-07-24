package mongo

import (
	"context"
	"github/simson613/webrtc-project/auth/dto"
)

func (m *mongoDB) CreateLoginRefreshToken(
	refreshToken *dto.CreateLoginRefreshToken) (interface{}, error) {
	coll := m.db.Collection("user_refresh")

	result, err := coll.InsertOne(context.TODO(), &refreshToken)
	return result.InsertedID, err
}
