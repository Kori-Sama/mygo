package utils

import (
	"mygo/internal/pkg/common"
	"mygo/internal/pkg/constants"

	"github.com/gin-gonic/gin"
)

func GetLoginUser(ctx *gin.Context) *common.LoginUser {
	user := GetValueFromContext[common.LoginUser](ctx, constants.LOGIN_USER)
	return user
}

func GetValueFromContext[T any](ctx *gin.Context, key string) *T {
	value, ok := ctx.Get(key)
	if !ok {
		return nil
	}
	user, ok := value.(T)
	if !ok {
		return nil
	}
	return &user
}
