package app

import (
    "fmt"
    "log"
    "review-service/internal/config"
    "review-service/internal/db"
    "review-service/internal/parser"
    "review-service/internal/processor"
    "review-service/internal/s3"
    "review-service/internal/validator"
)

type App struct {
    cfg       *config.Config
    s3Client  *s3.Client
    repo      *db.Repository
    parser    *parser.Parser
    validator *validator.Validator
    processor *processor.Processor
}

func NewApp(cfg *config.Config) (*App, error) {
    s3Client, err := s3.NewClient(cfg)
    if err != nil {
        return nil, fmt.Errorf("failed to create S3 client: %w", err)
    }

    repo, err := db.NewRepository(cfg.DB.DSN)
    if err != nil {
        return nil, fmt.Errorf("failed to create repository: %w", err)
    }

    return &App{
        cfg:       cfg,
        s3Client:  s3Client,
        repo:      repo,
        parser:    parser.NewParser(),
        validator: validator.NewValidator(),
        processor: processor.NewProcessor(repo),
    }, nil
}

// Run starts the ingestion and processing workflow
func (a *App) Run() error {
    files, err := a.s3Client.ListNewFiles(a.cfg.S3.Bucket, a.cfg.S3.Prefix)
    if err != nil {
        return fmt.Errorf("failed to list S3 files: %w", err)
    }

    for _, fileKey := range files {
        data, err := a.s3Client.GetFile(a.cfg.S3.Bucket, fileKey)
        if err != nil {
            log.Printf("failed to download file %s: %v", fileKey, err)
            continue
        }

        reviews, err := a.parser.ParseFile(data)
        if err != nil {
            log.Printf("failed to parse file %s: %v", fileKey, err)
            continue
        }

        validReviews := a.validator.FilterValid(reviews)

        if err := a.processor.ProcessReviews(validReviews); err != nil {
            log.Printf("failed to process reviews from file %s: %v", fileKey, err)
            continue
        }

        log.Printf("successfully processed file %s with %d reviews", fileKey, len(validReviews))
    }
    return nil
}
