package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	playerDom "player_management_system/internal/domains/players"
	playerRepoImpl "player_management_system/internal/repositories/player/postgres"
)

func TestCreatePlayer(t *testing.T) {
	// Create sqlmock database connection and a mock
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Create sqlx db instance from the mocked connection
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Create the repository
	repo := playerRepoImpl.NewPlayerRepository(sqlxDB)

	// Create a new player
	p := &playerDom.Player{
		ID:              uuid.New(),
		Name:            "Test Player",
		Sport:           "Football",
		Team:            "Test Team",
		ProfileImageURL: "http://example.com/image.jpg",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	// Expect the query to be executed with the correct parameters
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO players (id, name, sport, team, profile_image_url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`)).
		WithArgs(p.ID, p.Name, p.Sport, p.Team, p.ProfileImageURL, p.CreatedAt, p.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1)) // 1 row affected

	// Test CreatePlayer
	err = repo.CreatePlayer(context.Background(), p)
	assert.NoError(t, err)

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetPlayerByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := playerRepoImpl.NewPlayerRepository(sqlxDB)

	playerID := uuid.New()
	rows := sqlmock.NewRows([]string{"id", "name", "sport", "team", "profile_image_url", "created_at", "updated_at"}). // profile_image_url 추가
																AddRow(playerID, "Test Player", "Football", "Test Team", "http://example.com/image.jpg", time.Now(), time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, sport, team, profile_image_url, created_at, updated_at FROM players WHERE id = $1`)).
		WithArgs(playerID).
		WillReturnRows(rows)

	player, err := repo.GetPlayerByID(context.Background(), playerID)
	assert.NoError(t, err)
	assert.NotNil(t, player)
	assert.Equal(t, "Test Player", player.Name)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUpdatePlayer(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := playerRepoImpl.NewPlayerRepository(sqlxDB)

	p := &playerDom.Player{
		ID:              uuid.New(),
		Name:            "Test Player",
		Sport:           "Football",
		Team:            "Test Team",
		ProfileImageURL: "http://example.com/image.jpg",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE players SET name = $1, sport = $2, team = $3, profile_image_url = $4, updated_at = $5 WHERE id = $6`)).
		WithArgs(p.Name, p.Sport, p.Team, p.ProfileImageURL, p.UpdatedAt, p.ID).
		WillReturnResult(sqlmock.NewResult(0, 1)) // 1 row affected

	err = repo.UpdatePlayer(context.Background(), p)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestDeletePlayer(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := playerRepoImpl.NewPlayerRepository(sqlxDB)

	playerID := uuid.New()

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM players WHERE id = $1`)).
		WithArgs(playerID).
		WillReturnResult(sqlmock.NewResult(0, 1)) // 1 row affected

	err = repo.DeletePlayer(context.Background(), playerID)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetPlayers(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := playerRepoImpl.NewPlayerRepository(sqlxDB)

	rows := sqlmock.NewRows([]string{"id", "name", "sport", "team", "profile_image_url", "created_at", "updated_at"}).
		AddRow(uuid.New(), "Test Player 1", "Football", "Test Team", "http://example.com/image1.jpg", time.Now(), time.Now()).
		AddRow(uuid.New(), "Test Player 2", "Basketball", "Test Team", "http://example.com/image2.jpg", time.Now(), time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, sport, team, profile_image_url, created_at, updated_at FROM players`)).
		WillReturnRows(rows)

	players, err := repo.GetPlayers(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, players)
	assert.Equal(t, 2, len(players))

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
