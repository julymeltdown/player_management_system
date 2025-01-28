package player

import (
	"context"

	"github.com/google/uuid"
	"player_management_system/internal/domains/players"
	playerRepo "player_management_system/internal/repositories/player" // 수정된 부분
)

// PlayerService defines the interface for player-related operations.
type PlayerService interface {
	CreatePlayer(ctx context.Context, player *player.Player) error
	GetPlayerByID(ctx context.Context, id uuid.UUID) (*player.Player, error)
	UpdatePlayer(ctx context.Context, player *player.Player) error
	DeletePlayer(ctx context.Context, id uuid.UUID) error
	GetPlayers(ctx context.Context) ([]*player.Player, error)
}

type playerService struct {
	repo playerRepo.PlayerRepository // 수정된 부분 (인터페이스 타입 사용)
}

// NewPlayerService creates a new PlayerService instance.
func NewPlayerService(repo playerRepo.PlayerRepository) PlayerService {
	return &playerService{repo: repo}
}

// CreatePlayer creates a new player.
func (s *playerService) CreatePlayer(ctx context.Context, p *player.Player) error {
	return s.repo.CreatePlayer(ctx, p)
}

// GetPlayerByID retrieves a player by their ID.
func (s *playerService) GetPlayerByID(ctx context.Context, id uuid.UUID) (*player.Player, error) {
	return s.repo.GetPlayerByID(ctx, id)
}

// UpdatePlayer updates an existing player.
func (s *playerService) UpdatePlayer(ctx context.Context, p *player.Player) error {
	return s.repo.UpdatePlayer(ctx, p)
}

// DeletePlayer deletes a player by their ID.
func (s *playerService) DeletePlayer(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeletePlayer(ctx, id)
}

// GetPlayers retrieves all players.
func (s *playerService) GetPlayers(ctx context.Context) ([]*player.Player, error) {
	return s.repo.GetPlayers(ctx)
}
