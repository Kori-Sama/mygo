package middlewares

import (
	"mygo/config"
	"mygo/pkg/common"
	"mygo/pkg/constants"
	"mygo/pkg/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader(constants.TOKEN_NAME)
		if token == "" || !strings.HasPrefix(token, constants.TOKEN_PREFIX) {
			ctx.AbortWithStatusJSON(403, common.NoAuth())
			return
		}

		token = token[len(constants.TOKEN_PREFIX):]
		claims, err := utils.ParseToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(403, common.NoAuth())
			return
		}

		duration, err := utils.GetTokenDuration(token)
		if err != nil {
			ctx.AbortWithStatusJSON(403, common.NoAuth())
			return
		}
		if duration < 0 {
			ctx.AbortWithStatusJSON(403, common.NoAuth())
			return
		}

		if duration < time.Duration(config.JwtConfig.RefreshExpire)*time.Minute {
			newToken, err := utils.GenerateToken(claims.Id, claims.Name)
			if err != nil {
				ctx.AbortWithStatusJSON(403, common.InternalError(err.Error()))
				return
			}

			ctx.Header(constants.TOKEN_NAME, constants.TOKEN_PREFIX+newToken)
		}

		ctx.Set(constants.LOGIN_USER, common.LoginUser{Id: claims.Id, Name: claims.Name})

		ctx.Next()
	}
}
