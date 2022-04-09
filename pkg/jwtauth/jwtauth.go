package jwtauth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Options struct {
	Secret    string `yaml:"secret"`
	ExpiresIn int64  `yaml:"expires-in"`
}

type AuthUser struct {
	Id int64
	Username string
}

type JWTAuth struct {
	Options *Options
}

type Claims struct {
	Username string
	jwt.RegisteredClaims
}

func NewJWTAuth(o *Options) *JWTAuth {
	return &JWTAuth{
		Options: o,
	}
}

var (
	ErrTokenExpired = errors.New("token is expired")
	ErrInvalidToken = errors.New("couldn't handle this token")
)

func (j *JWTAuth) CreateToken(username string) (string, error) {
	tokenExpireDuration := time.Second * time.Duration(j.Options.ExpiresIn)
	// set our claims
	claims := Claims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpireDuration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.Options.Secret))
}

func (j *JWTAuth) ParseToken(tokenString string) (*Claims, error) {

	secret := []byte(j.Options.Secret)
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			// Token is expired
			if ve.Errors & jwt.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			}
		}
		return nil, ErrInvalidToken
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrInvalidToken
}

func (j *JWTAuth) RefreshToken(tokenString string) (string, error) {
	claims, err := j.ParseToken(tokenString)
	if err == nil {
		return j.CreateToken(claims.Username)
	}
	return "", err
}
