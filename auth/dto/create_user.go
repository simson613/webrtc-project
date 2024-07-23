package dto

type SubscribeCreateUser struct {
	Key      string `json:"user_key" bson:"_id"`
	Id       string `json:"user_id" bson:"user_id"`
	Name     string `json:"user_name" bson:"user_name"`
	Password string `json:"password" bson:"password"`
}
