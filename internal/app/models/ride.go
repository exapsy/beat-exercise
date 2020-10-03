package models

import (
	"fmt"
	"time"
)

// Ride describes a ride of a taxi driver
type Ride struct {
	ID       string
	Segments []RideSegment
}

// FareFlag is the standard charge for each start of a ride
const FareFlag float32 = 1.30

// FareMinimum is the minimum fare for a ride
const FareMinimum float32 = 3.47

// FareMidnightPerKm is the fare applied in midnight per each killometre
const FareMidnightPerKm float32 = 1.30

// FareMiddayPerKm is the fare applied in miday per each killometre
const FareMiddayPerKm float32 = 0.74

// FareIdlePerHour is the fare applied for each hour staying idle
const FareIdlePerHour float32 = 11.90

// MakeRide is a factory for a ride
func MakeRide(id string, segments []RideSegment) (ride *Ride) {
	if segments == nil || len(segments) == 0 {
		segments = []RideSegment{}
	}
	ride = &Ride{
		ID: id,
	}
	if len(segments) >= 1 {
		ride.Segments = segments[:1]
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
			fmt.Println("[Warning] Potential desync smell: A ride segment had lower timestamp than the previous segment")
		}
	}
	return
}

// EstimateFare returns the ride's fare
// Please use this instead of trying to cheat the customer Mr.Taxi driver :(
func (r *Ride) EstimateFare() (fare float32) {
	fare = FareFlag

	if len(r.Segments) < 2 {
		return FareMinimum
	}

	var totalIdleTime = time.Time{}
	var totalKmMidnight float64 = 0
	var totalKmMidday float64 = 0

	// Aggregate through all the segments
	for i, segment := range r.Segments[1:] {
		previousSegment := r.Segments[i]
		// Todo: If previous segment is on 4:58 and next segment on 5:05
		// it should probably seperate segments in appropriate pieces
		// to estimate a fair fare
		if segment.GetVelocity(previousSegment) > 10 {
			hour, _, _ := segment.Timestamp.Clock()
			isMidnight := hour <= 5 &&
				hour > 0
			if isMidnight {
				totalKmMidnight += segment.DistanceFrom(previousSegment)
			} else {
				totalKmMidday += segment.DistanceFrom(previousSegment)
			}
		} else if segment.GetVelocity(previousSegment) <= 10 {
			totalIdleTime = totalIdleTime.Add(
				time.Duration(segment.Timestamp.Sub(previousSegment.Timestamp)),
			)
		}
	}

	// Idle time fare
	fare += float32(totalIdleTime.Hour()) * FareIdlePerHour
	// Midnight fare
	fare += float32(totalKmMidnight) * FareMidnightPerKm
	// Midday fare
	fare += float32(totalKmMidday) * FareMiddayPerKm

	if fare < FareMinimum {
		fare = FareMinimum
	}

	return
}
