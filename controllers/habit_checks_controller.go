package controllers

import (
	"net/http"
	"strconv"

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

func Check(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	habitCheckID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := habit_checks.Check(uint(habitCheckID), user.ID); err != nil {
		if err.Error() == "check-in não encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Check-in realizado com sucesso"})
}
