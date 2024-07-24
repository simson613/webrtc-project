package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type DeleteTokenId struct {
	Id primitive.ObjectID `bson:"_id"`
}
