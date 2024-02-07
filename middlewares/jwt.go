package middlewares

import (
	"mygo/config"
	"mygo/pkg/common"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	Id   int
	Name string
	jwt.RegisteredClaims
}

var secret = []byte(config.JwtConfig.Secret)

func GenerateToken(id int, name string) (string, error) {
	jwtClaims := JwtClaims{
		id,
		name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(config.JwtConfig.TokenExpire))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "MyGO",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	return token.SignedString(secret)
}

func ParseToken(tokenStr string) (JwtClaims, error) {
	jwtClaims := JwtClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &jwtClaims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err == nil && !token.Valid {
		err = common.ErrorInvalidToken
	}
	return jwtClaims, err
}

func IsTokenValid(tokenStr string) bool {
	_, err := ParseToken(tokenStr)
	return err == nil
}
