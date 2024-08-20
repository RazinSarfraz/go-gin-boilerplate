package service

import (
	"time"
	"user-backend/conf"
	"user-backend/models"

	"github.com/golang-jwt/jwt"
)

type JwtService interface {
	CreateLoginToken(user models.User) (string, error)
	VerifyLoginToken(tokenStr string) (*models.DecodeJWTClaims, error)
	CreateAuthToken(user models.User) (string, error)
	VerifyAuthToken(tokenStr string) (*models.DecodeJWTClaims, error)
}

type jwtService struct {
	cacheService CacheService
	config       conf.Config
}

func NewJWTService(cacheService CacheService, config *conf.Config) JwtService {
	return &jwtService{
		cacheService: cacheService,
		config:       *config,
	}
}

func (j *jwtService) CreateLoginToken(user models.User) (string, error) {
	claim := jwt.MapClaims{
		"exp":         time.Now().Add(time.Minute * 60).Unix(),
		"uid":         user.UserID,
		"phoneNumber": user.Phone,
		"tokenType":   models.LoginToken,
		"loginType":   models.UserLogin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	secret := j.config.JwtSecret

	return token.SignedString([]byte(secret))
}

func (j *jwtService) CreateAuthToken(user models.User) (string, error) {
	claim := jwt.MapClaims{
		"exp":         time.Now().Add(time.Minute * 60).Unix(),
		"uid":         user.UserID,
		"phoneNumber": user.Phone,
		"tokenType":   models.AuthToken,
		"loginType":   models.UserLogin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	secret := j.config.JwtSecret

	return token.SignedString([]byte(secret))
}

/*
* This method does following
* get jwt secret from conf, parse incoming token, get data stored in token, check if token has expired, check if uid exist,
* check active account id value, check account type
* @params token string
* @returns decoded jwt token, erro
 */
func (j jwtService) VerifyLoginToken(tokenStr string) (*models.DecodeJWTClaims, error) {
	secret := j.config.JwtSecret

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, models.GetStandardError("", models.INVALID_TOKEN, models.INVALID_TOKEN_MESSAGE)
	}
	if !token.Valid {
		return nil, models.GetStandardError("", models.INVALID_TOKEN, models.INVALID_TOKEN_MESSAGE)
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, models.GetStandardError("", models.INVALID_TOKEN, "cannot convert claim to MapClaims")
	}

	err = claim.Valid()
	if err != nil {
		return nil, models.GetStandardError("", models.TOKEN_EXPIRED, models.TOKEN_EXPIRED_MESSAGE)
	}

	expiredAtVal, found := claim["exp"]
	if !found {
		return nil, models.GetStandardError("", models.INVALID_TOKEN, models.INVALID_TOKEN_MESSAGE)
	}

	uidVal, found := claim["uid"]
	if !found {
		return nil, models.GetStandardError("", models.INVALID_TOKEN, models.INVALID_TOKEN_MESSAGE)
	}
	phoneNumber, found := claim["phoneNumber"]
	if !found {
		return nil, models.GetStandardError("", models.INVALID_TOKEN, models.INVALID_TOKEN_MESSAGE)
	}

	tokenType, found := claim["tokenType"]
	if !found || tokenType != models.LoginToken {
		return nil, models.GetStandardError("", models.INVALID_TOKEN, models.INVALID_TOKEN_MESSAGE)
	}

	loginType, found := claim["loginType"]
	if !found || loginType != models.UserLogin {
		return nil, models.GetStandardError("", models.INVALID_TOKEN, models.INVALID_TOKEN_MESSAGE)
	}

	return &models.DecodeJWTClaims{
		ExpiredAt: expiredAtVal.(float64),
		UserId:    uidVal.(string),
		Phone:     phoneNumber.(string),
		TokenType: tokenType.(string),
		LoginType: loginType.(string),
	}, nil
}

func (j jwtService) VerifyAuthToken(tokenStr string) (*models.DecodeJWTClaims, error) {
	secret := j.config.JwtSecret

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, models.GetStandardError("", models.INVALID_TOKEN, models.INVALID_TOKEN_MESSAGE)
	}
	if !token.Valid {
		return nil, models.GetStandardError("", models.INVALID_TOKEN, models.INVALID_TOKEN_MESSAGE)
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, models.GetStandardError("", models.INVALID_TOKEN, "cannot convert claim to MapClaims")
	}

	err = claim.Valid()
	if err != nil {
		return nil, models.GetStandardError("", models.TOKEN_EXPIRED, models.TOKEN_EXPIRED_MESSAGE)
	}

	expiredAtVal, found := claim["exp"]
	if !found {
		return nil, models.GetStandardError("", models.INVALID_TOKEN, models.INVALID_TOKEN_MESSAGE)
	}

	uidVal, found := claim["uid"]
	if !found {
		return nil, models.GetStandardError("", models.INVALID_TOKEN, models.INVALID_TOKEN_MESSAGE)
	}
	phoneNumber, found := claim["phoneNumber"]
	if !found {
		return nil, models.GetStandardError("", models.INVALID_TOKEN, models.INVALID_TOKEN_MESSAGE)
	}

	tokenType, found := claim["tokenType"]
	if !found || tokenType != models.AuthToken {
		return nil, models.GetStandardError("", models.INVALID_TOKEN, models.INVALID_TOKEN_MESSAGE)
	}

	loginType, found := claim["loginType"]
	if !found || loginType != models.UserLogin {
		return nil, models.GetStandardError("", models.INVALID_TOKEN, models.INVALID_TOKEN_MESSAGE)
	}

	return &models.DecodeJWTClaims{
		ExpiredAt: expiredAtVal.(float64),
		UserId:    uidVal.(string),
		Phone:     phoneNumber.(string),
		TokenType: tokenType.(string),
		LoginType: loginType.(string),
	}, nil
}
