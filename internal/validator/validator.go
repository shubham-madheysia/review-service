package validator

import (
	"errors"
	"reviewservice/internal/models"
)

func Validate(review *models.Review) error {
	if review.HotelID == 0 {
		return errors.New("missing hotel ID")
	}
	if review.Platform == "" {
		return errors.New("missing platform")
	}
	if review.ReviewDate.IsZero() {
		return errors.New("missing or invalid review date")
	}
	if review.Reviewer == "" {
		return errors.New("missing reviewer")
	}
	if review.ReviewText == "" {
		return errors.New("empty review text")
	}
	return nil
}
