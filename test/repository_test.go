package tests

import (
    "testing"
    "review-service/internal/db"
    "review-service/internal/models"
    "os"
)

func setupTestDB(t *testing.T) *db.Repository {
    dsn := os.Getenv("TEST_DB_DSN")
    if dsn == "" {
        t.Skip("TEST_DB_DSN not set, skipping DB tests")
    }
    repo, err := db.NewRepository(dsn)
    if err != nil {
        t.Fatalf("failed to connect to test db: %v", err)
    }
    return repo
}

func TestSaveReview(t *testing.T) {
    repo := setupTestDB(t)
    review := models.Review{
        HotelId: 999,
        Platform: "TestPlatform",
        HotelName: "Test Hotel",
        Comment: models.Comment{
            HotelReviewId: 1,
            Rating: 8.5,
        },
    }

    err := repo.SaveReview(&review)
    if err != nil {
        t.Fatalf("failed to save review: %v", err)
    }
}
