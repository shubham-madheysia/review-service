package tests

import (
    "testing"
    "strings"
    "review-service/internal/parser"
)

func TestParseLine_ValidJSON(t *testing.T) {
    line := `{"hotelId":10984,"platform":"Agoda","hotelName":"Oscar Saigon Hotel","comment":{"hotelReviewId":948353737,"rating":6.4}}`
    
    review, err := parser.ParseLine(strings.NewReader(line))
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if review.HotelId != 10984 {
        t.Errorf("expected hotelId 10984, got %d", review.HotelId)
    }
}

func TestParseLine_InvalidJSON(t *testing.T) {
    line := `{invalid json}`
    _, err := parser.ParseLine(strings.NewReader(line))
    if err == nil {
        t.Fatal("expected error for invalid JSON, got nil")
    }
}
