package workers

import (
	"log"
	"time"

	"github.com/dev-Gois/habbits-api/workers/jobs"
	"github.com/robfig/cron/v3"
)

var scheduler *cron.Cron

func InitScheduler() {
	scheduler = cron.New(cron.WithLocation(time.UTC))

	addJobs()

	scheduler.Start()
	log.Println("Cron scheduler iniciado")
}

func addJobs() {
	scheduler.AddFunc("0 0 * * *", jobs.CreateDailyHabitChecks)

	log.Println("Cron jobs registrados:")
	log.Println("- Criar check-ins diários: 0 0 * * * (todo dia à 00:00)")
}

func StopScheduler() {
	if scheduler != nil {
		scheduler.Stop()
		log.Println("Cron scheduler parado")
	}
}
