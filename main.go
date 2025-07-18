package main

import (
	"os"

	"github.com/dev-Gois/habbits-api/config"
	"github.com/dev-Gois/habbits-api/models"
	"github.com/dev-Gois/habbits-api/routes"
	"github.com/dev-Gois/habbits-api/workers"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{}, &models.Habit{}, &models.HabitCheck{})

	// Inicializar cron scheduler
	workers.InitScheduler()

	router := gin.Default()

	router.Use(config.SetupCors())

	routes.SetupRoutes(router)
	router.Run(":" + os.Getenv("PORT"))
}
