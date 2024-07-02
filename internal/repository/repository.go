package repository

import (
	"time"

	"github.com/carlos19960601/fiber-boilerplate/internal/pkg/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(cfg *config.Config) *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)

	driver := cfg.DB.Dirver
	dsn := cfg.DB.DSN

	switch driver {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	default:
		log.Fatal().Msg("gorm不支持的driver")
	}
	if err != nil {
		log.Fatal().Err(err).Msg("gorm初始化失败")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("gorm初始化失败")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
