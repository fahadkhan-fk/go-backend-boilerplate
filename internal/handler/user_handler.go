package handler

import (
	"go-backend-boilerplate/internal/service"
	"go-backend-boilerplate/internal/validator"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{service: svc}
}

type registerRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req registerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := validator.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.service.Register(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req loginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := validator.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	token, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"token": token})
}

func (h *UserHandler) Me(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	user, err := h.service.GetById(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(fiber.Map{
		"id":    user.ID,
		"email": user.Email,
	})
}

func (h *UserHandler) List(c *fiber.Ctx) error {
	users, err := h.service.List()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func (h *UserHandler) Get(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.service.GetById(uint(id))
	if err != nil {
		return fiber.ErrNotFound
	}
	return c.JSON(user)
}

type updateUserRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	var req updateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	if err := validator.ValidateStruct(req); err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.service.Update(uint(id), req.Email); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{"message": "User updated successfully"})
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.service.Delete(uint(id)); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}
