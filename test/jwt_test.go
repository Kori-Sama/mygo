package test

import (
	"mygo/config"
	"mygo/middlewares"
	"testing"
)

func TestJwt(t *testing.T) {
	config.JwtConfig.Secret = "mygo"
	config.JwtConfig.TokenExpire = 60
	token, _ := middlewares.GenerateToken(0, "name")
	t.Log(token)
	t.Log(middlewares.ParseToken(token))
	actual := middlewares.IsTokenValid(token)
	want := true
	if actual != want {
		t.Errorf("got %v, want %v", actual, want)
	}
}
