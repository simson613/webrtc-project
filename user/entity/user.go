package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Key       string         `gorm:"primary_key; type:varchar(32);"`
	Id        string         `gorm:"uniqueIndex:idx_user_id; type:varchar(16); not null;"`
	Name      string         `gorm:"uniqueIndex:idx_user_name; type:varchar(10); not null;"`
	CreatedAt time.Time      `gorm:"default: current_timestamp; not null"`
	UpdatedAt time.Time      `gorm:"default: null"`
	DeletedAt gorm.DeletedAt `gorm:"default: null"`
}
