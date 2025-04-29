package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"go-backend-boilerplate/config"
)

func Connect(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DatabaseHost,
		cfg.DatabasePort,
		cfg.DatabaseUsername,
		cfg.DatabaseDb,
		cfg.DatabasePassword,
		cfg.DatabaseSslMode,
	)

	var err error
	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return client
}
