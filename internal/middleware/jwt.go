package middleware

import (
	"github.com/carlos19960601/fiber-boilerplate/api/v1"
	"github.com/gofiber/fiber/v3"
)

func StrictAuth() fiber.Handler {
	return func(c fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return v1.HandleError(c, v1.ErrUnauthorized)
		}
		return c.Next()
	}
}
