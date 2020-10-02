package models

import "time"

// RideSegment describes either the start or the end of a ride
type RideSegment struct {
	Point     Point
	Timestamp time.Time
}
