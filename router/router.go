package router

import (
	"personal/TODO-list/pkg/setting"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.Use(location.Default())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     setting.CorsSetting.AllowOrigin,
		AllowMethods:     setting.CorsSetting.AllowMethods,
		AllowHeaders:     setting.CorsSetting.AllowHeaders,
		ExposeHeaders:    setting.CorsSetting.ExposeHeaders,
		AllowCredentials: setting.CorsSetting.AllowCredentials,
	}))

	// r.POST("/checkFeeTally", controller.CheckFeeTally)

	return r
}
