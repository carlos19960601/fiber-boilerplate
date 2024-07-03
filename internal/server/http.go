package server

import (
	"github.com/carlos19960601/fiber-boilerplate/internal/handler"
	"github.com/carlos19960601/fiber-boilerplate/internal/middleware"
	"github.com/carlos19960601/fiber-boilerplate/internal/pkg/config"
	"github.com/carlos19960601/fiber-boilerplate/internal/pkg/jwt"
	"github.com/carlos19960601/fiber-boilerplate/internal/pkg/server/http"
)

func NewHTTPServer(
	cfg *config.Config,
	jwt *jwt.JWT,
	userHandler *handler.UserHandler,
) *http.Server {

	s := http.NewServer(http.WithPort(cfg.HTTP.Port), http.WithHost(cfg.HTTP.Host))

	s.Use(middleware.CORSMiddleware())

	v1 := s.Group("/v1")
	{
		noAuthRouter := v1.Group("/")
		{
			noAuthRouter.Post("/register", userHandler.Register)
			noAuthRouter.Post("/login", userHandler.Login)
		}

		strictAuthRouter := v1.Group("/").Use(middleware.StrictAuth(jwt))
		{
			strictAuthRouter.Get("/profile", userHandler.GetProfile)
		}
	}

	return s
}
