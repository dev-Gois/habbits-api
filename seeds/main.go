package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/dev-Gois/habbits-api/config"
	"github.com/dev-Gois/habbits-api/models"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {
	// Carregar variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis do sistema")
	}

	// Conectar ao banco de dados
	config.ConnectDB()
	db := config.DB

	// Verificar argumentos da linha de comando
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go [seed|clear]")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "seed":
		runSeeds(db)
	case "clear":
		clearDatabase(db)
	default:
		fmt.Println("Comando inválido. Use 'seed' ou 'clear'")
		os.Exit(1)
	}
}

func runSeeds(db *gorm.DB) {
	fmt.Println("🌱 Iniciando seeds...")

	// Limpar dados existentes
	clearDatabase(db)

	// Criar usuários
	users := createUsers(db, 5)
	fmt.Printf("✅ Criados %d usuários\n", len(users))

	// Criar hábitos para cada usuário
	var allHabits []models.Habit
	for _, user := range users {
		habits := createHabitsForUser(db, user.ID, 8)
		allHabits = append(allHabits, habits...)
	}
	fmt.Printf("✅ Criados %d hábitos\n", len(allHabits))

	// Criar habit checks para os últimos 60 dias
	totalChecks := createHabitChecks(db, allHabits, 60)
	fmt.Printf("✅ Criados %d habit checks\n", totalChecks)

	fmt.Println("🎉 Seeds concluídos com sucesso!")
}

func clearDatabase(db *gorm.DB) {
	fmt.Println("🧹 Limpando banco de dados...")

	// Deletar em ordem devido às foreign keys
	db.Exec("DELETE FROM habit_checks")
	db.Exec("DELETE FROM habits")
	db.Exec("DELETE FROM users")

	fmt.Println("✅ Banco de dados limpo")
}

func createUsers(db *gorm.DB, count int) []models.User {
	var users []models.User

	// Usuários específicos definidos
	usersList := []struct {
		Name  string
		Email string
	}{
		{"Pedro Feijó", "feijo@gmail.com"},
		{"Sextou da Cantina", "sextou@gmail.com"},
		{"Germano Fenner", "germano@gmail.com"},
		{"Douglas Saboia", "saboia@gmail.com"},
		{"Jose Henrique", "henrique@gmail.com"},
	}

	for i := 0; i < count && i < len(usersList); i++ {
		// Hash da senha padrão
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)

		user := models.User{
			Name:     usersList[i].Name,
			Email:    usersList[i].Email,
			Password: string(hashedPassword),
		}

		if err := db.Create(&user).Error; err != nil {
			log.Printf("Erro ao criar usuário: %v", err)
			continue
		}

		users = append(users, user)
	}

	return users
}

func createHabitsForUser(db *gorm.DB, userID uint, count int) []models.Habit {
	var habits []models.Habit

	habitNames := []string{
		"Beber 2L de água",
		"Exercitar-se por 30 minutos",
		"Meditar",
		"Ler por 1 hora",
		"Estudar programação",
		"Caminhar 10.000 passos",
		"Dormir 8 horas",
		"Comer frutas",
		"Fazer alongamento",
		"Escrever no diário",
		"Praticar gratidão",
		"Organizar o ambiente",
		"Tomar vitaminas",
		"Evitar redes sociais",
		"Fazer yoga",
	}

	// Ícones para os hábitos
	habitIcons := []string{
		"café", "halter", "coração", "livro aberto", "impressora", "folha", "cama", "maçã",
		"halter", "email", "halter", "casa", "coração", "telefone", "pena",
	}

	// Diferentes padrões de dias da semana
	weekdayPatterns := [][]bool{
		{false, true, true, true, true, true, false},    // Segunda a sexta
		{true, true, true, true, true, true, true},      // Todos os dias
		{false, true, false, true, false, true, false},  // Dias alternados
		{false, false, true, false, true, false, true},  // Terça, quinta, sábado
		{true, false, false, false, false, false, true}, // Fins de semana
		{false, true, true, true, true, true, true},     // Segunda a sábado
		{true, false, false, false, true, false, false}, // Segunda, quinta, domingo
	}

	for i := 0; i < count && i < len(habitNames); i++ {
		pattern := weekdayPatterns[rand.Intn(len(weekdayPatterns))]

		habit := models.Habit{
			UserID:    userID,
			Title:     habitNames[i],
			Icon:      habitIcons[i],
			Sunday:    pattern[0],
			Monday:    pattern[1],
			Tuesday:   pattern[2],
			Wednesday: pattern[3],
			Thursday:  pattern[4],
			Friday:    pattern[5],
			Saturday:  pattern[6],
		}

		if err := db.Create(&habit).Error; err != nil {
			log.Printf("Erro ao criar hábito: %v", err)
			continue
		}

		habits = append(habits, habit)
	}

	return habits
}

func createHabitChecks(db *gorm.DB, habits []models.Habit, daysBack int) int {
	totalChecks := 0
	now := time.Now()

	for _, habit := range habits {
		// Simular diferentes padrões de completude
		completionRate := getCompletionRateForHabit(habit.Title)

		for i := 0; i < daysBack; i++ {
			checkDate := now.AddDate(0, 0, -i)

			// Verificar se o hábito deve ser feito neste dia da semana
			if !shouldHabitBeCheckedOnDay(habit, checkDate) {
				continue
			}

			// Decidir se o hábito foi completado baseado na taxa de completude
			if rand.Float64() < completionRate {
				// Gerar horário aleatório para o check
				checkTime := generateRandomTimeForDay(checkDate)

				habitCheck := models.HabitCheck{
					HabitID: habit.ID,
					Done:    true,
					Date:    checkTime,
				}

				if err := db.Create(&habitCheck).Error; err != nil {
					log.Printf("Erro ao criar habit check: %v", err)
					continue
				}

				totalChecks++
			}
		}
	}

	return totalChecks
}

func getCompletionRateForHabit(habitName string) float64 {
	// Diferentes taxas de completude para tornar os dados mais realistas
	rates := map[string]float64{
		"Beber 2L de água":            0.85,
		"Exercitar-se por 30 minutos": 0.60,
		"Meditar":                     0.45,
		"Ler por 1 hora":              0.55,
		"Estudar programação":         0.70,
		"Caminhar 10.000 passos":      0.40,
		"Dormir 8 horas":              0.65,
		"Comer frutas":                0.80,
		"Fazer alongamento":           0.50,
		"Escrever no diário":          0.35,
		"Praticar gratidão":           0.75,
		"Organizar o ambiente":        0.30,
		"Tomar vitaminas":             0.90,
		"Evitar redes sociais":        0.25,
		"Fazer yoga":                  0.40,
	}

	if rate, exists := rates[habitName]; exists {
		return rate
	}

	// Taxa padrão para hábitos não mapeados
	return 0.60
}

func shouldHabitBeCheckedOnDay(habit models.Habit, date time.Time) bool {
	dayOfWeek := date.Weekday()

	switch dayOfWeek {
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
	}

	return false
}

func generateRandomTimeForDay(date time.Time) time.Time {
	// Gerar horários mais realistas baseados em padrões de comportamento
	hourDistribution := []int{
		6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 8, 8, 9, 9, 9, 9, // Manhã (6h-9h)
		10, 10, 11, 11, 12, 12, // Meio da manhã
		13, 13, 14, 14, 15, 15, // Tarde
		16, 16, 17, 17, 18, 18, 18, 19, 19, 19, // Final da tarde
		20, 20, 20, 21, 21, 21, 22, 22, 22, 23, // Noite
	}

	hour := hourDistribution[rand.Intn(len(hourDistribution))]
	minute := rand.Intn(60)
	second := rand.Intn(60)

	return time.Date(date.Year(), date.Month(), date.Day(), hour, minute, second, 0, date.Location())
}
