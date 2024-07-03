package handler

import (
	pkgjwt "github.com/carlos19960601/fiber-boilerplate/internal/pkg/jwt"
	"github.com/gofiber/fiber/v3"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func GetUserIdFromCtx(c fiber.Ctx) string {
	claims := c.Locals("claims")
	if claims == nil {
		return ""
	}

	return (claims.(*pkgjwt.MyCustomClaims)).UserId
}
