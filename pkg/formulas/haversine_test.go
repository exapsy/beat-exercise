package formulas_test

import (
	"testing"

	"github.com/exapsy/beat-exercise/pkg/formulas"
)

func TestHaversine(t *testing.T) {
	t.Run("Distance=1.286km", func(t *testing.T) {
		startPoint := formulas.Point{
			Latitude:  49.1715000,
			Longitude: -121.7493500,
		}
		endPoint := formulas.Point{
			Latitude:  49.18258,
			Longitude: -121.75441,
		}
		distance := formulas.CalculateHaversine(startPoint, endPoint)

		if distance != 1267.4277181680873 {
			t.Error("Expected distance 1267.4277181680873 meters but got", distance, "meters")
		}
	})
	t.Run("Distance=0", func(t *testing.T) {
		startPoint := formulas.Point{
			Latitude:  0,
			Longitude: 0,
		}
		endPoint := formulas.Point{
			Latitude:  0,
			Longitude: 0,
		}
		distance := formulas.CalculateHaversine(startPoint, endPoint)

		if distance != 0 {
			t.Error("Expected distance 0 meters but got", distance, "meters")
		}
	})
}
