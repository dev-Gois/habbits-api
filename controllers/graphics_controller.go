package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dev-Gois/habbits-api/config"
	"github.com/dev-Gois/habbits-api/models"
	"github.com/gin-gonic/gin"
)

type DashboardData struct {
	HourlyDistribution []HourlyData   `json:"hourly_distribution"`
	WeeklyAverage      []WeeklyData   `json:"weekly_average"`
	TimelineEvolution  []TimelineData `json:"timeline_evolution"`
}

type HourlyData struct {
	TimeRange string `json:"time_range"`
	Count     int    `json:"count"`
}

type WeeklyData struct {
	DayOfWeek string  `json:"day_of_week"`
	Average   float64 `json:"average"`
}

type TimelineData struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

func GetDashboardData(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	// Parâmetro opcional para definir quantos dias considerar (padrão: 30)
	daysParam := c.DefaultQuery("days", "30")
	days, err := strconv.Atoi(daysParam)
	if err != nil || days <= 0 {
		days = 30
	}

	// Data de início para análise
	startDate := time.Now().AddDate(0, 0, -days)

	// 1. Distribuição por horário do dia
	hourlyDistribution, err := getHourlyDistribution(user.ID, startDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar distribuição por horário"})
		return
	}

	// 2. Média por dia da semana
	weeklyAverage, err := getWeeklyAverage(user.ID, startDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar média semanal"})
		return
	}

	// 3. Evolução temporal
	timelineEvolution, err := getTimelineEvolution(user.ID, startDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar evolução temporal"})
		return
	}

	dashboardData := DashboardData{
		HourlyDistribution: hourlyDistribution,
		WeeklyAverage:      weeklyAverage,
		TimelineEvolution:  timelineEvolution,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Dados da dashboard obtidos com sucesso",
		"data":    dashboardData,
		"period": map[string]interface{}{
			"days":       days,
			"start_date": startDate.Format("2006-01-02"),
			"end_date":   time.Now().Format("2006-01-02"),
		},
	})
}

// getHourlyDistribution retorna a distribuição de conclusões por faixa horária
func getHourlyDistribution(userID uint, startDate time.Time) ([]HourlyData, error) {
	var results []struct {
		Hour  int `json:"hour"`
		Count int `json:"count"`
	}

	query := `
		SELECT 
			EXTRACT(HOUR FROM hc.updated_at) as hour,
			COUNT(*) as count
		FROM habit_checks hc
		JOIN habits h ON h.id = hc.habit_id
		WHERE h.user_id = ? 
			AND hc.done = true 
			AND hc.updated_at >= ?
			AND h.deleted_at IS NULL
		GROUP BY EXTRACT(HOUR FROM hc.updated_at)
		ORDER BY hour
	`

	err := config.DB.Raw(query, userID, startDate).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	// Agrupar em faixas horárias
	hourlyData := []HourlyData{
		{"00h–06h", 0},
		{"06h–09h", 0},
		{"09h–12h", 0},
		{"12h–15h", 0},
		{"15h–18h", 0},
		{"18h–21h", 0},
		{"21h–24h", 0},
	}

	for _, result := range results {
		hour := result.Hour
		count := result.Count

		switch {
		case hour >= 0 && hour < 6:
			hourlyData[0].Count += count
		case hour >= 6 && hour < 9:
			hourlyData[1].Count += count
		case hour >= 9 && hour < 12:
			hourlyData[2].Count += count
		case hour >= 12 && hour < 15:
			hourlyData[3].Count += count
		case hour >= 15 && hour < 18:
			hourlyData[4].Count += count
		case hour >= 18 && hour < 21:
			hourlyData[5].Count += count
		case hour >= 21 && hour < 24:
			hourlyData[6].Count += count
		}
	}

	return hourlyData, nil
}

// getWeeklyAverage retorna a média de conclusões por dia da semana
func getWeeklyAverage(userID uint, startDate time.Time) ([]WeeklyData, error) {
	var results []struct {
		DayOfWeek int     `json:"day_of_week"`
		Average   float64 `json:"average"`
	}

	query := `
		SELECT 
			EXTRACT(DOW FROM hc.date) as day_of_week,
			COUNT(*)::float / COUNT(DISTINCT DATE(hc.date)) as average
		FROM habit_checks hc
		JOIN habits h ON h.id = hc.habit_id
		WHERE h.user_id = ? 
			AND hc.done = true 
			AND hc.date >= ?
			AND h.deleted_at IS NULL
		GROUP BY EXTRACT(DOW FROM hc.date)
		ORDER BY day_of_week
	`

	err := config.DB.Raw(query, userID, startDate).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	// Mapear números dos dias para nomes
	dayNames := []string{"Domingo", "Segunda", "Terça", "Quarta", "Quinta", "Sexta", "Sábado"}
	weeklyData := make([]WeeklyData, 7)

	// Inicializar com zeros
	for i := 0; i < 7; i++ {
		weeklyData[i] = WeeklyData{
			DayOfWeek: dayNames[i],
			Average:   0,
		}
	}

	// Preencher com dados reais
	for _, result := range results {
		if result.DayOfWeek >= 0 && result.DayOfWeek < 7 {
			weeklyData[result.DayOfWeek].Average = result.Average
		}
	}

	return weeklyData, nil
}

// getTimelineEvolution retorna a evolução temporal de conclusões
func getTimelineEvolution(userID uint, startDate time.Time) ([]TimelineData, error) {
	var results []struct {
		Date  time.Time `json:"date"`
		Count int       `json:"count"`
	}

	query := `
		SELECT 
			hc.date,
			COUNT(*) as count
		FROM habit_checks hc
		JOIN habits h ON h.id = hc.habit_id
		WHERE h.user_id = ? 
			AND hc.done = true 
			AND hc.date >= ?
			AND h.deleted_at IS NULL
		GROUP BY hc.date
		ORDER BY hc.date
	`

	err := config.DB.Raw(query, userID, startDate).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	// Criar mapa para facilitar busca
	dataMap := make(map[string]int)
	for _, result := range results {
		dateStr := result.Date.Format("2006-01-02")
		dataMap[dateStr] = result.Count
	}

	// Preencher todos os dias no intervalo (incluindo dias sem dados)
	var timelineData []TimelineData
	currentDate := startDate
	endDate := time.Now()

	for currentDate.Before(endDate) || currentDate.Equal(endDate) {
		dateStr := currentDate.Format("2006-01-02")
		count := dataMap[dateStr] // Se não existir, será 0

		timelineData = append(timelineData, TimelineData{
			Date:  dateStr,
			Count: count,
		})

		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return timelineData, nil
}
