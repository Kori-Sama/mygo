package main

import (
	"mygo/config"
	_ "mygo/docs"
	"mygo/internal/server/controller"
	"mygo/internal/server/middlewares"
	"mygo/internal/server/model"
	"mygo/internal/server/service"
	"strings"

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

	service.Seg.LoadDict(strings.Join(config.Server.Dict, ","))

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
		// bc.POST("/transfer", controller.Transfer)
	}

	user := auth.Group("/user")
	{
		user.GET("/", controller.GetAllUsers)
		user.GET("/:id", controller.GetUser)
	}

	t := auth.Group("/transaction")
	{
		t.GET("/", controller.GetAllTransactions)
		t.GET("/:id", controller.GetTransaction)
		t.GET("/limit", controller.GetLimitedTransactions)
		t.GET("/self", controller.GetOwnTransactions)
		t.GET("/by", controller.GetTransactionByStatus)
		t.GET("/search", controller.SearchTransaction)
		t.POST("/new", controller.NewTransaction)
		t.POST("/save", controller.SaveTransaction)
		t.POST("/publish", controller.PublishTransaction)
		t.POST("/delete", controller.DeleteTransaction)
		t.POST("/censor", controller.CensorTransaction)
	}

	his := auth.Group("/history")
	{
		his.GET("/", controller.GetAllHistories)
		his.GET("/query", controller.QueryHistory)
	}

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.Run(":" + config.Server.Port)
}
