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
// @Success		200			{object}	Result			"OK"
// @Router			/api/blockchain/createWallet/{username}/{passphrase} [post]
func CreateWallet(ctx *gin.Context) {
	username := ctx.Param("username")
	passphrase := ctx.Param("passphrase")
	if err := service.CreateWallet(username, passphrase); err != nil {
		common.SelectInternalError(ctx, err)
		ctx.JSON(400, common.Bad(err.Error()))
		return
	}

	ctx.JSON(200, common.Ok(nil))
}

// @Summary		get balance
// @Description	get balance amount from wallet
// @Tags			blockchain
// @Param			username	path	string	true	"username"
// @Produce		json
// @Success		200			{object}	Result			"OK"
// @Router			/api/blockchain/getBalance/{username} [get]
func GetBalance(ctx *gin.Context) {
	username := ctx.Param("username")
	balance, err := service.GetBalance(username)
	if err != nil {
		common.SelectInternalError(ctx, err)
		ctx.JSON(400, common.Bad(err.Error()))
		return
	}
	ctx.JSON(200, common.Ok(balance))
}

// @Summary		transfer funds
// @Description	transfer funds from one wallet to another
// @Tags			blockchain
// @Accept			json
// @Produce		json
// @Param			transfer	body		TransferRequest	true	"transfer request json"
// @Success		200			{object}	Result			"OK"
// @Router			/api/blockchain/transfer [post]
func Transfer(ctx *gin.Context) {
	var transferRequest TransferRequest
	if err := ctx.ShouldBindJSON(&transferRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := service.Transfer(transferRequest.Username, transferRequest.Passphrase, transferRequest.ToName, transferRequest.Amount); err != nil {
		common.SelectInternalError(ctx, err)
		ctx.JSON(400, common.Bad(err.Error()))
		return
	}

	ctx.JSON(200, common.Ok(nil))
}

type TransferRequest struct {
	Username   string `json:"username"`
	Passphrase string `json:"passphrase"`
	ToName     string `json:"toName"`
	Amount     string `json:"amount"`
}
