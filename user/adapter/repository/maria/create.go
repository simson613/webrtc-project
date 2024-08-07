package maria

import (
	"github/simson613/webrtc-project/user/domain/entity"
	"github/simson613/webrtc-project/user/dto"
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
