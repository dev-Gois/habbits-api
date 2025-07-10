package routes

import (
	"github.com/dev-Gois/habbits-api/config"
	"github.com/dev-Gois/habbits-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	config.ConnectDB()
	api := router.Group("/api")
	{
		api.GET("/", controllers.GetApplication)
		api.POST("/register", controllers.CreateUser)
		api.POST("/login", controllers.Login)
	}
}