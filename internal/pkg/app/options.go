package app

import "github.com/carlos19960601/fiber-boilerplate/internal/pkg/server"

type Option func(a *App)

func WithServer(servers ...server.Server) Option {
	return func(a *App) {
		a.servers = servers
	}
}
