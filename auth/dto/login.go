package dto

type LoginParam struct {
	Id       string `json:"user_id" binding:"required" example:"tester123"`
	Password string `json:"password" binding:"required" example:"123"`
}
