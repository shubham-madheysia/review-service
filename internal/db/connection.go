package db

import (
	"fmt"
	"log"
	"reviewservice/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/cenkalti/backoff/v4"
)

// Connect establishes a connection to PostgreSQL with retry and returns *gorm.DB
func Connect(cfg *config.Config) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	operation := func() error {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Printf("Failed to connect to DB: %v", err)
			return err
		}

		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Failed to get sql.DB from gorm DB: %v", err)
			return err
		}

		// Configure connection pool
		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetConnMaxLifetime(time.Hour)

		return nil
	}

	// Retry with exponential backoff for up to ~30 seconds
	backOffConfig := backoff.NewExponentialBackOff()
	backOffConfig.MaxElapsedTime = 30 * time.Second

	if err = backoff.Retry(operation, backOffConfig); err != nil {
		return nil, fmt.Errorf("could not connect to DB after retries: %w", err)
	}

	log.Println("Connected to PostgreSQL database.")
	return db, nil
}
