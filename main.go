package main

import (
	"personal/TODO-list/database"
	"personal/TODO-list/pkg/setting"
	"personal/TODO-list/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("hello world")
	gin.SetMode(setting.ServerSetting.RunMode)

	setting.Setup()
	database.SetupDatabaseConnection()

	routersInit := router.SetupRouter()
	routersInit.Run(setting.ServerSetting.HttpPort)
}
