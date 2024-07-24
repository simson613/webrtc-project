package dto

type LoginParam struct {
	Id       string `json:"user_id"`
	Password string `json:"password"`
}

type ReadUserByIdParam struct {
	Id string `bson:"user_id"`
}

type ReadUserById struct {
	Name     string `bson:"user_name"`
	Password string `bson:"password"`
}
