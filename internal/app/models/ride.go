package models

import (
	"github.com/exapsy/beat-exercise/pkg/formulas"
)

// Ride describes a ride of a taxi driver
type Ride struct {
	ID    string
	Start RideSegment
	End   RideSegment
}

// GetVelocity returns the km/h calculated
// between the two segments of a ride
// using the Haversine distance formula
func (r *Ride) GetVelocity() float64 {
	distance := formulas.CalculateHaversine(
		formulas.Point{
			Latitude:  r.Start.Point.Latitude,
			Longitude: r.Start.Point.Longitude,
		},
		formulas.Point{
			Latitude:  r.End.Point.Latitude,
			Longitude: r.End.Point.Longitude,
		},
	)

	timestampDifference := r.Start.Timestamp.Sub(
		r.End.Timestamp,
	)

	return distance / timestampDifference.Hours()
}
