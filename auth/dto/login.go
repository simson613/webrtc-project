package dto

type LoginParam struct {
	Id       string `json:"user_id"`
	Password string `json:"password"`
}
