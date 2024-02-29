package utils

import (
	"github.com/gin-gonic/gin"
)

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
