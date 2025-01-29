package http

import (
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	playerDomain "player_management_system/internal/domains/players"
	customErrors "player_management_system/internal/pkg/errors"
	playerService "player_management_system/internal/services/player"
)

// PlayerHandler handles HTTP requests for player operations.
type PlayerHandler struct {
	playerService playerService.PlayerService
}

// NewPlayerHandler creates a new PlayerHandler.
func NewPlayerHandler(playerService playerService.PlayerService) *PlayerHandler {
	return &PlayerHandler{playerService: playerService}
}

// RegisterRoutes registers the player routes with the Echo router.
func (h *PlayerHandler) RegisterRoutes(e *echo.Echo) {
	e.POST("/players", h.CreatePlayer)
	e.GET("/players/:id", h.GetPlayer)
}

// CreatePlayerRequest represents the request body for creating a new player.
type CreatePlayerRequest struct {
	Name            string `json:"name"`
	Sport           string `json:"sport"`
	Team            string `json:"team"`
	ProfileImageURL string `json:"profile_image_url"`
}

// CreatePlayer handles the POST /players request.
func (h *PlayerHandler) CreatePlayer(c echo.Context) error {
	var req CreatePlayerRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, customErrors.NewInvalidArgumentError("Invalid request body"))
	}

	p, err := playerDomain.NewPlayer(req.Name, req.Sport, req.Team, req.ProfileImageURL)
	if err != nil {
		// Assuming NewPlayer can only return InvalidArgumentError
		return c.JSON(http.StatusBadRequest, err)
	}

	err = h.playerService.CreatePlayer(c.Request().Context(), p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, customErrors.NewInternalError("Failed to create player"))
	}

	return c.JSON(http.StatusCreated, p)
}

// GetPlayer handles the GET /players/:id request.
func (h *PlayerHandler) GetPlayer(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, customErrors.NewInvalidArgumentError("Invalid player ID"))
	}

	p, err := h.playerService.GetPlayerByID(c.Request().Context(), id)
	if err != nil {
		var notFoundErr *customErrors.Error
		if errors.As(err, &notFoundErr) && notFoundErr.Code == customErrors.NotFoundError {
			return c.JSON(http.StatusNotFound, notFoundErr)
		}
		return c.JSON(http.StatusInternalServerError, customErrors.NewInternalError("Failed to get player"))
	}

	return c.JSON(http.StatusOK, p)
}
