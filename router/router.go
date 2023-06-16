package router

import (
	"personal/TODO-list/controller"
	"personal/TODO-list/middleware"
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

	r.Use(middleware.ApiLogger())

	r.POST("/login", controller.Login)

	// fb callback
	r.GET("/auth/facebook/callback", controller.FacebookCallbackHandler)
	// gmail callback
	r.GET("/auth/google/callback", controller.GoogleCallbackHandler)
	// github callback
	r.GET("/auth/github/callback", controller.GithubCallbackHandler)

	r.POST("/logout", middleware.AuthorizeJWT(), controller.Logout)

	toDoRoutes := r.Group("/api/todo", middleware.AuthorizeJWT())
	{
		toDoRoutes.POST("/addTodo", controller.AddTodo)
		toDoRoutes.DELETE("/deleteTodo", controller.DeleteTodo)
		toDoRoutes.POST("/listTodo", controller.ListTodo)
		toDoRoutes.PUT("/markTodo", controller.MarkTodo)
	}

	return r
}
