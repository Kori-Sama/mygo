package controller

import (
	"mygo/pkg/common"
	"mygo/service"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	if err := service.Login(username, password); err != nil {
		common.SelectInternalError(ctx, err)
		ctx.JSON(400, common.Bad(err.Error()))
		return
	}
	ctx.JSON(200, common.Ok(nil))
}
