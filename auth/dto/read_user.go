package dto

type ReadUserByIdParam struct {
	Id string `bson:"user_id"`
}

type ReadUserById struct {
	Name     string `bson:"user_name"`
	Password string `bson:"password"`
}
