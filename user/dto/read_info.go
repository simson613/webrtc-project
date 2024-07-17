package dto

type ReadUserInfoParam struct {
	Key string `uri:"user_key" bson:"_id"`
}

type ReadUserInfo struct {
	Name      string `bson:"user_name"`
	CreatedAt string `bson:"registration_date"`
}
