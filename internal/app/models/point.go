package models

import "github.com/exapsy/beat-exercise/pkg/formulas"

// Point contains the latitude and the longitude of a point
type Point struct {
	Latitude  float64
	Longitude float64
}

// HoversineDistanceFrom returns the hoversine distance in meters between two points
func (p *Point) HoversineDistanceFrom(p2 Point) (distance float64) {
	distance = formulas.CalculateHaversine(
		formulas.Point{
			Latitude:  p.Latitude,
			Longitude: p.Longitude,
		},
		formulas.Point{
			Latitude:  p2.Latitude,
			Longitude: p2.Longitude,
		},
	)
	return
}
