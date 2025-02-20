//go:build wireinject
// +build wireinject

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

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	//repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
)

func newApp(httpServer *http.Server) *app.App {
	return app.NewApp(
		app.WithServer(httpServer),
	)
}

func NewWire(cfg *config.Config) (*app.App, func(), error) {
	panic(wire.Build(
		serverSet,
		handlerSet,
		serviceSet,
		repositorySet,
		jwt.NewJwt,
		sid.NewSid,
		newApp,
	))
}
