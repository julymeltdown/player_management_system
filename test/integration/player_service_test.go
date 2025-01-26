package integration

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	testcontainers_pg "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	playerDom "player_management_system/internal/domains/players"
	playerRepo "player_management_system/internal/repositories/player/postgres"
	playerSvc "player_management_system/internal/services/player"
)

func TestCreateAndGetPlayer(t *testing.T) {
	ctx := context.Background()

	// Create a new PostgreSQL container
	pgContainer, err := testcontainers_pg.RunContainer(
		ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		testcontainers_pg.WithInitScripts("testdata/init.sql"), // init.sql 파일의 경로 확인 필요
		testcontainers_pg.WithDatabase("testdb"),
		testcontainers_pg.WithUsername("testuser"),
		testcontainers_pg.WithPassword("testpassword"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer pgContainer.Terminate(ctx)

	// Get the connection string
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	assert.NoError(t, err)

	// Connect to the database using GORM
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	assert.NoError(t, err)

	// Migrate the schema
	err = db.AutoMigrate(&playerDom.Player{}, &playerDom.PlayerDescription{}, &playerDom.Media{})
	assert.NoError(t, err)

	// Create the repository and service
	repo := playerRepo.NewPlayerRepository(db)
	service := playerSvc.NewPlayerService(repo)

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
	err = service.CreatePlayer(ctx, p)
	assert.NoError(t, err)

	// Get the player by ID
	retrievedPlayer, err := service.GetPlayerByID(ctx, p.ID)
	assert.NoError(t, err)
	assert.NotNil(t, retrievedPlayer)
	assert.Equal(t, p.Name, retrievedPlayer.Name)
}
