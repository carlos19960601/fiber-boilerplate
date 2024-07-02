package server

import (
	"github.com/carlos19960601/fiber-boilerplate/internal/handler"
	"github.com/carlos19960601/fiber-boilerplate/internal/middleware"
	"github.com/carlos19960601/fiber-boilerplate/internal/pkg/config"
	"github.com/carlos19960601/fiber-boilerplate/internal/pkg/server/http"
)

func NewHTTPServer(
	cfg *config.Config,
	userHandler *handler.UserHandler,
) *http.Server {

	s := http.NewServer(http.WithPort(cfg.HTTP.Port), http.WithHost(cfg.HTTP.Host))

	s.Use(middleware.CORSMiddleware())

	v1 := s.Group("/v1")
	{
		v1.Get("/register", userHandler.Register)
	}

	return s
}
