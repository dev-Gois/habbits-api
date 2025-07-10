package controllers

import (
	"net/http"

	"github.com/dev-Gois/habbits-api/models"
	"github.com/dev-Gois/habbits-api/services/habit_checks"
	"github.com/gin-gonic/gin"
)

func GetToday(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	habitChecks, err := habit_checks.GetAllToday(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, habitChecks)
}
