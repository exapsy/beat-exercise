package models

import (
	"math"
	"time"

	"github.com/exapsy/beat-exercise/pkg/formulas"
)

// RideSegment describes either the start or the end of a ride
type RideSegment struct {
	Point     Point
	Timestamp time.Time
}

// GetVelocity returns the km/h calculated
// between the two segments of a ride
// using the Haversine distance formula
func (s *RideSegment) GetVelocity(segment RideSegment) float64 {
	distance := formulas.CalculateHaversine(
		formulas.Point{
			Latitude:  s.Point.Latitude,
			Longitude: s.Point.Longitude,
		},
		formulas.Point{
			Latitude:  segment.Point.Latitude,
			Longitude: segment.Point.Longitude,
		},
	)
	distanceKm := distance / 1000

	timestampDifference := s.Timestamp.Sub(
		segment.Timestamp,
	)

	diff := timestampDifference.Hours()
	velocity := math.Abs(distanceKm / diff)
	if math.IsNaN(velocity) {
		return 0
	}
	return velocity
}

func (s *RideSegment) Equals(seg RideSegment) bool {
	if s.Point.Latitude != seg.Point.Latitude {
		return false
	}
	if s.Point.Longitude != seg.Point.Longitude {
		return false
	}
	if s.Timestamp != seg.Timestamp {
		return false
	}
	return true
}
