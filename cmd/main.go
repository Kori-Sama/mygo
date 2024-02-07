package main

import (
	"mygo/config"
	"mygo/controller"
	"mygo/middlewares"
	"mygo/model"

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

// @host	localhost:8888
func main() {
	config.InitLog()
	config.InitConfig()

	model.InitEngine()
	model.SyncTables()

	app := gin.Default()
	app.Use(middlewares.Cors())

	bcGroup := app.Group("/api/blockchain")
	{
		bcGroup.POST("/createWallet/:username/:passphrase", controller.CreateWallet)
		bcGroup.GET("/getBalance/:username", controller.GetBalance)
		bcGroup.POST("/transfer", controller.Transfer)
	}

	// app.POST("/api/blockchain/transfer", controller.Transfer)

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.Run(config.Server.Host + ":" + config.Server.Port)
}
