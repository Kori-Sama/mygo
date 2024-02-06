package controller

import (
	"mygo/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateWallet(ctx *gin.Context) {
	username := ctx.Query("username")
	passphrase := ctx.Query("passphrase")
	if err := service.CreateWallet(username, passphrase); err != nil {
		ctx.JSON(http.StatusBadRequest, Bad(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, Ok(nil))
}

// func Transfer(ctx *gin.Context) {

// }
