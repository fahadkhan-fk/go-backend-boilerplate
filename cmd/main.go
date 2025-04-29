package main

import (
	"go-backend-boilerplate/config"
	"go-backend-boilerplate/internal/db"
	"log"
)

func main() {
	cfg := config.Load()
	client := db.Connect(cfg)
	sqlDB, err := client.DB()
	if err != nil {
		log.Fatalf("Failed to retrieve underlying DB: %v", err)
	}
	defer sqlDB.Close()

	log.Println("Application initialized successfully.")

}
