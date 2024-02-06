package controller

import (
	result "mygo/internal/pkg"
	"mygo/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateWallet(ctx *gin.Context) {
	username := ctx.Query("username")
	passphrase := ctx.Query("passphrase")
	if err := service.CreateWallet(username, passphrase); err != nil {
		ctx.JSON(http.StatusBadRequest, result.Error(400, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Ok(nil))
}

// func Transfer(ctx *gin.Context) {

// }
