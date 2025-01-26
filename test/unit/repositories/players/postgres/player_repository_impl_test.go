package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"player-ms/internal/domains/player"
	"player-ms/internal/platform/postgres"
)

func TestCreatePlayer(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := postgres.NewPostgresDB(db) // sqlxDB는 *sqlx.DB 타입
	repo := NewPlayerRepository(sqlxDB)

	p := &player.Player{
		ID:              uuid.New(),
		Name:            "Test Player",
		Sport:           "Football",
		Team:            "Test Team",
		ProfileImageURL: "http://example.com/image.jpg",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	mock.ExpectExec("INSERT INTO players").
		WithArgs(p.ID, p.Name, p.Sport, p.Team, p.ProfileImageURL, p.CreatedAt, p.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreatePlayer(context.Background(), p)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
