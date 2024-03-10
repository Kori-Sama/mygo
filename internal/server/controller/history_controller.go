package controller

import (
	"mygo/internal/pkg/common"
	"mygo/internal/pkg/utils"
	"mygo/internal/server/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllHistories(ctx *gin.Context) {
	histories, err := service.GetAllHistory()
	if err != nil {
		if common.CheckInternalError(err) {
			ctx.JSON(common.INTERNAL_SERVER_ERROR, common.InternalError(err.Error()))
			return
		}
		ctx.JSON(common.BAD_REQUEST, common.Bad(err.Error()))
		return
	}
	ctx.JSON(common.OK, common.Ok(histories))
}

func QueryHistory(ctx *gin.Context) {
	userIdStr, ok := ctx.GetQuery("userId")
	if ok {
		id, err := strconv.Atoi(userIdStr)
		if err != nil {
			ctx.JSON(common.BAD_REQUEST, common.Bad(err.Error()))
			return
		}
		histories, err := service.GetHistoryByUserId(id)
		if err != nil {
			if common.CheckInternalError(err) {
				ctx.JSON(common.INTERNAL_SERVER_ERROR, common.InternalError(err.Error()))
				return
			}
			ctx.JSON(common.BAD_REQUEST, common.Bad(err.Error()))
			return
		}
		ctx.JSON(common.OK, common.Ok(histories))
	}

	transactionId, ok := ctx.GetQuery("transactionId")
	if ok {
		id, err := strconv.Atoi(transactionId)
		if err != nil {
			ctx.JSON(common.BAD_REQUEST, common.Bad(err.Error()))
			return
		}
		history, err := service.GetHistoryByTransactionId(id)
		if err != nil {
			if common.CheckInternalError(err) {
				ctx.JSON(common.INTERNAL_SERVER_ERROR, common.InternalError(err.Error()))
				return
			}
			ctx.JSON(common.BAD_REQUEST, common.Bad(err.Error()))
			return
		}
		ctx.JSON(common.OK, common.Ok(history))
	}

	actionStr, ok := ctx.GetQuery("action")
	if ok {
		action, err := utils.FilterAction(actionStr)
		if err != nil {
			ctx.JSON(common.BAD_REQUEST, common.Bad(err.Error()))
			return
		}
		histories, err := service.GetHistoryByAction(action)
		if err != nil {
			if common.CheckInternalError(err) {
				ctx.JSON(common.INTERNAL_SERVER_ERROR, common.InternalError(err.Error()))
				return
			}
			ctx.JSON(common.BAD_REQUEST, common.Bad(err.Error()))
			return
		}
		ctx.JSON(common.OK, common.Ok(histories))

	}
}

// TODO!
//
// func GetHistoryByTimestamp(ctx *gin.Context) {
// 	timestamp := ctx.Param("unix")
// 	time, err := strconv.Atoi(timestamp)
// 	if err != nil {
// 		ctx.JSON(common.BAD_REQUEST, common.Bad(err.Error()))
// 		return
// 	}
// 	histories, err := service.GetHistoryByTimestamp(time)
// 	if err != nil {
// 		if common.CheckInternalError(err) {
// 			ctx.JSON(common.INTERNAL_SERVER_ERROR, common.InternalError(err.Error()))
// 			return
// 		}
// 		ctx.JSON(common.BAD_REQUEST, common.Bad(err.Error()))
// 		return
// 	}
// 	ctx.JSON(common.OK, common.Ok(histories))
// }
