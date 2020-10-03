package csvutils

import (
	"os"
	"strconv"

	"github.com/exapsy/beat-exercise/internal/app/models"
)

// WriteRide outputs the ride estimates along with their id in a file
func WriteRide(file *os.File, ride models.Ride) (err error) {
	if err != nil {
		return err
	}

	rideOutput := GetRideOutputString(ride)
	file.WriteString(rideOutput)

	return
}

// GetRidesOutputString returns a ride's string output
func GetRidesOutputString(rides []models.Ride) (ridesOutput string) {
	for _, ride := range rides {
		ridesOutput += GetRideOutputString(ride)
	}
	return
}

// GetRideOutputString returns a ride's string output
func GetRideOutputString(ride models.Ride) (rideOutput string) {
	fare := float64(ride.EstimateFare())
	rideOutput = ride.ID +
		", " +
		strconv.FormatFloat(fare, 'g', -1, 32) +
		"\r\n"
	return
}
