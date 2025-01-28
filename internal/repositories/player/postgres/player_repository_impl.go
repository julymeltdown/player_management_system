package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"player_management_system/internal/domains/players"
	"player_management_system/internal/pkg/errors"
	playerRepo "player_management_system/internal/repositories/player"
)

type playerRepository struct {
	db *sqlx.DB
}

func NewPlayerRepository(db *sqlx.DB) playerRepo.PlayerRepository {
	return &playerRepository{db: db}
}

// CreatePlayer implements playerRepo.PlayerRepository.
func (r *playerRepository) CreatePlayer(ctx context.Context, p *player.Player) error {
	query := `
        INSERT INTO players (id, name, sport, team, profile_image_url, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `

	_, err := r.db.ExecContext(
		ctx,
		query,
		p.ID,
		p.Name,
		p.Sport,
		p.Team,
		p.ProfileImageURL,
		p.CreatedAt,
		p.UpdatedAt,
	)
	if err != nil {
		return errors.NewInternalError(err.Error())
	}

	return nil
}

// DeletePlayer implements playerRepo.PlayerRepository.
func (r *playerRepository) DeletePlayer(ctx context.Context, id uuid.UUID) error {
	query := `
        DELETE FROM players
        WHERE id = $1
    `

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return errors.NewInternalError(err.Error())
	}

	return nil
}

// GetPlayerByID implements playerRepo.PlayerRepository.
func (r *playerRepository) GetPlayerByID(ctx context.Context, id uuid.UUID) (*player.Player, error) {
	var p player.Player
	query := `
        SELECT id, name, sport, team, profile_image_url, created_at, updated_at
        FROM players
        WHERE id = $1
    `

	err := r.db.GetContext(ctx, &p, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("player not found")
		}
		return nil, errors.NewInternalError(err.Error())
	}

	return &p, nil
}

// GetPlayers implements playerRepo.PlayerRepository.
func (r *playerRepository) GetPlayers(ctx context.Context) ([]*player.Player, error) {
	var players []*player.Player
	query := `
        SELECT id, name, sport, team, profile_image_url, created_at, updated_at
        FROM players
    `

	err := r.db.SelectContext(ctx, &players, query)
	if err != nil {
		return nil, errors.NewInternalError(err.Error())
	}

	return players, nil
}

// UpdatePlayer implements playerRepo.PlayerRepository.
func (r *playerRepository) UpdatePlayer(ctx context.Context, p *player.Player) error {
	query := `
        UPDATE players
        SET name = $1, sport = $2, team = $3, profile_image_url = $4, updated_at = $5
        WHERE id = $6
    `

	_, err := r.db.ExecContext(
		ctx,
		query,
		p.Name,
		p.Sport,
		p.Team,
		p.ProfileImageURL,
		p.UpdatedAt,
		p.ID,
	)
	if err != nil {
		return errors.NewInternalError(err.Error())
	}

	return nil
}
