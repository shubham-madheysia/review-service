package db

import (
	"context"
	"log"
	"reviewservice/internal/models"
	"time"

	"github.com/cenkalti/backoff/v4"
	"gorm.io/gorm"
)

// ReviewRepository defines methods to work with review data
type ReviewRepository interface {
	SaveReview(ctx context.Context, review *models.Review) error
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db: db}
}

// SaveReview inserts a review with retry logic on transient errors
func (r *reviewRepository) SaveReview(ctx context.Context, review *models.Review) error {
	operation := func() error {
		result := r.db.WithContext(ctx).Create(review)
		if result.Error != nil {
			log.Printf("DB insert failed: %v", result.Error)
			return result.Error
		}
		return nil
	}

	// Retry on transient DB errors with exponential backoff
	backOffConfig := backoff.NewExponentialBackOff()
	backOffConfig.MaxElapsedTime = 10 * time.Second

	err := backoff.Retry(operation, backOffConfig)
	if err != nil {
		return err
	}
	return nil
}
