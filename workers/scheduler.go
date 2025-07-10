package workers

import (
	"log"
	"time"

	"github.com/dev-Gois/habbits-api/config"
	"github.com/dev-Gois/habbits-api/models"
	"github.com/robfig/cron/v3"
)

var scheduler *cron.Cron

// InitScheduler inicializa o scheduler de cron jobs
func InitScheduler() {
	scheduler = cron.New(cron.WithLocation(time.UTC))

	// Adicionar jobs
	addJobs()

	// Iniciar o scheduler
	scheduler.Start()
	log.Println("Cron scheduler iniciado")
}

// addJobs adiciona todos os cron jobs
func addJobs() {
	// Criar check-ins diários para hábitos ativos (todo dia à 00:00)
	scheduler.AddFunc("0 0 * * *", func() {
		createDailyHabitChecks(time.Now())
	})

	// Limpar habit_checks antigos (todo domingo à 02:00)
	scheduler.AddFunc("0 2 * * 0", cleanOldHabitChecks)

	log.Println("Cron jobs registrados:")
	log.Println("- Criar check-ins diários: 0 0 * * * (todo dia à 00:00)")
	log.Println("- Limpar check-ins antigos: 0 2 * * 0 (todo domingo à 02:00)")
}

// CreateDailyHabitChecksForDate cria check-ins para uma data específica
func CreateDailyHabitChecksForDate(targetDate time.Time) (int, error) {
	log.Printf("Executando: Criar check-ins para %s", targetDate.Format("2006-01-02"))

	var habits []models.Habit
	weekday := targetDate.Weekday()

	// Buscar hábitos ativos que devem ter check-in na data especificada
	err := config.DB.Where("deleted_at IS NULL").Find(&habits).Error
	if err != nil {
		log.Printf("Erro ao buscar hábitos: %v", err)
		return 0, err
	}

	createdCount := 0
	for _, habit := range habits {
		// Verificar se o hábito está ativo para o dia da semana
		if isHabitActiveToday(&habit, weekday) {
			// Verificar se já existe check-in para a data
			exists, err := habitCheckExistsForDate(habit.ID, targetDate)
			if err != nil {
				log.Printf("Erro ao verificar check-in existente: %v", err)
				continue
			}

			if !exists {
				// Criar check-in para a data específica
				if err := createHabitCheckForDate(habit.ID, targetDate); err != nil {
					log.Printf("Erro ao criar check-in para hábito %d: %v", habit.ID, err)
					continue
				}
				createdCount++
			}
		}
	}

	log.Printf("Criados %d check-ins para %s", createdCount, targetDate.Format("2006-01-02"))
	return createdCount, nil
}

// createDailyHabitChecks cria check-ins para hábitos ativos do dia (versão automática)
func createDailyHabitChecks(today time.Time) {
	CreateDailyHabitChecksForDate(today)
}

// cleanOldHabitChecks remove habit_checks antigos (mais de 30 dias)
func cleanOldHabitChecks() {
	log.Println("Executando: Limpar check-ins antigos")

	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)

	result := config.DB.Where("date < ?", thirtyDaysAgo).Delete(&models.HabitCheck{})
	if result.Error != nil {
		log.Printf("Erro ao limpar check-ins antigos: %v", result.Error)
		return
	}

	log.Printf("Removidos %d check-ins antigos", result.RowsAffected)
}

// isHabitActiveToday verifica se o hábito está ativo para o dia da semana
func isHabitActiveToday(habit *models.Habit, weekday time.Weekday) bool {
	switch weekday {
	case time.Sunday:
		return habit.Sunday
	case time.Monday:
		return habit.Monday
	case time.Tuesday:
		return habit.Tuesday
	case time.Wednesday:
		return habit.Wednesday
	case time.Thursday:
		return habit.Thursday
	case time.Friday:
		return habit.Friday
	case time.Saturday:
		return habit.Saturday
	default:
		return false
	}
}

// habitCheckExistsForDate verifica se já existe check-in para uma data específica
func habitCheckExistsForDate(habitID uint, targetDate time.Time) (bool, error) {
	var count int64
	dateStr := targetDate.Format("2006-01-02")

	err := config.DB.Model(&models.HabitCheck{}).
		Where("habit_id = ? AND date = ?", habitID, dateStr).
		Count(&count).Error

	return count > 0, err
}

// createHabitCheckForDate cria um check-in para uma data específica
func createHabitCheckForDate(habitID uint, targetDate time.Time) error {
	habitCheck := &models.HabitCheck{
		HabitID: habitID,
		Done:    false,
		Date:    targetDate,
	}

	return habitCheck.Create()
}

// StopScheduler para o scheduler de cron jobs
func StopScheduler() {
	if scheduler != nil {
		scheduler.Stop()
		log.Println("Cron scheduler parado")
	}
}
