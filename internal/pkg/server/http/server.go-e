package http

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog/log"
)

const (
	DefaultHost = "0.0.0.0"
	DefaultPort = 8000
)

type Server struct {
	*fiber.App
	host string
	port int
}

func NewServer(opts ...Option) *Server {
	s := &Server{
		App:  fiber.New(),
		host: DefaultHost,
		port: DefaultPort,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Server) Start(ctx context.Context) error {
	if err := s.App.Listen(fmt.Sprintf("%s:%d", s.host, s.port)); err != nil {
		log.Fatal().Err(err).Str("host", s.host).Int("port", s.port).Msg("监听端口")
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	log.Info().Msg("关闭服务器...")
	if err := s.App.ShutdownWithTimeout(5 * time.Second); err != nil {
		log.Fatal().Err(err).Msg("关闭服务失败")
	}

	log.Info().Msg("关闭服务器成功")

	return nil
}
