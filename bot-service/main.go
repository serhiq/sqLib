package main

import (
	"bot-service/app"
	"bot-service/config"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	log.Info("Initializing bot...")

	cfg, err := config.New()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	an, err := app.New(cfg)
	if err != nil {
		log.Error(err)
		log.Fatal("Cannot initialize %s", "An Bot")
		os.Exit(1)
	}

	an.Start()
}
