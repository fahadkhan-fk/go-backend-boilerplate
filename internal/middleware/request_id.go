package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const RequestIDKey = "request_id"

func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqID := c.Get("X-Request-ID")

		if reqID == "" {
			reqID = uuid.New().String()
		}

		// Store in context
		c.Locals(RequestIDKey, reqID)

		// Set header for downstream services and response
		c.Set("X-Request-ID", reqID)

		return c.Next()
	}
}

func GetRequestID(c *fiber.Ctx) string {
	id, ok := c.Locals(RequestIDKey).(string)
	if !ok {
		return ""
	}
	return id
}
