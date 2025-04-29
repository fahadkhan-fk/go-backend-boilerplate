package main

import (
	"log"

	"go-backend-boilerplate/config"
	"go-backend-boilerplate/internal/db"
	"go-backend-boilerplate/internal/handler"
	"go-backend-boilerplate/internal/repo"
	"go-backend-boilerplate/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Load()
	client := db.Connect(cfg)

	sqlDB, err := client.DB()
	if err != nil {
		log.Fatalf("Failed to retrieve underlying DB: %v", err)
	}
	defer sqlDB.Close()

	app := fiber.New()

	userRepo := repo.NewUserRepo(client)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	api := app.Group("/api/v1")
	api.Post("/register", userHandler.Register)

	log.Printf("Server running on port %s...", cfg.ServerPort)
	log.Fatal(app.Listen(":" + cfg.ServerPort))
}
