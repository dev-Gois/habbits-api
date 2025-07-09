package main

import (
	"github.com/dev-Gois/habbits-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":3000")
}