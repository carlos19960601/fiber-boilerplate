package v1

import "github.com/gofiber/fiber/v3"

var (
	// common errors
	ErrBadRequest          = newBizError(fiber.StatusBadRequest, "Bad Request")
	ErrUnauthorized        = newBizError(fiber.StatusUnauthorized, "Unauthorized")
	ErrNotFound            = newBizError(fiber.StatusNotFound, "Not Found")
	ErrInternalServerError = newBizError(fiber.StatusInternalServerError, "Internal Server Error")
)
