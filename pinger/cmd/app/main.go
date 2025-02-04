package main

import (
	"log"
	"os"
	"pinger/internal/services"
	"time"
)

func main() {
	backendURL := os.Getenv("BACKEND_URL")
	if backendURL == "" {
		log.Fatal("BACKEND_URL не задан")
	}

	interval := os.Getenv("PINGER_INTERVAL_SECONDS")
	if interval == "" {
		log.Fatal("PINGER_INTERVAL_SECONDS не задан")
	}

	inter, err := time.ParseDuration(interval)
	if err != nil {
		log.Fatal("PINGER_INTERVAL_SECONDS неправильно задан")
	}

	ticker := time.NewTicker(inter)
	defer ticker.Stop()

	service := services.NewService(backendURL)

	for {
		service.RunPingCycle()
		<-ticker.C
	}
}
