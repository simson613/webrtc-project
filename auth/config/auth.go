package config

import (
	"os"
	"strconv"
)

type AuthInterface interface {
	TokenSecret() string
	AccessExpiration() int
	RefreshExpiration() int
}

type Auth struct {
	tokenSecret       string
	accessExpiration  int
	refreshExpiration int
}

func initAuthConfig() *Auth {
	tokenSecret := os.Getenv("TOKEN_SECRET")
	if tokenSecret == "" {
		tokenSecret = "habreljn53j"
	}
	accessExpiration, _ := strconv.Atoi(os.Getenv("ACCESS_TIME"))
	if accessExpiration == 0 {
		accessExpiration = 30
	}
	refreshExpiration, _ := strconv.Atoi(os.Getenv("REFRESH_TIME"))
	if refreshExpiration == 0 {
		refreshExpiration = 24
	}

	return &Auth{
		tokenSecret:       tokenSecret,
		accessExpiration:  accessExpiration,
		refreshExpiration: refreshExpiration,
	}
}

func (auth *Auth) TokenSecret() string {
	return auth.tokenSecret
}

func (auth *Auth) AccessExpiration() int {
	return auth.accessExpiration
}

func (auth *Auth) RefreshExpiration() int {
	return auth.refreshExpiration
}
