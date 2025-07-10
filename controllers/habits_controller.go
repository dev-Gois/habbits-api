package controllers

import (
	"net/http"

	"github.com/dev-Gois/habbits-api/models"
	"github.com/dev-Gois/habbits-api/services/habits"
	"github.com/gin-gonic/gin"
)

func CreateHabit(c *gin.Context) {
	habit := &models.Habit{}

	if err := c.ShouldBindJSON(habit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := c.MustGet("user").(models.User)

	if err := habits.Create(habit, user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, habit)
}
