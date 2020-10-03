package csvfilereader

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/exapsy/beat-exercise/internal/app/models"
)

// ErrEOF signifies a file error when there's nothing else to read
var ErrEOF = errors.New("End of file")

// File is a file
type File struct {
	file   *os.File
	reader *csv.Reader
	// Keep the last read record so when the cursor has bypassed it
	// in case it's the next ride
	// we dont lose it
	nextRideFirstRecord []string
}

// OpenFile opens a file by filename path
func OpenFile(path string) (file File, err error) {
	fileStream, err := os.Open(path)
	if err != nil {
		return
	}
	file = File{
		file: fileStream,
	}
	return
}

// Close closes the file buffer
func (f *File) Close() {
	f.file.Close()
}

// ReadRide returns a ride
func (f *File) ReadRide() (ride *models.Ride, err error) {
	return ReadRide(f.file, f.reader, &f.nextRideFirstRecord)
}

// ReadRide returns a ride from a buffer with ride segments
// assuming that the file it's continuous and pre-sorted per-ride
func ReadRide(
	reader io.Reader,
	csvReader *csv.Reader,
	nextRideFirstRecord *[]string,
) (ride *models.Ride, err error) {
	records := [][]string{}
	if csvReader == nil {
		csvReader = csv.NewReader(reader)
	}
	var rideID string

	if nextRideFirstRecord != nil && len(*nextRideFirstRecord) > 0 {
		records = append(records, *nextRideFirstRecord)
		rideID = (*nextRideFirstRecord)[0]
		// Reset
		*nextRideFirstRecord = []string{}
	}

	// Read and parse to string records
	for {
		record, errTmp := csvReader.Read()
		if errTmp != nil {
			break
		}
		if rideID != "" && rideID != string(record[0][0]) {
			if nextRideFirstRecord != nil {
				*nextRideFirstRecord = record
			}
			break
		} else if rideID == "" {
			rideID = string(record[0][0])
		}
		records = append(records, record)
	}

	segments := []models.RideSegment{}

	// Read records and parse to ride segments
	for _, record := range records {
		la, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, errors.New("Couldn't parse latitude")
		}
		lo, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, errors.New("Couldn't parse longitude")
		}
		timestamp, err := strconv.ParseInt(record[3], 10, 64)
		if err != nil {
			return nil, errors.New("Coulnd't parse timestamp")
		}

		segments = append(
			segments,
			models.RideSegment{
				Timestamp: time.Unix(timestamp, 0),
				Point: models.Point{
					Latitude:  la,
					Longitude: lo,
				},
			},
		)
	}

	ride = models.MakeRide(
		rideID,
		segments,
	)
	return
}
