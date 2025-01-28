package player

import (
	"time"

	"github.com/google/uuid"
	"player_management_system/internal/pkg/errors"
)

// Player represents a player entity.
type Player struct {
	ID              uuid.UUID `json:"id" db:"id"`
	Name            string    `json:"name" db:"name"`
	Sport           string    `json:"sport" db:"sport"`
	Team            string    `json:"team" db:"team"`
	ProfileImageURL string    `json:"profile_image_url" db:"profile_image_url"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// NewPlayer creates a new Player entity.
func NewPlayer(name, sport, team, profileImageURL string) (*Player, error) {
	if name == "" {
		return nil, errors.NewInvalidArgumentError("name is required")
	}
	if sport == "" {
		return nil, errors.NewInvalidArgumentError("sport is required")
	}
	if team == "" {
		return nil, errors.NewInvalidArgumentError("team is required")
	}

	return &Player{
		ID:              uuid.New(),
		Name:            name,
		Sport:           sport,
		Team:            team,
		ProfileImageURL: profileImageURL,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}, nil
}
