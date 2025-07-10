package controllers

import (
	"net/http"
	"time"

	"github.com/dev-Gois/habbits-api/workers"
	"github.com/gin-gonic/gin"
)

// CreateHabitChecksForDate executa o job de criar check-ins para uma data específica
func CreateHabitChecksForDate(c *gin.Context) {
	var targetDate time.Time
	var err error

	// Primeiro tenta pegar do query parameter
	dateStr := c.Query("date")

	// Se não tiver no query, tenta pegar do body JSON
	if dateStr == "" {
		var requestBody struct {
			Date string `json:"date"`
		}

		if err := c.ShouldBindJSON(&requestBody); err == nil && requestBody.Date != "" {
			dateStr = requestBody.Date
		}
	}

	// Se ainda não tiver data, usa hoje
	if dateStr == "" {
		targetDate = time.Now()
	} else {
		targetDate, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Formato de data inválido. Use YYYY-MM-DD",
			})
			return
		}
	}

	// Executar o job
	createdCount, err := workers.CreateDailyHabitChecksForDate(targetDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Job executado com sucesso",
		"date":    targetDate.Format("2006-01-02"),
		"created": createdCount,
	})
}
