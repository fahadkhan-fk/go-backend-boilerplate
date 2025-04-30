package main

import (
	"log"

	"go-backend-boilerplate/config"
	"go-backend-boilerplate/internal/db"
	"go-backend-boilerplate/internal/repo"
	"go-backend-boilerplate/internal/routes"
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

	// Initialize repositories
	userRepo := repo.NewUserRepo(client)

	// Initialize services
	userService := service.NewUserService(userRepo)

	// Dependency container
	deps := &routes.Dependencies{
		UserService: userService,
	}

	// Register all routes with dependencies
	routes.RegisterRoutes(app, deps)

	log.Printf("Server running on port %s...", cfg.ServerPort)
	log.Fatal(app.Listen(":" + cfg.ServerPort))
}
