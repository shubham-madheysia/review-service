package models

import "time"

type Review struct {
	ID          uint      `gorm:"primaryKey" json:"-"`                // DB primary key, ignored in JSON
	HotelID     int       `json:"hotelId" validate:"required"`        // Required field
	Platform    string    `json:"platform" validate:"required"`
	HotelName   string    `json:"hotelName" validate:"required"`
	ReviewDate  time.Time `json:"reviewDate" validate:"required"`
	Rating      float64   `json:"rating" validate:"required,gte=0,lte=10"`
	Reviewer    string    `json:"reviewer" validate:"required"`
	Country     string    `json:"country" validate:"required"`
	RoomType    string    `json:"roomType" validate:"required"`
	ReviewText  string    `json:"reviewText" validate:"required"`
	ReviewTitle string    `json:"reviewTitle" validate:"required"`
}
