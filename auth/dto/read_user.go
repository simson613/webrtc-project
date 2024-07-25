package dto

type ReadUserByIdParam struct {
	Id string `bson:"user_id"`
}

type ReadUserById struct {
	Key      string `bson:"_id"`
	Name     string `bson:"user_name"`
	Password string `bson:"password"`
}
