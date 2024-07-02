package app

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/carlos19960601/fiber-boilerplate/internal/pkg/server"
	"github.com/rs/zerolog/log"
)

type App struct {
	wg      sync.WaitGroup
	servers []server.Server
}

func NewApp(opts ...Option) *App {
	a := &App{}

	for _, opt := range opts {
		opt(a)
	}

	return a
}

func (a *App) Run(ctx context.Context) error {
	a.Start(ctx)
	a.waitForSignals()
	a.Shutdown(ctx)

	return nil
}

func (a *App) Start(ctx context.Context) {
	for _, srv := range a.servers {
		a.wg.Add(1)
		go func(srv server.Server) {
			defer a.wg.Done()
			err := srv.Start(ctx)
			if err != nil {
				log.Error().Err(err).Msg("HTTP服务启动")
			}
		}(srv)
	}
}

func (a *App) waitForSignals() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals
	log.Info().Msg("收到关闭的信号")
}

func (a *App) Shutdown(ctx context.Context) {
	for _, srv := range a.servers {
		srv.Stop(ctx)
	}
	a.wg.Wait()
	log.Info().Msg("服务退出")
}
