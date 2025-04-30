package routes

import (
	"go-backend-boilerplate/internal/service"
)

// Dependencies holds all injected services
type Dependencies struct {
	UserService service.UserService
}
