package models_test

import (
	"testing"
	"time"

	"github.com/exapsy/beat-exercise/internal/app/models"
)

func TestMakeRide(t *testing.T) {
	t.Run("no_segments", func(t *testing.T) {
		ride := models.MakeRide("0", []models.RideSegment{})
		if len(ride.Segments) != 0 {
			t.Error("Expected segments with length of 0 but got", len(ride.Segments))
		}
	})
	t.Run("skip_big_velocity_segment", func(t *testing.T) {
		// 2 segments with in between > 100kmh velocity
		segmentA := models.RideSegment{
			Timestamp: time.Time{},
			Point: models.Point{
				Latitude:  0,
				Longitude: 0,
			},
		}
		// SHOULD BE FILTERED OUT - velocity > 100kmh velocity from segmentA
		segmentB := models.RideSegment{
			Timestamp: time.Time{}.Add(time.Second),
			Point: models.Point{
				Latitude:  1,
				Longitude: 1,
			},
		}
		// < 100kmh velocity from segmentA
		segmentC := models.RideSegment{
			Timestamp: time.Time{}.Add(time.Second * 2),
			Point: models.Point{
				Latitude:  0.00005,
				Longitude: 0.00005,
			},
		}
		// SHOULD BE FILTERED OUT - velocity > 100kmh velocity from segmentC
		segmentD := models.RideSegment{
			Timestamp: time.Time{}.Add(time.Second * 3),
			Point: models.Point{
				Latitude:  10,
				Longitude: 12,
			},
		}
		// < 100kmh velocity from segmentC
		segmentE := models.RideSegment{
			Timestamp: time.Time{}.Add(time.Second * 8),
			Point: models.Point{
				Latitude:  0.0006,
				Longitude: 0.0006,
			},
		}
		v := segmentE.VelocityFrom(segmentC)
		t.Log(v)
		ride := models.MakeRide("0", []models.RideSegment{
			segmentA,
			segmentB,
			segmentC,
			segmentD,
			segmentE,
		})
		if ride.ID != "0" {
			t.Error("Expected ID '0' but instead got", ride.ID)
		}
		if len(ride.Segments) != 3 {
			t.Error("Expected 3 segments but instead got", len(ride.Segments))
		}
		for i := range ride.Segments[:1] {
			if ride.Segments[i+1].VelocityFrom(ride.Segments[i+2]) > 100 {
				t.Error("Expected velocity <= 100kmh but got", ride.Segments[i+1].VelocityFrom(ride.Segments[i+2]))
			}
		}
	})
	t.Run("segments_length_1", func(t *testing.T) {
		segmentA := models.RideSegment{
			Timestamp: time.Time{},
			Point: models.Point{
				Latitude:  0,
				Longitude: 0,
			},
		}
		ride := models.MakeRide("0", []models.RideSegment{
			segmentA,
		})
		if ride.ID != "0" {
			t.Error("Expected ID '0' but instead got", ride.ID)
		}
		if len(ride.Segments) != 1 {
			t.Error("Expected 1 segments but instead got", len(ride.Segments))
		}
	})
}

func TestEstimateFare(t *testing.T) {
	t.Run("minimum_fare_no_segments", func(t *testing.T) {
		ride := models.MakeRide("0", []models.RideSegment{})
		fare := ride.EstimateFare()
		if fare != models.FareMinimum {
			t.Error("Expected minimum fare of", models.FareMinimum, "but got instead", fare)
		}
	})
	t.Run("idle_3_hours", func(t *testing.T) {
		segmentA := models.RideSegment{
			Timestamp: time.Time{},
			Point: models.Point{
				Latitude:  0,
				Longitude: 0,
			},
		}
		segmentB := models.RideSegment{
			Timestamp: time.Time{}.Add(time.Hour * 3),
			Point: models.Point{
				Latitude:  0.1,
				Longitude: 0.2,
			},
		}
		ride := models.MakeRide("0", []models.RideSegment{
			segmentA,
			segmentB,
		})
		fare := ride.EstimateFare()
		expectedFare := models.FareFlag + models.FareIdlePerHour*3
		if fare != expectedFare {
			t.Error("Expected idle fare of", expectedFare, "but got", fare)
		}
	})
	t.Run("3_km_midnight", func(t *testing.T) {
		segmentA := models.RideSegment{
			Timestamp: time.Time{},
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
		segmentC := models.RideSegment{
			Timestamp: time.Time{}.Add(time.Hour * 2),
			Point: models.Point{
				Latitude:  1,
				Longitude: 1,
			},
		}
		ride := models.MakeRide("0", []models.RideSegment{
			segmentA,
			segmentB,
			segmentC,
		})
		fare := ride.EstimateFare()
		expectedFare := models.FareFlag +
			models.FareMidnightPerKm*
				(float32(segmentB.DistanceFrom(segmentA))) +
			models.FareMidnightPerKm*
				(float32(segmentC.DistanceFrom(segmentB)))
		if fare != expectedFare {
			t.Error("Expected fare of", expectedFare, "but got", fare)
		}
	})
	t.Run("3_km_midday", func(t *testing.T) {
		segmentA := models.RideSegment{
			Timestamp: time.Time{}.Add(time.Hour * 6),
			Point: models.Point{
				Latitude:  0,
				Longitude: 0,
			},
		}
		segmentB := models.RideSegment{
			Timestamp: time.Time{}.Add(time.Hour * 7),
			Point: models.Point{
				Latitude:  0.5,
				Longitude: 0.5,
			},
		}
		segmentC := models.RideSegment{
			Timestamp: time.Time{}.Add(time.Hour * 8),
			Point: models.Point{
				Latitude:  1,
				Longitude: 1,
			},
		}
		ride := models.MakeRide("0", []models.RideSegment{
			segmentA,
			segmentB,
			segmentC,
		})
		fare := ride.EstimateFare()
		expectedFare := models.FareFlag +
			models.FareMiddayPerKm*
				(float32(segmentB.DistanceFrom(segmentA))) +
			models.FareMiddayPerKm*
				(float32(segmentC.DistanceFrom(segmentB)))
		if fare != expectedFare {
			t.Error("Expected fare of", expectedFare, "but got", fare)
		}
	})
	t.Run("minimum_fare", func(t *testing.T) {
		segmentA := models.RideSegment{
			Timestamp: time.Time{},
			Point: models.Point{
				Latitude:  0,
				Longitude: 0,
			},
		}
		segmentB := models.RideSegment{
			Timestamp: time.Time{}.Add(time.Minute * 1),
			Point: models.Point{
				Latitude:  0.001,
				Longitude: 0.001,
			},
		}
		ride := models.MakeRide("0", []models.RideSegment{
			segmentA,
			segmentB,
		})
		fare := ride.EstimateFare()
		if fare != models.FareMinimum {
			t.Error("Expected minimum fare of", models.FareMinimum, "but got instead", fare)
		}
	})
}
