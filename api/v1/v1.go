package v1

import (
	"github.com/gofiber/fiber/v3"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleError(c fiber.Ctx, err error) error {
	resp := &Response{}
	if bizErr, ok := err.(*BizError); ok {
		resp.Message = bizErr.Error()
		resp.Code = bizErr.Code
		return c.Status(bizErr.Status).JSON(resp)
	}

	resp.Message = ErrInternalServerError.Error()
	resp.Code = ErrInternalServerError.Code

	return c.Status(fiber.StatusInternalServerError).JSON(resp)
}

func HandleSuccess(c fiber.Ctx, data any) error {
	if data == nil {
		data = map[string]any{}
	}

	resp := &Response{
		Code:    ErrSuccess.Code,
		Message: ErrSuccess.Message,
		Data:    data,
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

type BizError struct {
	Code    int
	Status  int
	Message string
}

func (e *BizError) Error() string {
	return e.Message
}

func NewBizError(status, code int, message string) *BizError {
	return &BizError{
		Code:    code,
		Status:  status,
		Message: message,
	}
}
