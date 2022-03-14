package main

import (
	"os"
	"strconv"

	"github.com/stevenwilkin/buyer/binance"

	_ "github.com/joho/godotenv/autoload"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

var (
	schedule string
	usdt     float64
)

func main() {
	if schedule = os.Getenv("SCHEDULE"); schedule == "" {
		log.Fatal("No cron schedule")
	}

	if usdt, _ = strconv.ParseFloat(os.Getenv("USDT"), 64); usdt <= 10.0 {
		log.Fatal("USDT quantity must be greater than 10")
	}

	log.WithFields(log.Fields{
		"usdt":     usdt,
		"schedule": schedule,
	}).Info("Starting")

	b := &binance.Binance{
		ApiKey:    os.Getenv("BINANCE_API_KEY"),
		ApiSecret: os.Getenv("BINANCE_API_SECRET")}

	c := cron.New()
	c.AddFunc(schedule, func() {
		log.WithField("usdt", usdt).Info("Buying")

		if err := b.Buy(usdt); err != nil {
			log.Error(err)
		}
	})
	c.Run()
}
