package test

import (
	"mygo/config"
	"mygo/pkg/utils"
	"testing"
)

func TestJwt(t *testing.T) {
	config.JwtConfig.Secret = "mygo"
	config.JwtConfig.TokenExpire = 60
	token, _ := utils.GenerateToken(0, "name")
	t.Log(token)
	t.Log(utils.ParseToken(token))
	actual := utils.IsTokenValid(token)
	want := true
	if actual != want {
		t.Errorf("got %v, want %v", actual, want)
	}
}
