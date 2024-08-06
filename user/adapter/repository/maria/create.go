package maria

import (
	"github/simson613/webrtc-project/user/dto"
	"github/simson613/webrtc-project/user/entity"
)

func (m *mariaDB) CreateUser(user *dto.CreateUser) error {
	newUser := entity.User{
		Key:       user.Key,
		Id:        user.Id,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}

	return m.db.Create(&newUser).Error
}
