package main

import (
	"github.com/dev-Gois/habbits-api/config"
	"github.com/dev-Gois/habbits-api/models"
	"github.com/dev-Gois/habbits-api/routes"
	"github.com/dev-Gois/habbits-api/services/cron"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{}, &models.Habit{}, &models.HabitCheck{})

	// Inicializar cron scheduler
	cron.InitScheduler()

	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":3000")
}
