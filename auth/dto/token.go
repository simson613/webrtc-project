package dto

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type CreateLoginTokenParam struct {
	Key  string
	Id   string
	Name string
}

type LoginTokenCliams struct {
	Key  string `json:"user_key"`
	Id   string `json:"user_id"`
	Name string `json:"user_name"`
	jwt.StandardClaims
}

type LoginToken struct {
	AccessToken    *CreateLoginAccessToken
	RefreshTokenId string
}

type CreateLoginAccessToken struct {
	Token string `json:"token"`
}

type CreateLoginRefreshToken struct {
	Token      string    `bson:"token"`
	Expiration time.Time `bson:"expiration"`
}
