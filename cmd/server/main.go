package main

import (
	"context"
	"flag"

	"github.com/carlos19960601/fiber-boilerplate/cmd/server/wire"
	"github.com/carlos19960601/fiber-boilerplate/internal/pkg/config"
	"github.com/rs/zerolog/log"
)

var (
	ConfigFile string
)

func init() {
	flag.StringVar(&ConfigFile, "conf", "./config/local.yaml", "配置文件路径, 例如: -conf ./config/local.yaml")
	flag.Parse()
}

func main() {
	cfg, err := config.ParseWithPath(ConfigFile)
	if err != nil {
		log.Fatal().Err(err).Str("配置文件", ConfigFile).Msg("读取配置文件")
	}
	app, cleanup, err := wire.NewWire(cfg)
	defer cleanup()
	if err != nil {
		log.Fatal().Err(err).Msg("New Wire")
	}

	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}
}
