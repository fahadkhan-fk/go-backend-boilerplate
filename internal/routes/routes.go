package routes

import (
	"go-backend-boilerplate/internal/handler"

	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes sets up all API routes.
func RegisterRoutes(app *fiber.App, deps *Dependencies) {

	// Initialize handlers with dependencies
	userHandler := handler.NewUserHandler(deps.UserService)

	api := app.Group("/api/v1")
	api.Post("/register", userHandler.Register)
}
