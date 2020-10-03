package csvutils_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/exapsy/beat-exercise/internal/app/drivers/csvutils"
	"github.com/exapsy/beat-exercise/internal/app/models"
)

func TestGetRideOutputString(t *testing.T) {
	t.Run("2_rides", func(t *testing.T) {
		segmentA := models.RideSegment{
			Point: models.Point{
				Latitude:  0,
				Longitude: 0,
			},
		}
		segmentB := models.RideSegment{
			Timestamp: time.Time{}.Add(time.Hour * 1),
			Point: models.Point{
				Latitude:  0.5,
				Longitude: 0.5,
			},
		}
		ride := models.Ride{
			ID: "0",
			Segments: []models.RideSegment{
				segmentA,
				segmentB,
			},
		}
		rideOutput := csvutils.GetRideOutputString(ride)
		fare := strconv.FormatFloat(float64(ride.EstimateFare()), 'g', -1, 32)
		expectedOutput := ride.ID + ", " + fare + "\r\n"
		if rideOutput != expectedOutput {
			t.Error("Expected output `", expectedOutput, "` but instead got `", rideOutput, "`")
		}
	})
}
