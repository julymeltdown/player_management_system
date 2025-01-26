package player

import (
	"time"

	"github.com/google/uuid"
)

// Media represents a media entity related to a player.
type Media struct {
	ID           uuid.UUID `json:"id"`
	PlayerID     uuid.UUID `json:"player_id"` // Foreign key to Player
	Source       string    `json:"source"`
	URL          string    `json:"url"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	PublishedAt  time.Time `json:"published_at"`
	ThumbnailURL string    `json:"thumbnail_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
