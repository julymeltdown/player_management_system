package player

import (
	"github.com/google/uuid"
	player "player_management_system/internal/domains/players"
	"testing"

	"github.com/stretchr/testify/assert"
	"player_management_system/internal/pkg/errors"
)

func TestNewPlayer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		p, err := player.NewPlayer("김도영", "야구", "기아", "https://example.com/image.jpg")

		assert.NoError(t, err)
		assert.NotNil(t, p)
		assert.NotEqual(t, uuid.UUID{}, p.ID)
		assert.Equal(t, "김도영", p.Name)
		assert.Equal(t, "야구", p.Sport)
		assert.Equal(t, "기아", p.Team)
		assert.Equal(t, "https://example.com/image.jpg", p.ProfileImageURL)
		assert.False(t, p.CreatedAt.IsZero())
		assert.False(t, p.UpdatedAt.IsZero())
	})

	t.Run("missing name", func(t *testing.T) {
		p, err := player.NewPlayer("", "야구", "기아", "https://example.com/image.jpg")
		assert.Error(t, err)
		assert.Nil(t, p)

		var invalidArgErr *errors.Error
		if assert.ErrorAs(t, err, &invalidArgErr) {
			assert.Equal(t, errors.InvalidArgumentError, invalidArgErr.Code)
			assert.Equal(t, "name is required", invalidArgErr.Message)
		}
	})

	t.Run("missing sport", func(t *testing.T) {
		p, err := player.NewPlayer("김도영", "", "기아", "https://example.com/image.jpg")
		assert.Error(t, err)
		assert.Nil(t, p)

		var invalidArgErr *errors.Error
		if assert.ErrorAs(t, err, &invalidArgErr) {
			assert.Equal(t, errors.InvalidArgumentError, invalidArgErr.Code)
			assert.Equal(t, "sport is required", invalidArgErr.Message)
		}
	})

	t.Run("missing team", func(t *testing.T) {
		p, err := player.NewPlayer("김도영", "야구", "", "https://example.com/image.jpg")
		assert.Error(t, err)
		assert.Nil(t, p)

		var invalidArgErr *errors.Error
		if assert.ErrorAs(t, err, &invalidArgErr) {
			assert.Equal(t, errors.InvalidArgumentError, invalidArgErr.Code)
			assert.Equal(t, "team is required", invalidArgErr.Message)
		}
	})
}
