package routes

import (
	"github.com/dev-Gois/habbits-api/config"
	"github.com/dev-Gois/habbits-api/controllers"
	"github.com/dev-Gois/habbits-api/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	config.ConnectDB()
	api := router.Group("/api")
	{
		api.GET("/", controllers.GetApplication)
		api.POST("/register", controllers.CreateUser)
		api.POST("/login", controllers.Login)
		api.GET("/user", middlewares.Authorization(), controllers.GetUser)
		api.POST("/habits", middlewares.Authorization(), controllers.CreateHabit)
		api.GET("/habits", middlewares.Authorization(), controllers.GetHabits)
		api.DELETE("/habits/:id", middlewares.Authorization(), controllers.DeleteHabit)
		api.PUT("/habits/:id", middlewares.Authorization(), controllers.UpdateHabit)
		api.GET("/habit-checks", middlewares.Authorization(), controllers.GetToday)
		api.PUT("/habit-checks/:id/check", middlewares.Authorization(), controllers.Check)
	}
}
