package jobs

import (
	"log"
	"time"

	"github.com/dev-Gois/habbits-api/config"
	"github.com/dev-Gois/habbits-api/models"
	"github.com/dev-Gois/habbits-api/services/habit_checks"
)

func CreateDailyHabitChecks() {
	log.Println("Executando: Criar check-ins diários")

	var habits []models.Habit
	today := time.Now()
	weekday := today.Weekday()

	err := config.DB.Where("deleted_at IS NULL").Find(&habits).Error
	if err != nil {
		log.Printf("Erro ao buscar hábitos: %v", err)
		return
	}

	createdCount := 0
	for _, habit := range habits {
		if isHabitActiveToday(&habit, weekday) {
			exists, err := habitCheckExistsToday(habit.ID, today)
			if err != nil {
				log.Printf("Erro ao verificar check-in existente: %v", err)
				continue
			}

			if !exists {
				if err := habit_checks.Create(habit.ID); err != nil {
					log.Printf("Erro ao criar check-in para hábito %d: %v", habit.ID, err)
					continue
				}
				createdCount++
			}
		}
	}

	log.Printf("Criados %d check-ins diários", createdCount)
}

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

func habitCheckExistsToday(habitID uint, today time.Time) (bool, error) {
	var count int64
	todayStr := today.Format("2006-01-02")

	err := config.DB.Model(&models.HabitCheck{}).
		Where("habit_id = ? AND date = ?", habitID, todayStr).
		Count(&count).Error

	return count > 0, err
}
