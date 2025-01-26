package player

import (
	"time"

	"github.com/google/uuid"
)

// PlayerDescription represents a detailed description of a player.
type PlayerDescription struct {
	ID        uuid.UUID `json:"id"`
	PlayerID  uuid.UUID `json:"player_id"` // Foreign key to Player
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
