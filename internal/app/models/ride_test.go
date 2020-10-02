package models_test

import (
	"testing"
	"time"

	"github.com/exapsy/beat-exercise/internal/app/models"
)

func TestVelocity(t *testing.T) {
	t.Run("Velocity=0", func(t *testing.T) {
		startSegment := models.RideSegment{
			Point: models.Point{
				Latitude:  0,
				Longitude: 0,
			},
			Timestamp: time.Time{},
		}
		endSegment := models.RideSegment{
			Point: models.Point{
				Latitude:  0,
				Longitude: 0,
			},
			Timestamp: time.Time{},
		}
		ride := models.Ride{
			ID:    "",
			Start: startSegment,
			End:   endSegment,
		}

		velocity := ride.GetVelocity()

		if velocity != 0 {
			t.Error("Expected 0 velocity, got", velocity, "instead")
		}
	})
}
