package database

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBClient *gorm.DB

func InitDB() *gorm.DB {
	if DBClient != nil {
		return DBClient
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.dbname"),
	)

	gormLogLevel := logger.Warn
	if viper.GetString("log.level") == "debug" {
		gormLogLevel = logger.Info
	} else if viper.GetString("log.level") == "silent" {
		gormLogLevel = logger.Silent
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                   logger.Default.LogMode(gormLogLevel),
		PrepareStmt:              false,
		DisableNestedTransaction: true,
		CreateBatchSize:          100,
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Database migrations are handled by Goose
	// Run: go run ./cmd/migrate -command=up -dir=./migrations
	// Or: ./migrate.sh

	DBClient = db
	return db
}

func WithTx(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return DBClient.WithContext(ctx).Transaction(fn)
}

// GetDB returns the transaction from context if available, otherwise the default DB client
func GetDB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value("tx").(*gorm.DB)
	if ok {
		return tx
	}
	return DBClient.WithContext(ctx)
}

type TransactionManager interface {
	RunInTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

type GormTransactionManager struct {
	db *gorm.DB
}

func NewGormTransactionManager(db *gorm.DB) *GormTransactionManager {
	return &GormTransactionManager{db: db}
}

func (m *GormTransactionManager) RunInTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Inject transaction into context
		txCtx := context.WithValue(ctx, "tx", tx)
		return fn(txCtx)
	})
}
