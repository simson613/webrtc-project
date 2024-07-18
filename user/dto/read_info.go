package dto

type ReadUserInfoParam struct {
	Key string `uri:"user_key" bson:"_id"`
}

type ReadUserInfo struct {
	Name string `json:"user_name" bson:"user_name" example:"테스터"`
}
