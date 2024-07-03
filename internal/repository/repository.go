package repository

import (
	"context"
	"time"

	"github.com/carlos19960601/fiber-boilerplate/internal/pkg/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TxKey string

const ctxTxKey TxKey = "TxKey"

type Transaction interface {
	Transaction(ctx context.Context, fn func(ctx context.Context) error) error
}

type Repository struct {
	db *gorm.DB
	//rdb    *redis.Client
}

func NewRepository(
	db *gorm.DB,
	// rdb *redis.Client,
) *Repository {
	return &Repository{
		db: db,
		//rdb:    rdb,
	}
}

func NewTransaction(r *Repository) Transaction {
	return r
}

func (r *Repository) DB(ctx context.Context) *gorm.DB {
	v := ctx.Value(ctxTxKey)
	if v != nil {
		if tx, ok := v.(*gorm.DB); ok {
			return tx
		}
	}
	return r.db.WithContext(ctx)
}

func (r *Repository) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, ctxTxKey, tx)
		return fn(ctx)
	})
}

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
