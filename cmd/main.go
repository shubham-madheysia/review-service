package main

import (
	"log"
	"os"
	"reviewservice/config"
	"reviewservice/internal/db"
	"reviewservice/internal/processor"
	"reviewservice/internal/s3"
)

// Run encapsulates the main workflow and returns errors instead of exiting directly.
// This allows unit tests to call Run() and check for errors.
func Run() error {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	// Connect to DB
	dbConn, err := db.Connect(cfg)
	if err != nil {
		return err
	}
	defer func() {
		if err := dbConn.Close(); err != nil {
			log.Printf("Warning: error closing DB connection: %v", err)
		}
	}()

	// Initialize S3 client
	s3Client, err := s3.NewClient(cfg)
	if err != nil {
		return err
	}

	// Run ingestion process
	if err := processor.RunIngestion(cfg, s3Client, dbConn); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatalf("Application error: %v", err)
	}
	log.Println("Review ingestion completed successfully.")
}
