package middlewares

import (
	"mygo/config"
	"mygo/model"
	"mygo/pkg/common"
	"mygo/pkg/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	TokenName   = "Authorization"
	TokenPrefix = "Bearer: "
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader(TokenName)
		if token == "" || !strings.HasPrefix(token, TokenPrefix) {
			ctx.AbortWithStatusJSON(403, common.NoAuth())
			return
		}

		token = token[len(TokenPrefix):]
		claims, err := utils.ParseToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(403, common.NoAuth())
			return
		}

		user, err := model.GetUserById(claims.Id)
		if err != nil || user.Token != token {
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
			if err = user.UpdateToken(); err != nil {
				ctx.AbortWithStatusJSON(500, common.InternalError(err.Error()))
				return
			}
		}

		ctx.Next()
	}
}
