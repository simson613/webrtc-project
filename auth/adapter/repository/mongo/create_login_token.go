package mongo

import (
	"context"
	"github/simson613/webrtc-project/auth/dto"
)

func (m *mongoDB) CreateLoginRefreshToken(
	refreshToken *dto.LoginRefreshToken) (interface{}, error) {
	collection := m.db.Collection("user_refresh")

	result, err := collection.InsertOne(context.TODO(), &refreshToken)
	return result.InsertedID, err
}
