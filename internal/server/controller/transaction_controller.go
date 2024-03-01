package controller

import (
	"mygo/internal/pkg/common"
	"mygo/internal/pkg/utils"
	"mygo/internal/server/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetTransaction(ctx *gin.Context) {
	transactionIDStr := ctx.Param("id")
	transactionID, err := strconv.Atoi(strings.TrimSpace(transactionIDStr))
	if err != nil {
		ctx.JSON(common.BAD_REQUEST, common.Bad(common.ErrorInvalidParam.Error()))
		return
	}

	user := utils.GetLoginUser(ctx)
	if user == nil {
		ctx.JSON(common.UNAUTHORIZED, common.NoAuth(common.ErrorGetInfoFromToken.Error()))
		return
	}
	transaction, err := service.GetTransaction(user.ID, transactionID)
	if err != nil {
		if common.CheckInternalError(err) {
			ctx.JSON(common.INTERNAL_SERVER_ERROR, common.InternalError(err.Error()))
			return
		}
		ctx.JSON(common.BAD_REQUEST, common.Bad(err.Error()))
		return
	}
	ctx.JSON(common.OK, common.Ok(transaction))
}

func GetTransactions(ctx *gin.Context) {
	user := utils.GetLoginUser(ctx)
	if user == nil {
		ctx.JSON(common.UNAUTHORIZED, common.NoAuth(common.ErrorGetInfoFromToken.Error()))
		return
	}
	transactions, err := service.GetTransactions(user.ID)
	if err != nil {
		if common.CheckInternalError(err) {
			ctx.JSON(common.INTERNAL_SERVER_ERROR, common.InternalError(err.Error()))
			return
		}
		ctx.JSON(common.BAD_REQUEST, common.Bad(err.Error()))
		return
	}
	ctx.JSON(common.OK, common.Ok(transactions))
}

func NewTransaction(ctx *gin.Context) {
	user := utils.GetLoginUser(ctx)
	if user == nil {
		ctx.JSON(common.UNAUTHORIZED, common.NoAuth(common.ErrorGetInfoFromToken.Error()))
		return
	}
	transaction := common.NewTransactionRequest{}
	ctx.ShouldBind(&transaction)
	transactionID, err := service.NewTransaction(user.ID, transaction)
	if err != nil {
		if common.CheckInternalError(err) {
			ctx.JSON(common.INTERNAL_SERVER_ERROR, common.InternalError(err.Error()))
			return
		}
		ctx.JSON(common.BAD_REQUEST, common.Bad(err.Error()))
		return
	}
	ctx.JSON(common.OK, common.Ok(gin.H{"id": transactionID}))
}

func SaveTransaction(ctx *gin.Context) {
	user := utils.GetLoginUser(ctx)
	if user == nil {
		ctx.JSON(common.UNAUTHORIZED, common.NoAuth(common.ErrorGetInfoFromToken.Error()))
		return
	}
	transaction := common.TransactionRequest{}
	ctx.ShouldBind(&transaction)
	err := service.SaveTransaction(user.ID, transaction)
	if err != nil {
		if common.CheckInternalError(err) {
			ctx.JSON(common.INTERNAL_SERVER_ERROR, common.InternalError(err.Error()))
			return
		}
		ctx.JSON(common.BAD_REQUEST, common.Bad(err.Error()))
		return
	}
	ctx.JSON(common.OK, common.Ok(nil))
}

func PublishTransaction(ctx *gin.Context) {
	user := utils.GetLoginUser(ctx)
	if user == nil {
		ctx.JSON(common.UNAUTHORIZED, common.NoAuth(common.ErrorGetInfoFromToken.Error()))
		return
	}
	transaction := common.TransactionRequest{}
	ctx.ShouldBind(&transaction)
	err := service.PublishTransaction(user.ID, transaction)
	if err != nil {
		if common.CheckInternalError(err) {
			ctx.JSON(common.INTERNAL_SERVER_ERROR, common.InternalError(err.Error()))
			return
		}
		ctx.JSON(common.BAD_REQUEST, common.Bad(err.Error()))
		return
	}
	ctx.JSON(common.OK, common.Ok(nil))
}
