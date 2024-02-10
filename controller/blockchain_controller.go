package controller

import (
	"mygo/pkg/common"
	"mygo/pkg/constants"
	"mygo/pkg/utils"
	"mygo/service"

	"github.com/gin-gonic/gin"
)

// @Summary		create wallet
// @Description	create wallet in blockchain
// @Tags			blockchain
// @Param			passphrase	path	string	true	"passphrase"
// @Produce		json
// @Success		200	{object}	common.Result	"OK"
// @Router			/api/blockchain/createWallet/{passphrase} [post]
func CreateWallet(ctx *gin.Context) {
	passphrase := ctx.Param("passphrase")
	user := utils.GetValueFromContext[common.LoginUser](ctx, constants.LOGIN_USER)
	if user == nil {
		ctx.JSON(403, common.NoAuth(common.ErrorGetInfoFromToken.Error()))
		return
	}

	if err := service.CreateWallet(user.Id, passphrase); err != nil {
		if common.CheckInternalError(err) {
			ctx.JSON(500, common.InternalError(err.Error()))
			return
		}
		ctx.JSON(400, common.Bad(err.Error()))
		return
	}

	ctx.JSON(200, common.Ok(nil))
}

// @Summary		get balance
// @Description	get balance amount from wallet
// @Tags			blockchain
// @Produce		json
// @Success		200	{object}	common.Result	"OK"
// @Router			/api/blockchain/getBalance [get]
func GetBalance(ctx *gin.Context) {
	user := utils.GetValueFromContext[common.LoginUser](ctx, constants.LOGIN_USER)
	if user == nil {
		ctx.JSON(403, common.NoAuth(common.ErrorGetInfoFromToken.Error()))
		return
	}

	balance, err := service.GetBalance(user.Id)
	if err != nil {
		if common.CheckInternalError(err) {
			ctx.JSON(500, common.InternalError(err.Error()))
			return
		}
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
// @Param			transfer	body		common.TransferRequest	true	"transfer request json"
// @Success		200			{object}	common.Result			"OK"
// @Router			/api/blockchain/transfer [post]
func Transfer(ctx *gin.Context) {
	user := utils.GetValueFromContext[common.LoginUser](ctx, constants.LOGIN_USER)
	if user == nil {
		ctx.JSON(403, common.NoAuth(common.ErrorGetInfoFromToken.Error()))
		return
	}

	var transferRequest common.TransferRequest
	if err := ctx.ShouldBindJSON(&transferRequest); err != nil {
		ctx.JSON(400, common.Bad(err.Error()))
		return
	}

	if err := service.Transfer(user.Name, transferRequest.Passphrase, transferRequest.ToName, transferRequest.Amount); err != nil {
		if common.CheckInternalError(err) {
			ctx.JSON(500, common.InternalError(err.Error()))
			return
		}
		ctx.JSON(400, common.Bad(err.Error()))
		return
	}

	ctx.JSON(200, common.Ok(nil))
}
