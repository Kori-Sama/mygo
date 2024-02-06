package controller

import (
	"mygo/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// createWallet  godoc
// @Summary      create wallet
// @Description  create wallet in blockchain
// @Tags         blockchain
// @Accept       json
// @Produce      json
// @Param		 username	query	string	true	"username"
// @Param		 passphrase	query	string	true	"passphrase"
// @Success      200  {object}  nil "OK"
// @failure      400  {object}  string "Bad Request"
// @failure	  	 500  {object}  string "Internal Server Error"
// @Router       /api/blockchain/createWallet [post]
func CreateWallet(ctx *gin.Context) {
	username := ctx.Query("username")
	passphrase := ctx.Query("passphrase")
	if err := service.CreateWallet(username, passphrase); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// func Transfer(ctx *gin.Context) {

// }
