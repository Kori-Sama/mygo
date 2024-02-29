package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type allows []string

func (a *allows) join() string {
	return strings.Join(*a, ",")
}

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		origin := ctx.Request.Header.Get("Origin")

		allowOrigin := allows{"*"}
		allowMethod := allows{"GET", "POST"}
		allowHeader := allows{"Origin", "Content-Type", "Authorization", "Accept", "X-Requested-With"}
		exposeHeader := allows{"Authorization", "Content-Length",
			"Cache-Control", "Content-Language", "Content-Type",
			"Access-Control-Allow-Origin", "Access-Control-Allow-Headers"}

		if origin != "" {
			ctx.Header("Access-Control-Allow-Origin", allowOrigin.join())
			ctx.Header("Access-Control-Allow-Methods", allowMethod.join())
			ctx.Header("Access-Control-Allow-Headers", allowHeader.join())
			ctx.Header("Access-Control-Expose-Headers", exposeHeader.join())
			ctx.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	}
}
