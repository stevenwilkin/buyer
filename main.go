package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

var (
	schedule string
)

func main() {
	if schedule = os.Getenv("SCHEDULE"); schedule == "" {
		log.Fatal("No cron schedule")
	}

	log.Info("Starting")

	c := cron.New()
	c.AddFunc(schedule, func() {
		log.Info("Doing it")
	})
	c.Run()
}
