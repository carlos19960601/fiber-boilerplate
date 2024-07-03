// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/carlos19960601/fiber-boilerplate/internal/handler"
	"github.com/carlos19960601/fiber-boilerplate/internal/pkg/app"
	"github.com/carlos19960601/fiber-boilerplate/internal/pkg/config"
	"github.com/carlos19960601/fiber-boilerplate/internal/pkg/jwt"
	"github.com/carlos19960601/fiber-boilerplate/internal/pkg/server/http"
	"github.com/carlos19960601/fiber-boilerplate/internal/pkg/sid"
	"github.com/carlos19960601/fiber-boilerplate/internal/repository"
	"github.com/carlos19960601/fiber-boilerplate/internal/server"
	"github.com/carlos19960601/fiber-boilerplate/internal/service"
	"github.com/google/wire"
)

// Injectors from wire.go:

func NewWire(cfg *config.Config) (*app.App, func(), error) {
	jwtJWT := jwt.NewJwt(cfg)
	handlerHandler := handler.NewHandler()
	db := repository.NewDB(cfg)
	repositoryRepository := repository.NewRepository(db)
	transaction := repository.NewTransaction(repositoryRepository)
	sidSid := sid.NewSid()
	serviceService := service.NewService(transaction, sidSid, jwtJWT)
	userRepository := repository.NewUserRepository(repositoryRepository)
	userService := service.NewUserService(serviceService, userRepository)
	userHandler := handler.NewUserHandler(handlerHandler, userService)
	httpServer := server.NewHTTPServer(cfg, jwtJWT, userHandler)
	appApp := newApp(httpServer)
	return appApp, func() {
	}, nil
}

// wire.go:

var handlerSet = wire.NewSet(handler.NewHandler, handler.NewUserHandler)

var serviceSet = wire.NewSet(service.NewService, service.NewUserService)

var repositorySet = wire.NewSet(repository.NewDB, repository.NewRepository, repository.NewTransaction, repository.NewUserRepository)

var serverSet = wire.NewSet(server.NewHTTPServer)

func newApp(httpServer *http.Server) *app.App {
	return app.NewApp(app.WithServer(httpServer))
}
