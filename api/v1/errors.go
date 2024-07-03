package v1

import "github.com/gofiber/fiber/v3"

var (
	// common errors
	ErrSuccess             = NewBizError(fiber.StatusOK, fiber.StatusOK, "success")
	ErrBadRequest          = NewBizError(fiber.StatusBadRequest, fiber.StatusBadRequest, "Bad Request")
	ErrUnauthorized        = NewBizError(fiber.StatusUnauthorized, fiber.StatusUnauthorized, "Unauthorized")
	ErrNotFound            = NewBizError(fiber.StatusNotFound, fiber.StatusNotFound, "Not Found")
	ErrInternalServerError = NewBizError(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "Internal Server Error")

	// more biz errors
	ErrEmailAlreadyUse = NewBizError(fiber.StatusBadRequest, 1001, "邮箱已使用")
)
