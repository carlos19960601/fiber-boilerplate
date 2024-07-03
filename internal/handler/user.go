package handler

import (
	v1 "github.com/carlos19960601/fiber-boilerplate/api/v1"
	"github.com/carlos19960601/fiber-boilerplate/internal/service"
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog/log"
)

type UserHandler struct {
	*Handler
	userService service.UserService
}

func NewUserHandler(handler *Handler, userService service.UserService) *UserHandler {
	return &UserHandler{
		Handler:     handler,
		userService: userService,
	}
}

func (h *UserHandler) Register(c fiber.Ctx) error {
	req := new(v1.RegisterRequest)
	if err := c.Bind().JSON(req); err != nil {
		return v1.HandleError(c, v1.ErrBadRequest)
	}
	if err := h.userService.Register(c.Context(), req); err != nil {
		log.Error().Err(err).Msg("User Register")
		return v1.HandleError(c, err)
	}

	return v1.HandleSuccess(c, nil)
}

func (h *UserHandler) Login(c fiber.Ctx) error {
	var req v1.LoginRequest
	if err := c.Bind().JSON(&req); err != nil {
		return v1.HandleError(c, v1.ErrBadRequest)
	}

	token, err := h.userService.Login(c.Context(), &req)
	if err != nil {
		return v1.HandleError(c, v1.ErrUnauthorized)
	}

	return v1.HandleSuccess(c, v1.LoginResponseData{
		AccessToken: token,
	})
}

func (h *UserHandler) GetProfile(c fiber.Ctx) error {
	userId := GetUserIdFromCtx(c)
	if userId == "" {
		return v1.HandleError(c, v1.ErrUnauthorized)
	}

	user, err := h.userService.GetProfile(c.Context(), userId)
	if err != nil {
		return v1.HandleError(c, v1.ErrBadRequest)
	}

	return v1.HandleSuccess(c, user)
}
