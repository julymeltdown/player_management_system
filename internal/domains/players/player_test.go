package player

import (
	customErrors "player_management_system/internal/pkg/errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewPlayer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		p, err := NewPlayer("김도영", "야구", "기아", "https://example.com/image.jpg")

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
		_, err := NewPlayer("", "야구", "기아", "https://example.com/image.jpg")

		assert.Error(t, err)
		assert.IsType(t, &customErrors.Error{}, err)

		var customErr *customErrors.Error
		if assert.ErrorAs(t, err, &customErr) {
			assert.Equal(t, customErrors.InvalidArgumentError, customErr.Code)
			assert.Equal(t, "Invalid argument: name", customErr.Message)
		}
	})

	t.Run("missing sport", func(t *testing.T) {
		_, err := NewPlayer("김도영", "", "기아", "https://example.com/image.jpg")

		assert.Error(t, err)
		assert.IsType(t, &customErrors.Error{}, err)

		var customErr *customErrors.Error
		if assert.ErrorAs(t, err, &customErr) {
			assert.Equal(t, customErrors.InvalidArgumentError, customErr.Code)
			assert.Equal(t, "Invalid argument: sport", customErr.Message)
		}
	})

	t.Run("missing team", func(t *testing.T) {
		_, err := NewPlayer("김도영", "야구", "", "https://example.com/image.jpg")

		assert.Error(t, err)
		assert.IsType(t, &customErrors.Error{}, err)

		var customErr *customErrors.Error
		if assert.ErrorAs(t, err, &customErr) {
			assert.Equal(t, customErrors.InvalidArgumentError, customErr.Code)
			assert.Equal(t, "Invalid argument: team", customErr.Message)
		}
	})
}
