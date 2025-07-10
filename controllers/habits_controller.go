package controllers

import (
	"net/http"
	"strconv"

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

func GetHabits(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	habits, err := habits.FindAllByUserID(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, habits)
}

func DeleteHabit(c *gin.Context) {
	habitID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	user := c.MustGet("user").(models.User)

	if err := habits.Delete(uint(habitID), user.ID); err != nil {
		if err.Error() == "hábito não encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hábito deletado com sucesso"})
}

func UpdateHabit(c *gin.Context) {
	habitID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	user := c.MustGet("user").(models.User)

	updatedHabit := &models.Habit{}

	if err := c.ShouldBindJSON(updatedHabit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := habits.Update(uint(habitID), user.ID, updatedHabit); err != nil {
		if err.Error() == "hábito não encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Busca o hábito atualizado do banco para retornar dados completos
	habit, err := habits.FindByID(uint(habitID), user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar hábito atualizado"})
		return
	}

	c.JSON(http.StatusOK, habit)
}
