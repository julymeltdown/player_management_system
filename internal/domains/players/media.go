package player

import (
	"time"

	"github.com/google/uuid"
)

// Media represents a media entity related to a player.
type Media struct {
	ID           uuid.UUID `json:"id" db:"id"`
	PlayerID     uuid.UUID `json:"player_id" db:"player_id"`
	Source       string    `json:"source" db:"source"`
	URL          string    `json:"url" db:"url"`
	Title        string    `json:"title" db:"title"`
	Content      string    `json:"content" db:"content"`
	PublishedAt  time.Time `json:"published_at" db:"published_at"`
	ThumbnailURL string    `json:"thumbnail_url" db:"thumbnail_url"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
