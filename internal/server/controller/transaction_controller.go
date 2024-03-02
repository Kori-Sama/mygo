package controller

import (
	"mygo/internal/pkg/common"
	"mygo/internal/pkg/utils"
	"mygo/internal/server/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// @Summary		search transaction
// @Description	search transaction by title or description
// @Tags			transaction
// @Accept			json
// @Produce		json
// @Param			search	query		string							true	"search string"
// @Success		200		{object}	[]common.TransactionResponse	"OK"
// @Router			/api/transaction/search [get]
func SearchTransaction(ctx *gin.Context) {
	search := ctx.Query("search")
	transactions, err := service.SearchTransactions(search)
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

// @Summary		get transaction
// @Description	get single transaction by id, only allowed for admins
// @Tags			transaction
// @Accept			json
// @Produce		json
// @Param			id	path		int							true	"Transaction ID"
// @Success		200	{object}	common.TransactionResponse	"OK"
// @Router			/api/transaction/{id} [get]
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

	if user.Role != common.RoleAdmin {
		ctx.JSON(common.FORBIDDEN, common.Forbidden(common.ErrorNoPermission.Error()))
		return
	}

	transaction, err := service.GetTransaction(transactionID)
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

// @Summary		get self transactions
// @Description	get transactions of login user
// @Tags			transaction
// @Accept			json
// @Produce		json
// @Success		200	{object}	[]common.TransactionResponse	"OK"
// @Router			/api/transaction/self [get]
func GetOwnTransactions(ctx *gin.Context) {
	user := utils.GetLoginUser(ctx)
	if user == nil {
		ctx.JSON(common.UNAUTHORIZED, common.NoAuth(common.ErrorGetInfoFromToken.Error()))
		return
	}

	transactions, err := service.GetTransactionsByUserID(user.ID)
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

// @Summary		get all transactions
// @Description	get all transactions, admins can see all transactions, users can only see passed transactions
// @Tags			transaction
// @Accept			json
// @Produce		json
// @Success		200	{object}	[]common.TransactionResponse	"OK"
// @Router			/api/transaction [get]
func GetAllTransactions(ctx *gin.Context) {
	user := utils.GetLoginUser(ctx)
	if user == nil {
		ctx.JSON(common.UNAUTHORIZED, common.NoAuth(common.ErrorGetInfoFromToken.Error()))
		return
	}
	transactions, err := service.GetTransactions(user.Role)
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

// @Summary		get transactions by status
// @Description	get transactions by status, only admins have permission to access other statuses except passed
// @Tags			transaction
// @Accept			json
// @Produce		json
// @Param			status	query		string							true	"transaction status"
// @Success		200		{object}	[]common.TransactionResponse	"OK"
// @Router			/api/transaction/by [get]
func GetTransactionByStatus(ctx *gin.Context) {
	statusRaw := ctx.Query("status")
	status, err := utils.FilterStatus(statusRaw)
	if err != nil {
		ctx.JSON(common.BAD_REQUEST, common.Bad(err.Error()))
		return
	}

	user := utils.GetLoginUser(ctx)
	if user == nil {
		ctx.JSON(common.UNAUTHORIZED, common.NoAuth(common.ErrorGetInfoFromToken.Error()))
		return
	}

	if user.Role != common.RoleAdmin && status != common.StatusPassed {
		ctx.JSON(common.UNAUTHORIZED, common.Forbidden(common.ErrorNoPermission.Error()))
		return
	}

	transactions, err := service.GetTransactionByStatus(status)
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

// @Summary		Create a new transaction
// @Description	Create a new transaction
// @Tags			transaction
// @Accept			json
// @Produce		json
// @Success		200	{object}	common.Result	"OK"
// @Router			/api/transaction/new [post]
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

// @Summary		Save a transaction
// @Description	Save a transaction
// @Tags			transaction
// @Accept			json
// @Produce		json
// @Success		200	{object}	common.Result	"OK"
// @Router			/api/transaction/save [post]
func SaveTransaction(ctx *gin.Context) {
	user := utils.GetLoginUser(ctx)
	if user == nil {
		ctx.JSON(common.UNAUTHORIZED, common.NoAuth(common.ErrorGetInfoFromToken.Error()))
		return
	}
	transaction := common.TransactionRequest{}
	ctx.ShouldBind(&transaction)
	log.Debugf("SaveTransaction: %+v", transaction)
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

// @Summary		Publish a transaction
// @Description	Publish a transaction
// @Tags			transaction
// @Accept			json
// @Produce		json
// @Success		200	{object}	common.Result	"OK"
// @Router			/api/transaction/publish [post]
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

// @Summary		Delete a transaction
// @Description	Delete a transaction
// @Tags			transaction
// @Accept			json
// @Produce		json
// @Param			id	path		int				true	"Transaction ID"
// @Success		200	{object}	common.Result	"OK"
// @Router			/api/transaction/delete/{id} [post]
func DeleteTransaction(ctx *gin.Context) {
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
	err = service.DeleteTransaction(user.ID, transactionID)
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

// @Summary		Censor a transaction
// @Description	Censor a transaction
// @Tags			transaction
// @Accept			json
// @Produce		json
// @Param			censorRequest	body		common.CensorRequest	true	"censor request"
// @Success		200				{object}	common.Result			"OK"
// @Router			/api/transaction/censor [post]
func CensorTransaction(ctx *gin.Context) {
	var request common.CensorRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(common.BAD_REQUEST, common.Bad(err.Error()))
		return
	}

	transactionID, isPassed := request.ID, request.IsPassed

	user := utils.GetLoginUser(ctx)
	if user == nil {
		ctx.JSON(common.UNAUTHORIZED, common.NoAuth(common.ErrorGetInfoFromToken.Error()))
		return
	}
	if user.Role != common.RoleAdmin {
		ctx.JSON(common.FORBIDDEN, common.Forbidden(common.ErrorNoPermission.Error()))
		return
	}

	err = service.CensorTransaction(isPassed, transactionID)
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
