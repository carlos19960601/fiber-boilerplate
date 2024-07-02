package handler

import "github.com/gofiber/fiber/v3"

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Register(c fiber.Ctx) error {
	return c.SendString("ok")
}
