package formulas

import "math"

//
// Formula
// a = sin²(Δφ/2) + cos φ1 ⋅ cos φ2 ⋅ sin²(Δλ/2)
//

// EarthRadius is the approximate radius of earth in METERS
var EarthRadius float64 = 6371000

// Point contains the latitude and the longitude of a point
type Point struct {
	Latitude  float64
	Longitude float64
}

// hsin(φ) returns sin²(theta/2)
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// CalculateHaversine calculates the distance between two points in METERS
// considering that they're located on the earth with radius of 6371000 meters
//
// Returns the distance between the two points
// and the midpoint
// More information: http://en.wikipedia.org/wiki/Haversine_formula
func CalculateHaversine(pointA Point, pointB Point) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2 float64
	la1 = pointA.Latitude * math.Pi / 180
	lo1 = pointA.Longitude * math.Pi / 180
	la2 = pointB.Latitude * math.Pi / 180
	lo2 = pointB.Longitude * math.Pi / 180

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * EarthRadius * math.Asin(math.Sqrt(h))
}
