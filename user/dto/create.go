package dto

import "time"

type CreateUserParam struct {
	Id       string `json:"user_id" binding:"required" example:"tester"`
	Name     string `json:"user_name" binding:"required" exmaple:"테스터"`
	Password string `json:"password" binding:"required" example:"123"`
}

type CreateUser struct {
	Key       string    `json:"user_key" bson:"_id"`
	Id        string    `json:"user_id" bson:"user_id"`
	Name      string    `json:"user_name" bson:"user_name"`
	CreatedAt time.Time `json:"created_at" bson:"registration_date"`
}

// type CreateUserPassword struct {
// 	Id       string `bson:"user_id"`
// 	Password string `bson:"password"`
// }
