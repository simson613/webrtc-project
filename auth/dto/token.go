package dto

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateLoginTokenParam struct {
	Id   string
	Name string
	// Password string
}

type LoginTokenCliams struct {
	Id   string `json:"user_id"`
	Name string `json:"user_name"`
	jwt.StandardClaims
}

type LoginToken struct {
	AccessToken    *LoginAccessToken
	RefreshTokenId string
}

type LoginAccessToken struct {
	Token string `json:"token"`
}

type LoginRefreshToken struct {
	Token      string    `bson:"token"`
	Expiration time.Time `bson:"expiration"`
}

type DeleteTokenId struct {
	Id primitive.ObjectID `bson:"_id"`
}
