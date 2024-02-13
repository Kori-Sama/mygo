package controller

import (
	"mygo/pkg/common"
	"mygo/pkg/constants"
	"mygo/pkg/utils"
	"mygo/service"

	"github.com/gin-gonic/gin"
)

// @Summary		login
// @Description	login
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			loginRequest	body		common.LoginRequest	true	"login request"
// @Success		200				{object}	common.Result		"OK"
// @Router			/api/login [post]
func Login(ctx *gin.Context) {
	var loginRequest common.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(400, common.Bad(err.Error()))
		return
	}
	username, password := loginRequest.Username, loginRequest.Password

	id, err := service.Login(username, password)
	if err != nil {
		if common.CheckInternalError(err) {
			ctx.JSON(500, common.InternalError(err.Error()))
			return
		}
		ctx.JSON(400, common.Bad(err.Error()))
		return
	}

	token, err := utils.GenerateToken(id, username)
	if err != nil {
		ctx.JSON(500, common.InternalError(err.Error()))
		return
	}

	ctx.Header(constants.TOKEN_NAME, constants.TOKEN_PREFIX+token)
	ctx.JSON(200, common.Ok(nil))
}

// @Summary		register
// @Description	register
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			registerRequest	body		common.RegisterRequest	true	"register request"
// @Success		200				{object}	common.Result			"OK"
// @Router			/api/register [post]
func Register(ctx *gin.Context) {
	var registerRequest common.RegisterRequest
	if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
		ctx.JSON(400, common.Bad(err.Error()))
		return
	}
	username, password := registerRequest.Username, registerRequest.Password

	id, err := service.Register(username, password)
	if err != nil {
		if common.CheckInternalError(err) {
			ctx.JSON(500, common.InternalError(err.Error()))
			return
		}
		ctx.JSON(400, common.Bad(err.Error()))
		return
	}

	token, err := utils.GenerateToken(id, username)
	if err != nil {
		ctx.JSON(500, common.InternalError(err.Error()))
		return
	}

	ctx.Header(constants.TOKEN_NAME, constants.TOKEN_PREFIX+token)
	ctx.JSON(200, common.Ok(nil))
}
