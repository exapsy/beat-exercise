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
func (s *RideSegment) GetVelocity(previousSegment RideSegment) float64 {
	distance := formulas.CalculateHaversine(
		formulas.Point{
			Latitude:  s.Point.Latitude,
			Longitude: s.Point.Longitude,
		},
		formulas.Point{
			Latitude:  previousSegment.Point.Latitude,
			Longitude: previousSegment.Point.Longitude,
		},
	)
	distanceKm := distance / 1000

	timestampDifference := s.Timestamp.Sub(
		previousSegment.Timestamp,
	)

	diff := timestampDifference.Hours()
	velocity := math.Abs(distanceKm / diff)
	if math.IsNaN(velocity) {
		return 0
	}
	return velocity
}
