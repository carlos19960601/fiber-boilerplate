package middleware

import (
	v1 "github.com/carlos19960601/fiber-boilerplate/api/v1"
	"github.com/carlos19960601/fiber-boilerplate/internal/pkg/jwt"
	"github.com/gofiber/fiber/v3"
)

func StrictAuth(j *jwt.JWT) fiber.Handler {
	return func(c fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return v1.HandleError(c, v1.ErrUnauthorized)
		}

		claims, err := j.ParseToken(tokenString)
		if err != nil {
			return v1.HandleError(c, v1.ErrUnauthorized)
		}

		c.Locals("claims", claims)
		return c.Next()
	}
}
