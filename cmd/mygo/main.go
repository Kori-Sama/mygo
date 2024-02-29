package main

import (
	"mygo/config"
	_ "mygo/docs"
	"mygo/internal/server/controller"
	"mygo/internal/server/middlewares"
	"mygo/internal/server/model"

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
	config.InitConfig()
	config.InitLog()

	model.InitEngine()
	model.SyncTables()

	gin.SetMode(config.Server.Mode)

	app := gin.Default()
	app.Use(middlewares.Cors())

	auth := app.Group("/api")
	auth.Use(middlewares.JwtAuth())

	app.POST("/api/login", controller.Login)
	app.POST("/api/register", controller.Register)

	bc := auth.Group("/blockchain")
	{
		bc.POST("/createWallet/:passphrase", controller.CreateWallet)
		bc.GET("/getBalance", controller.GetBalance)
		bc.POST("/transfer", controller.Transfer)
	}

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.Run(":" + config.Server.Port)
}
