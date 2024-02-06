package controller

import (
	"mygo/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	if err := service.Login(username, password); err != nil {
		ctx.JSON(http.StatusBadRequest, Bad(err.Error()))
		return
	}
	ctx.JSON(200, Ok(nil))
}
