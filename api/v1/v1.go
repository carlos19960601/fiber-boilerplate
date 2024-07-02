package v1

import (
	"github.com/gofiber/fiber/v3"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleError(c fiber.Ctx, err error) error {
	resp := Response{}
	if bizErr, ok := err.(*BizError); ok {
		resp.Message = err.Error()
		return c.Status(bizErr.Status).JSON(resp)
	}

	return c.Status(fiber.StatusInternalServerError).JSON(resp)
}

type BizError struct {
	Status  int
	Message string
}

func (e BizError) Error() string {
	return e.Message
}

func newBizError(status int, message string) error {
	return BizError{
		Status:  status,
		Message: message,
	}
}
