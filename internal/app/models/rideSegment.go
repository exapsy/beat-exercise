package models

import (
	"math"
	"time"
)

// RideSegment describes either the start or the end of a ride
type RideSegment struct {
	Point     Point
	Timestamp time.Time
}

// GetVelocity returns the km/h calculated
// between the two segments of a ride
// using the Haversine distance formula
func (s *RideSegment) GetVelocity(segment RideSegment) (velocity float64) {
	distanceKm := s.DistanceFrom(segment)

	timestampDifference := s.Timestamp.Sub(
		segment.Timestamp,
	)

	diff := timestampDifference.Hours()
	velocity = math.Abs(distanceKm / diff)
	if math.IsNaN(velocity) {
		velocity = 0
	}
	return velocity
}

// DistanceFrom returns the distance in KILOMETRES between two segments.
// Distance is measured by haversine formula
func (s *RideSegment) DistanceFrom(segment RideSegment) (distance float64) {
	distance = s.Point.HaversineDistanceFrom(
		segment.Point,
	) / 1000

	return distance
}

// Equals tests all the values of two segments
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
