package tools

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var secret = []byte("test")

const TokenExpireDuration = time.Hour * 2

type MyClaims struct {
	UID int64 `json:"uid"`
	jwt.StandardClaims
}

func MakeJwtToken(id int64) (string, error) {
	c := MyClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "gin-api",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(secret)
}

func ParseJwtToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
