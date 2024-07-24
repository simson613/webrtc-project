package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReadLoginTokenParam struct {
	Id primitive.ObjectID `bson:"_id"`
}

type ReadLoginToken struct {
	Token      string    `bson:"token"`
	Expiration time.Time `bson:"expiration"`
}

type LoginTokenInfo struct {
	Id   string `json:"user_id"`
	Name string `json:"user_name"`
}
