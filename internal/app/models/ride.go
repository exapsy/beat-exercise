package models

// Ride describes a ride of a taxi driver
type Ride struct {
	ID       string
	Segments []RideSegment
}
