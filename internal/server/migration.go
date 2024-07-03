package server

import (
	"context"
	"os"

	"github.com/carlos19960601/fiber-boilerplate/internal/model"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Migrate struct {
	db *gorm.DB
}

func NewMigrate(db *gorm.DB) *Migrate {
	return &Migrate{
		db: db,
	}
}
func (m *Migrate) Start(ctx context.Context) error {
	if err := m.db.AutoMigrate(&model.User{}); err != nil {
		log.Error().Err(err).Msg("user migrate")
		return err
	}
	log.Info().Msg("AutoMigrate success")
	os.Exit(0)
	return nil
}
func (m *Migrate) Stop(ctx context.Context) error {
	log.Info().Msg("AutoMigrate stop")
	return nil
}
