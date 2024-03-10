package middlewares

import (
	"mygo/config"
	"mygo/internal/pkg/common"
	"mygo/internal/pkg/constants"
	"mygo/internal/pkg/utils"

	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader(constants.TOKEN_NAME)
		if token == "" || !strings.HasPrefix(token, constants.TOKEN_PREFIX) {
			ctx.AbortWithStatusJSON(common.UNAUTHORIZED, common.NoAuth(common.ErrorInvalidToken.Error()))
			return
		}

		token = token[len(constants.TOKEN_PREFIX):]
		claims, err := utils.ParseToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(common.UNAUTHORIZED, common.NoAuth(common.ErrorInvalidToken.Error()))
			return
		}

		duration, err := utils.GetTokenDuration(token)
		if err != nil {
			ctx.AbortWithStatusJSON(common.UNAUTHORIZED, common.NoAuth(common.ErrorInvalidToken.Error()))
			return
		}
		if duration < 0 {
			ctx.AbortWithStatusJSON(common.UNAUTHORIZED, common.NoAuth(common.ErrorExpiredToken.Error()))
			return
		}

		if duration < time.Duration(config.Jwt.RefreshExpire)*time.Minute {
			newToken, err := utils.GenerateToken(claims.ID, claims.Name, claims.Role)
			if err != nil {
				ctx.AbortWithStatusJSON(common.UNAUTHORIZED, common.InternalError(err.Error()))
				return
			}

			ctx.Header(constants.TOKEN_NAME, constants.TOKEN_PREFIX+newToken)
		}

		ctx.Set(constants.LOGIN_USER, common.LoginUser{ID: claims.ID, Name: claims.Name, Role: claims.Role})

		ctx.Next()
	}
}
