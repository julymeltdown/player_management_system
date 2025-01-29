package http

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	playerDomain "player_management_system/internal/domains/players"
	customErrors "player_management_system/internal/pkg/errors"
	_ "player_management_system/internal/services/player"
)

// MockPlayerService is a mock implementation of the PlayerService interface for testing.
type MockPlayerService struct {
	mock.Mock
}

func (m *MockPlayerService) CreatePlayer(ctx context.Context, player *playerDomain.Player) error {
	args := m.Called(ctx, player)
	return args.Error(0)
}

func (m *MockPlayerService) GetPlayerByID(ctx context.Context, id uuid.UUID) (*playerDomain.Player, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*playerDomain.Player), args.Error(1)
}

func (m *MockPlayerService) UpdatePlayer(ctx context.Context, player *playerDomain.Player) error {
	args := m.Called(ctx, player)
	return args.Error(0)
}

func (m *MockPlayerService) DeletePlayer(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockPlayerService) GetPlayers(ctx context.Context) ([]*playerDomain.Player, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*playerDomain.Player), args.Error(1)
}

func (m *MockPlayerService) GetPlayersWithPagination(ctx context.Context, page, pageSize int) ([]*playerDomain.Player, error) {
	args := m.Called(ctx, page, pageSize)
	return args.Get(0).([]*playerDomain.Player), args.Error(1)
}

func TestCreatePlayer_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/players", strings.NewReader(`{"name":"Test Player","sport":"Football","team":"Test Team","profile_image_url":"http://example.com"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockPlayerService)
	mockService.On("CreatePlayer", mock.Anything, mock.AnythingOfType("*player.Player")).Return(nil)

	handler := NewPlayerHandler(mockService)

	// Assertions
	if assert.NoError(t, handler.CreatePlayer(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestCreatePlayer_InvalidRequestBody(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/players", strings.NewReader(`{invalid json}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockPlayerService)
	handler := NewPlayerHandler(mockService)

	// 실행
	err := handler.CreatePlayer(c)

	// 검증
	assert.Error(t, err)

	var httpErr *echo.HTTPError
	if errors.As(err, &httpErr) {
		assert.Equal(t, http.StatusBadRequest, httpErr.Code)
		assert.Equal(t, "Invalid request body", httpErr.Message) // 변경된 메시지 확인
	} else {
		assert.Fail(t, "Expected *echo.HTTPError")
	}
}

func TestCreatePlayer_NewPlayerError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/players", strings.NewReader(`{"name":"","sport":"Football","team":"Test Team","profile_image_url":"http://example.com"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockPlayerService)
	handler := NewPlayerHandler(mockService)

	// 실행
	err := handler.CreatePlayer(c)

	// 검증
	assert.Error(t, err)

	var httpErr *echo.HTTPError
	if errors.As(err, &httpErr) {
		assert.Equal(t, http.StatusBadRequest, httpErr.Code)
	} else {
		assert.Fail(t, "Expected *echo.HTTPError")
	}
}

func TestCreatePlayer_CreatePlayerError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/players", strings.NewReader(`{"name":"Test Player","sport":"Football","team":"Test Team","profile_image_url":"http://example.com"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockPlayerService)
	mockService.On("CreatePlayer", mock.Anything, mock.AnythingOfType("*player.Player")).Return(customErrors.NewError(customErrors.DatabaseError, "database error"))
	handler := NewPlayerHandler(mockService)

	// 실행
	err := handler.CreatePlayer(c)

	// 검증
	assert.Error(t, err)

	var httpErr *echo.HTTPError
	if errors.As(err, &httpErr) {
		assert.Equal(t, http.StatusInternalServerError, httpErr.Code)
	} else {
		assert.Fail(t, "Expected *echo.HTTPError")
	}
}

func TestGetPlayer_Success(t *testing.T) {
	playerId := uuid.New()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/players/"+playerId.String(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/players/:id")
	c.SetParamNames("id")
	c.SetParamValues(playerId.String())

	mockService := new(MockPlayerService)
	expectedPlayer := &playerDomain.Player{
		ID:   playerId,
		Name: "Test Player",
	}
	mockService.On("GetPlayerByID", mock.Anything, playerId).Return(expectedPlayer, nil)

	handler := NewPlayerHandler(mockService)

	// Assertions
	if assert.NoError(t, handler.GetPlayer(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetPlayer_InvalidID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/players/invalid-id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/players/:id")
	c.SetParamNames("id")
	c.SetParamValues("invalid-id")

	mockService := new(MockPlayerService)
	handler := NewPlayerHandler(mockService)

	// 실행
	err := handler.GetPlayer(c)

	// 검증
	assert.Error(t, err)

	var httpErr *echo.HTTPError
	if errors.As(err, &httpErr) {
		assert.Equal(t, http.StatusBadRequest, httpErr.Code)
	} else {
		assert.Fail(t, "Expected *echo.HTTPError")
	}
}

func TestGetPlayer_PlayerNotFound(t *testing.T) {
	playerId := uuid.New()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/players/"+playerId.String(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/players/:id")
	c.SetParamNames("id")
	c.SetParamValues(playerId.String())

	mockService := new(MockPlayerService)
	mockService.On("GetPlayerByID", mock.Anything, playerId).Return((*playerDomain.Player)(nil), customErrors.NewError(customErrors.NotFoundError, "player not found"))
	handler := NewPlayerHandler(mockService)

	// 실행
	err := handler.GetPlayer(c)

	// 검증
	assert.Error(t, err)

	var httpErr *echo.HTTPError
	if errors.As(err, &httpErr) {
		assert.Equal(t, http.StatusNotFound, httpErr.Code)
	} else {
		assert.Fail(t, "Expected *echo.HTTPError")
	}
}

func TestGetPlayer_InternalError(t *testing.T) {
	playerId := uuid.New()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/players/"+playerId.String(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/players/:id")
	c.SetParamNames("id")
	c.SetParamValues(playerId.String())

	mockService := new(MockPlayerService)
	mockService.On("GetPlayerByID", mock.Anything, playerId).Return((*playerDomain.Player)(nil), customErrors.NewError(customErrors.InternalError, "internal error"))
	handler := NewPlayerHandler(mockService)

	// 실행
	err := handler.GetPlayer(c)

	// 검증
	assert.Error(t, err)

	var httpErr *echo.HTTPError
	if errors.As(err, &httpErr) {
		assert.Equal(t, http.StatusInternalServerError, httpErr.Code)
	} else {
		assert.Fail(t, "Expected *echo.HTTPError")
	}
}

func TestGetPlayers_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/players?page=1&size=10", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockPlayerService)
	expectedPlayers := []*playerDomain.Player{
		{
			ID:              uuid.New(),
			Name:            "Test Player 1",
			Sport:           "Football",
			Team:            "Test Team 1",
			ProfileImageURL: "http://example.com/image1.jpg",
		},
		{
			ID:              uuid.New(),
			Name:            "Test Player 2",
			Sport:           "Basketball",
			Team:            "Test Team 2",
			ProfileImageURL: "http://example.com/image2.jpg",
		},
	}
	mockService.On("GetPlayersWithPagination", mock.Anything, 1, 10).Return(expectedPlayers, nil)

	handler := NewPlayerHandler(mockService)

	// Assertions
	if assert.NoError(t, handler.GetPlayers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetPlayers_InvalidPage(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/players?page=invalid&size=10", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockPlayerService)
	// Page가 유효하지 않은 경우, 기본값으로 page=1, size=10을 사용하도록 설정
	mockService.On("GetPlayersWithPagination", mock.Anything, 1, 10).Return([]*playerDomain.Player{}, nil)
	handler := NewPlayerHandler(mockService)

	// Assertions
	assert.NoError(t, handler.GetPlayers(c))
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetPlayers_InvalidSize(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/players?page=1&size=invalid", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockPlayerService)
	// Size가 유효하지 않은 경우, 기본값으로 page=1, size=10을 사용하도록 설정
	mockService.On("GetPlayersWithPagination", mock.Anything, 1, 10).Return([]*playerDomain.Player{}, nil)
	handler := NewPlayerHandler(mockService)

	// Assertions
	assert.NoError(t, handler.GetPlayers(c))
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetPlayers_ServiceError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/players?page=1&size=10", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockPlayerService)
	mockService.On("GetPlayersWithPagination", mock.Anything, 1, 10).Return([]*playerDomain.Player{}, customErrors.NewError(customErrors.DatabaseError, "database error"))
	handler := NewPlayerHandler(mockService)

	// 실행
	err := handler.GetPlayers(c)

	// 검증
	assert.Error(t, err)

	var httpErr *echo.HTTPError
	if errors.As(err, &httpErr) {
		assert.Equal(t, http.StatusInternalServerError, httpErr.Code)
	} else {
		assert.Fail(t, "Expected *echo.HTTPError")
	}
}
