package main

import (
	"mygo/config"
	"mygo/internal/controller"
	"mygo/internal/model"

	_ "mygo/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			MyGO!!!!! API docs
//	@version		0.1
//	@description	This is a API docs for MyGO project.

//	@contact.name	KoriSama
//	@contact.url	https://kori-sama.github.io/
//	@contact.email	Miyohashikori457@gmail.com

//	@host		localhost:8888
//	@BasePath	/api
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

	app.POST("/api/blockchain/createWallet/:username/:passphrase", controller.CreateWallet)

	// app.POST("/api/blockchain/transfer", controller.Transfer)

	app.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.Run(config.Server.Host + ":" + config.Server.Port)
}
