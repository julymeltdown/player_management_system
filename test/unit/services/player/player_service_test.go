package player

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	playerDom "player_management_system/internal/domains/players"
	playerSvc "player_management_system/internal/services/player"
)

// MockPlayerRepository is a mock implementation of the PlayerRepository interface.
type MockPlayerRepository struct {
	mock.Mock
}

func (m *MockPlayerRepository) CreatePlayer(ctx context.Context, player *playerDom.Player) error {
	args := m.Called(ctx, player)
	return args.Error(0)
}

func (m *MockPlayerRepository) GetPlayerByID(ctx context.Context, id uuid.UUID) (*playerDom.Player, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*playerDom.Player), args.Error(1)
}

func (m *MockPlayerRepository) UpdatePlayer(ctx context.Context, player *playerDom.Player) error {
	args := m.Called(ctx, player)
	return args.Error(0)
}

func (m *MockPlayerRepository) DeletePlayer(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockPlayerRepository) GetPlayers(ctx context.Context) ([]*playerDom.Player, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*playerDom.Player), args.Error(1)
}

func TestCreatePlayer(t *testing.T) {
	mockRepo := new(MockPlayerRepository)
	service := playerSvc.NewPlayerService(mockRepo)

	p := &playerDom.Player{
		ID:              uuid.New(),
		Name:            "Test Player",
		Sport:           "Football",
		Team:            "Test Team",
		ProfileImageURL: "http://example.com/image.jpg",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	mockRepo.On("CreatePlayer", mock.Anything, p).Return(nil)

	err := service.CreatePlayer(context.Background(), p)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestGetPlayerByID(t *testing.T) {
	mockRepo := new(MockPlayerRepository)
	service := playerSvc.NewPlayerService(mockRepo)

	playerID := uuid.New()
	expectedPlayer := &playerDom.Player{
		ID:              playerID,
		Name:            "Test Player",
		Sport:           "Football",
		Team:            "Test Team",
		ProfileImageURL: "http://example.com/image.jpg",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	mockRepo.On("GetPlayerByID", mock.Anything, playerID).Return(expectedPlayer, nil)

	player, err := service.GetPlayerByID(context.Background(), playerID)
	assert.NoError(t, err)
	assert.Equal(t, expectedPlayer, player)

	mockRepo.AssertExpectations(t)
}

func TestUpdatePlayer(t *testing.T) {
	mockRepo := new(MockPlayerRepository)
	service := playerSvc.NewPlayerService(mockRepo)

	p := &playerDom.Player{
		ID:              uuid.New(),
		Name:            "Updated Player",
		Sport:           "Basketball",
		Team:            "Another Team",
		ProfileImageURL: "http://example.com/updated.jpg",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	mockRepo.On("UpdatePlayer", mock.Anything, p).Return(nil)

	err := service.UpdatePlayer(context.Background(), p)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeletePlayer(t *testing.T) {
	mockRepo := new(MockPlayerRepository)
	service := playerSvc.NewPlayerService(mockRepo)

	playerID := uuid.New()

	mockRepo.On("DeletePlayer", mock.Anything, playerID).Return(nil)

	err := service.DeletePlayer(context.Background(), playerID)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestGetPlayers(t *testing.T) {
	mockRepo := new(MockPlayerRepository)
	service := playerSvc.NewPlayerService(mockRepo)

	expectedPlayers := []*playerDom.Player{
		{
			ID:              uuid.New(),
			Name:            "Player 1",
			Sport:           "Football",
			Team:            "Team A",
			ProfileImageURL: "http://example.com/image1.jpg",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
		{
			ID:              uuid.New(),
			Name:            "Player 2",
			Sport:           "Basketball",
			Team:            "Team B",
			ProfileImageURL: "http://example.com/image2.jpg",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
	}

	mockRepo.On("GetPlayers", mock.Anything).Return(expectedPlayers, nil)

	players, err := service.GetPlayers(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedPlayers, players)

	mockRepo.AssertExpectations(t)
}
