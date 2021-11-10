package tools

import (
	"errors"
	"ginapi/config"
	"time"

	"github.com/golang-jwt/jwt"
)

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
			Issuer:    config.C.App.Name,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(config.C.Jwt.Secret))
}

func ParseJwtToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(config.C.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
