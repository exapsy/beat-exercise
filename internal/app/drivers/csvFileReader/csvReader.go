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
	file *os.File
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
	return ReadRide(f.file)
}

// ReadRide returns a ride from a buffer with ride segments
// assuming that the file it's continuous and pre-sorted per-ride
func ReadRide(reader io.Reader) (ride *models.Ride, err error) {
	records := [][]string{}
	csvReader := csv.NewReader(reader)
	var rideID string

	// Read and parse to string records
	for {
		record, errTmp := csvReader.Read()
		if errTmp != nil {
			break
		}
		if rideID != "" && rideID != string(record[0][0]) {
			break
		} else if rideID == "" {
			rideID = string(record[0][0])
		}
		records = append(records, record)
	}

	ride = &models.Ride{
		ID:       rideID,
		Segments: []models.RideSegment{},
	}

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

		ride.Segments = append(
			ride.Segments,
			models.RideSegment{
				Timestamp: time.Unix(timestamp, 0),
				Point: models.Point{
					Latitude:  la,
					Longitude: lo,
				},
			},
		)
	}
	return
}
