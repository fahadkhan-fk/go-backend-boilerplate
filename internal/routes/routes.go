package routes

import (
	"go-backend-boilerplate/internal/handler"
	"go-backend-boilerplate/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes sets up all API routes.
func RegisterRoutes(app *fiber.App, deps *Dependencies) {

	// Initialize handlers with dependencies
	userHandler := handler.NewUserHandler(deps.UserService)

	api := app.Group("/api/v1")
	api.Post("/register", userHandler.Register)
	api.Post("/login", userHandler.Login)

	users := api.Group("/users", middleware.JWTAuth(deps.JWTSecret))
	users.Get("/me", userHandler.Me)
}
