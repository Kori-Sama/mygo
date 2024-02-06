package main

import (
	"mygo/config"
	"mygo/internal/controller"
	"mygo/internal/model"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitLog()
	config.InitConfig()

	model.InitEngine()
	model.SyncTables()

	app := gin.Default()

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	app.POST("/api/blockchain/createWallet", controller.CreateWallet)

	// app.POST("/api/blockchain/transfer", controller.Transfer)

	app.Run(config.Server.Host + ":" + config.Server.Port)
}
