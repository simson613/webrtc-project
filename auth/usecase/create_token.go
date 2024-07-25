package usecase

import (
	"github/simson613/webrtc-project/auth/dto"
	"github/simson613/webrtc-project/auth/util"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *Usecase) createLoginToken(
	user *dto.CreateLoginTokenParam) (*dto.LoginToken, *util.Error) {
	tokenCliam := dto.LoginTokenCliams{
		Key:  user.Key,
		Id:   user.Id,
		Name: user.Name,
	}
	// create access token
	accessToken, err := uc.createAccessToken(&tokenCliam)
	if err != nil {
		return nil, util.DefaultErrorHandle(err)
	}

	// create refresh token
	refreshToken, err := uc.createRefreshToken(&tokenCliam)
	if err != nil {
		return nil, util.DefaultErrorHandle(err)
	}

	// insert refresh token in read db
	result, utilErr := uc.MongoDBTransactionHandler(refreshToken)
	if utilErr != nil {
		return nil, utilErr
	}
	refreshId := result.(primitive.ObjectID)

	return &dto.LoginToken{
		AccessToken:    accessToken,
		RefreshTokenId: refreshId.Hex(),
	}, nil
}

func (uc *Usecase) createAccessToken(
	payload *dto.LoginTokenCliams) (*dto.CreateLoginAccessToken, error) {
	now := time.Now()
	expiration := uc.config.Auth().AccessExpiration()
	expirationTime := now.Add(time.Minute * time.Duration(expiration))
	payload.StandardClaims = jwt.StandardClaims{ExpiresAt: jwt.At(expirationTime)}

	token, err := uc.createToken(payload)
	return &dto.CreateLoginAccessToken{Token: token}, err
}

func (uc *Usecase) createRefreshToken(
	payload *dto.LoginTokenCliams) (*dto.CreateLoginRefreshToken, error) {
	now := time.Now()
	expiration := uc.config.Auth().RefreshExpiration()
	expirationTime := now.Add(time.Hour * time.Duration(expiration))
	payload.StandardClaims = jwt.StandardClaims{ExpiresAt: jwt.At(expirationTime)}

	token, err := uc.createToken(payload)
	return &dto.CreateLoginRefreshToken{
		Token:      token,
		Expiration: expirationTime,
	}, err
}

func (uc *Usecase) createToken(payload *dto.LoginTokenCliams) (string, error) {
	jwtKey := uc.config.Auth().TokenSecret()
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, *payload)
	token, err := tokenClaim.SignedString([]byte(jwtKey))
	return token, err
}
