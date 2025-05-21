-- Drop existing tables if needed (for dev/testing)
DROP TABLE IF EXISTS reviews;

-- Main reviews table
CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    hotel_id INTEGER NOT NULL,
    platform VARCHAR(50) NOT NULL,
    hotel_name TEXT,
    review_date TIMESTAMP NOT NULL,
    rating REAL,
    reviewer TEXT,
    country TEXT,
    room_type TEXT,
    review_text TEXT,
    review_title TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for performance
CREATE INDEX idx_reviews_hotel_id ON reviews(hotel_id);
CREATE INDEX idx_reviews_platform ON reviews(platform);
CREATE INDEX idx_reviews_review_date ON reviews(review_date);
