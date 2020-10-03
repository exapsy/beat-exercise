package models

import "fmt"

// Ride describes a ride of a taxi driver
type Ride struct {
	ID       string
	Segments []RideSegment
}

// MakeRide is a factory for a ride
func MakeRide(id string, segments []RideSegment) (ride *Ride) {
	ride = &Ride{
		ID:       id,
		Segments: segments[:1],
	}
	if len(segments) < 2 {
		return
	}
	ride.Segments = append(ride.Segments, filterValidSegments(segments)...)

	return
}

func filterValidSegments(segments []RideSegment) (filteredSegments []RideSegment) {
	filteredSegments = []RideSegment{}
	// The next compared index should be compared to the last compared one
	lastComparedIndex := 0
	for i := 1; i < len(segments); i++ {
		if segments[i].GetVelocity(segments[lastComparedIndex]) <= 100 {
			filteredSegments = append(filteredSegments, segments[i])
			lastComparedIndex = i
		}
		if segments[lastComparedIndex].Timestamp.After(segments[i].Timestamp) {
			fmt.Println("[Warning] A ride segment had lower timestamp than the previous segment")
		}
	}
	return
}
