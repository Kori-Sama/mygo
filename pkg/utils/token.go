package utils

import (
	"mygo/config"
	"mygo/pkg/common"
	"mygo/pkg/constants"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	ID   int
	Name string
	Role common.Role
	jwt.RegisteredClaims
}

var secret = []byte(config.Jwt.Secret)

func GenerateToken(id int, name string, role common.Role) (string, error) {
	jwtClaims := JwtClaims{
		id,
		name,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(config.Jwt.TokenExpire))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   constants.TOKEN_SUBJECT,
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

func GetTokenDuration(tokenStr string) (time.Duration, error) {
	token, err := ParseToken(tokenStr)
	if err != nil {
		return 0, common.ErrorInvalidToken
	}
	return time.Until(token.ExpiresAt.Time), nil
}

func IsTokenValid(tokenStr string) bool {
	_, err := ParseToken(tokenStr)
	return err == nil
}
