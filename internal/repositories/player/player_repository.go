package player

import (
	"context"

	"github.com/google/uuid"
	"player_management_system/internal/domains/players"
)

// PlayerRepository defines the interface for player repository operations.
type PlayerRepository interface {
	CreatePlayer(ctx context.Context, player *player.Player) error
	GetPlayerByID(ctx context.Context, id uuid.UUID) (*player.Player, error)
	UpdatePlayer(ctx context.Context, player *player.Player) error
	DeletePlayer(ctx context.Context, id uuid.UUID) error
	GetPlayers(ctx context.Context) ([]*player.Player, error)
}
