package routes

import (
	"go-backend-boilerplate/internal/handler"
	"go-backend-boilerplate/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes sets up all API routes.
func RegisterRoutes(app *fiber.App, deps *Dependencies, secret string) {

	// Initialize handlers with dependencies
	userHandler := handler.NewUserHandler(deps.UserService)

	api := app.Group("/api/v1")
	api.Post("/register", userHandler.Register)
	api.Post("/login", userHandler.Login)

	users := api.Group("/users", middleware.JWTAuth(secret))
	users.Get("/me", userHandler.Me)
	users.Get("/", userHandler.List)
	users.Get("/:id", userHandler.Get)
	users.Patch("/:id", userHandler.Update)
	users.Delete("/:id", userHandler.Delete)
}
