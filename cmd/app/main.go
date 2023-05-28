package main

import (
	"log"

	"quiz-mtuci-server/config"
	"quiz-mtuci-server/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
