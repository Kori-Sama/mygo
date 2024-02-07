package controller

import (
	"mygo/pkg/common"
	"mygo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		create wallet
// @Description	create wallet in blockchain
// @Tags			blockchain
// @Param			username	path	string	true	"username"
// @Param			passphrase	path	string	true	"passphrase"
// @Produce		json
// @Success		200	{object}	nil		"OK"
// @failure		400	{object}	string	"Bad Request"
// @failure		500	{object}	string	"Internal Server Error"
// @Router			/api/blockchain/createWallet/{username}/{passphrase} [post]
func CreateWallet(ctx *gin.Context) {
	username := ctx.Param("username")
	passphrase := ctx.Param("passphrase")
	if err := service.CreateWallet(username, passphrase); err != nil {
		if common.CheckInternalError(err) {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// @Summary		get balance
// @Description get balance amount from wallet
// @Tags			blockchain
// @Param			username	path	string	true	"username"
// @Produce		json
// @Success		200	{object}	float64 "OK"
// @failure		400	{object}	string	"Bad Request"
// @failure		500	{object}	string	"Internal Server Error"
// @Router			/api/blockchain/getBalance/{username} [get]
func GetBalance(ctx *gin.Context) {
	username := ctx.Param("username")
	balance, err := service.GetBalance(username)
	if err != nil {
		if common.CheckInternalError(err) {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, balance)
}

// func Transfer(ctx *gin.Context) {

// }
