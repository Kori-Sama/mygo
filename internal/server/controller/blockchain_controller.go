package controller

import (
	"mygo/internal/pkg/common"
	"mygo/internal/pkg/utils"
	"mygo/internal/server/service"

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
	user := utils.GetLoginUser(ctx)
	if user == nil {
		ctx.JSON(common.UNAUTHORIZED, common.NoAuth(common.ErrorGetInfoFromToken.Error()))
		return
	}

	if err := service.CreateWallet(user.ID, passphrase); err != nil {
		if common.CheckInternalError(err) {
			ctx.JSON(common.INTERNAL_SERVER_ERROR, common.InternalError(err.Error()))
			return
		}
		ctx.JSON(common.BAD_REQUEST, common.Bad(err.Error()))
		return
	}

	ctx.JSON(common.OK, common.Ok(nil))
}

// @Summary		get balance
// @Description	get balance amount from wallet
// @Tags			blockchain
// @Produce		json
// @Success		200	{object}	common.Result	"OK"
// @Router			/api/blockchain/getBalance [get]
func GetBalance(ctx *gin.Context) {
	user := utils.GetLoginUser(ctx)
	if user == nil {
		ctx.JSON(common.UNAUTHORIZED, common.NoAuth(common.ErrorGetInfoFromToken.Error()))
		return
	}

	balance, err := service.GetBalance(user.ID)
	if err != nil {
		if common.CheckInternalError(err) {
			ctx.JSON(common.INTERNAL_SERVER_ERROR, common.InternalError(err.Error()))
			return
		}
		ctx.JSON(common.BAD_REQUEST, common.Bad(err.Error()))
		return
	}
	ctx.JSON(common.OK, common.Ok(balance))
}

// // @Summary		transfer funds
// // @Description	transfer funds from one wallet to another
// // @Tags			blockchain
// // @Accept			json
// // @Produce		json
// // @Param			transfer	body		common.TransferRequest	true	"transfer request json"
// // @Success		200			{object}	common.Result			"OK"
// // @Router			/api/blockchain/transfer [post]
// func Transfer(ctx *gin.Context) {
// 	user := utils.GetLoginUser(ctx)
// 	if user == nil {
// 		ctx.JSON(common.UNAUTHORIZED, common.NoAuth(common.ErrorGetInfoFromToken.Error()))
// 		return
// 	}

// 	var transferRequest common.TransferRequest
// 	if err := ctx.ShouldBindJSON(&transferRequest); err != nil {
// 		ctx.JSON(400, common.Bad(err.Error()))
// 		return
// 	}

// 	if err := service.Transfer(user.Name, transferRequest.Passphrase, transferRequest.ToName, transferRequest.Amount); err != nil {
// 		if common.CheckInternalError(err) {
// 			ctx.JSON(500, common.InternalError(err.Error()))
// 			return
// 		}
// 		ctx.JSON(400, common.Bad(err.Error()))
// 		return
// 	}

// 	ctx.JSON(200, common.Ok(nil))
// }
