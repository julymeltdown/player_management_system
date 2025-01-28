package player

import (
	"time"

	"github.com/google/uuid"
)

// PlayerDescription represents a detailed description of a player.
type PlayerDescription struct {
	ID        uuid.UUID `json:"id" db:"id"`
	PlayerID  uuid.UUID `json:"player_id" db:"player_id"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
