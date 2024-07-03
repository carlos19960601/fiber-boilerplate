//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/carlos19960601/fiber-boilerplate/internal/pkg/app"
	"github.com/carlos19960601/fiber-boilerplate/internal/pkg/config"
	"github.com/carlos19960601/fiber-boilerplate/internal/repository"
	"github.com/carlos19960601/fiber-boilerplate/internal/server"
	"github.com/google/wire"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	//repository.NewRedis,
	repository.NewRepository,
	repository.NewUserRepository,
)

var serverSet = wire.NewSet(
	server.NewMigrate,
)

func newApp(migrate *server.Migrate) *app.App {
	return app.NewApp(
		app.WithServer(migrate),
	)
}

func NewWire(cfg *config.Config) (*app.App, func(), error) {
	panic(wire.Build(
		serverSet,
		repositorySet,
		newApp,
	))
}
