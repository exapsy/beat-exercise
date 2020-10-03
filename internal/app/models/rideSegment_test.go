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
			Timestamp: time.Time{}.Add(time.Second * 1),
		}

		velocity := startSegment.VelocityFrom(endSegment)

		if velocity != 0 {
			t.Error("Expected 0 velocity, got", velocity, "instead")
		}
	})
	t.Run("Velocity=Positive", func(t *testing.T) {
		startTime, err := time.Parse(
			time.RFC3339,
			"2017-05-05T12:35:00+00:00",
		)
		if err != nil {
			t.Fatal(err)
		}
		endTime, err := time.Parse(
			time.RFC3339,
			"2017-05-05T13:10:00+00:00",
		)
		if err != nil {
			t.Fatal(err)
		}
		startSegment := models.RideSegment{
			Point: models.Point{
				Latitude:  38.920602,
				Longitude: 77.222329,
			},
			Timestamp: startTime,
		}
		endSegment := models.RideSegment{
			Point: models.Point{
				Latitude:  38.889011,
				Longitude: 77.050061,
			},
			Timestamp: endTime,
		}

		velocity := startSegment.VelocityFrom(endSegment)

		if velocity != 23.33051501315813 {
			t.Error("Expected 23.33051501315813 velocity, got", velocity, "instead")
		}
	})
}

func TestEquals(t *testing.T) {
	t.Run("equal", func(t *testing.T) {
		segmentA := models.RideSegment{
			Point: models.Point{
				Latitude:  38.920602,
				Longitude: 77.222329,
			},
			Timestamp: time.Time{}.Add(time.Hour),
		}
		segmentB := models.RideSegment{
			Point: models.Point{
				Latitude:  38.920602,
				Longitude: 77.222329,
			},
			Timestamp: time.Time{}.Add(time.Hour),
		}
		equals := segmentA.Equals(segmentB)
		if !equals {
			t.Error("Expected segments to be equal")
		}
	})
	t.Run("not_equal_longitude", func(t *testing.T) {
		segmentA := models.RideSegment{
			Point: models.Point{
				Latitude:  38.920602,
				Longitude: 77.222329,
			},
			Timestamp: time.Time{}.Add(time.Hour),
		}
		segmentB := models.RideSegment{
			Point: models.Point{
				Latitude:  38.920602,
				Longitude: 78.222329,
			},
			Timestamp: time.Time{}.Add(time.Hour),
		}
		equals := segmentA.Equals(segmentB)
		if equals {
			t.Error("Expected segments not be equal")
		}
	})
	t.Run("not_equal_latitude", func(t *testing.T) {
		segmentA := models.RideSegment{
			Point: models.Point{
				Latitude:  38.920602,
				Longitude: 77.222329,
			},
			Timestamp: time.Time{}.Add(time.Hour),
		}
		segmentB := models.RideSegment{
			Point: models.Point{
				Latitude:  37.920602,
				Longitude: 77.222329,
			},
			Timestamp: time.Time{}.Add(time.Hour),
		}
		equals := segmentA.Equals(segmentB)
		if equals {
			t.Error("Expected segments not be equal")
		}
	})
	t.Run("not_equal_timestamp", func(t *testing.T) {
		segmentA := models.RideSegment{
			Point: models.Point{
				Latitude:  38.920602,
				Longitude: 77.222329,
			},
			Timestamp: time.Time{}.Add(time.Hour),
		}
		segmentB := models.RideSegment{
			Point: models.Point{
				Latitude:  38.920602,
				Longitude: 77.222329,
			},
			Timestamp: time.Time{}.Add(time.Second),
		}
		equals := segmentA.Equals(segmentB)
		if equals {
			t.Error("Expected segments not be equal")
		}
	})
}
