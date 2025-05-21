package parser

import (
	"encoding/json"
	"fmt"
	"reviewservice/internal/models"
	"strings"
	"time"
)

// rawReview mirrors the JSON structure for unmarshalling.
type rawReview struct {
	HotelID   int    `json:"hotelId"`
	Platform  string `json:"platform"`
	HotelName string `json:"hotelName"`
	Comment   struct {
		ReviewDate     string  `json:"reviewDate"`
		Rating         float64 `json:"rating"`
		ReviewTitle    string  `json:"reviewTitle"`
		ReviewComments string  `json:"reviewComments"`
		ReviewerInfo   struct {
			DisplayMemberName string `json:"displayMemberName"`
			CountryName       string `json:"countryName"`
			RoomTypeName      string `json:"roomTypeName"`
		} `json:"reviewerInfo"`
	} `json:"comment"`
}

// ParseLine parses a single JSON line into a Review struct with detailed error info.
func ParseLine(line string) (*models.Review, error) {
	var raw rawReview
	if err := json.Unmarshal([]byte(line), &raw); err != nil {
		return nil, fmt.Errorf("json unmarshal error: %w", err)
	}

	// Validate and parse the date field.
	if strings.TrimSpace(raw.Comment.ReviewDate) == "" {
		return nil, fmt.Errorf("missing reviewDate field")
	}
	date, err := time.Parse(time.RFC3339, raw.Comment.ReviewDate)
	if err != nil {
		return nil, fmt.Errorf("invalid reviewDate format: %w", err)
	}

	// Additional validation can be added here (e.g. rating range, required fields).

	return &models.Review{
		HotelID:     raw.HotelID,
		Platform:    raw.Platform,
		HotelName:   raw.HotelName,
		ReviewDate:  date,
		Rating:      raw.Comment.Rating,
		ReviewTitle: raw.Comment.ReviewTitle,
		ReviewText:  raw.Comment.ReviewComments,
		Reviewer:    raw.Comment.ReviewerInfo.DisplayMemberName,
		Country:     raw.Comment.ReviewerInfo.CountryName,
		RoomType:    raw.Comment.ReviewerInfo.RoomTypeName,
	}, nil
}

// ParseLines parses multiple JSON lines and returns slices of parsed reviews and errors.
// Each error includes line number and context.
func ParseLines(lines []string) ([]*models.Review, []error) {
	var reviews []*models.Review
	var errs []error

	for i, line := range lines {
		// Skip empty lines early
		if strings.TrimSpace(line) == "" {
			continue
		}
		r, err := ParseLine(line)
		if err != nil {
			errs = append(errs, fmt.Errorf("line %d: %w", i+1, err))
			continue
		}
		reviews = append(reviews, r)
	}

	return reviews, errs
}
