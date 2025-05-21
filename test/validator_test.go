package tests

import (
    "testing"
    "review-service/internal/models"
    "review-service/internal/validator"
)

func TestValidateReview_Valid(t *testing.T) {
    review := models.Review{
        HotelId: 123,
        Platform: "Agoda",
        HotelName: "Test Hotel",
    }

    err := validator.ValidateReview(&review)
    if err != nil {
        t.Errorf("expected no validation error, got %v", err)
    }
}

func TestValidateReview_MissingRequiredFields(t *testing.T) {
    review := models.Review{} // missing required fields
    err := validator.ValidateReview(&review)
    if err == nil {
        t.Fatal("expected validation error, got nil")
    }
}
