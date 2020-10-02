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

// File contains the cursor and the file stream
type File struct {
	file   *os.File
	cursor int
}

// Close closes the file buffer
func (f *File) Close() {
	f.file.Close()
}

// ReadRide returns a ride
func (f *File) ReadRide() (ride *models.Ride, err error) {
	return readRide(f.file)
}

// readRide returns a ride from a buffer with ride segments
func readRide(reader io.Reader) (ride *models.Ride, err error) {
	records := [][]string{}
	csvReader := csv.NewReader(reader)
	for {
		record, errTmp := csvReader.Read()
		if errTmp != nil {
			break
		}
		records = append(records, record)
	}
	ride = models.Ride{
		ID:       records[0][0],
		Segments: []models.RideSegment{},
	}
	for _, record := range records {
		var la, lo float64 = record[1], record[2]

		timestamp, err := strconv.ParseInt(record[3], 10, 64)
		if err != nil {
			return nil, errors.New("Coulnd't parse timestamp")
		}

		ride.Segments = append(
			ride.Segments,
			models.RideSegment{
				Timestamp: time.Unix(timestamp, 0),
			},
		)
	}
	return
}

// OpenFile opens a file and initializes a cursor
func OpenFile(filePath string) (file File, err error) {
	fileStream, err := os.Open(filePath)
	if err != nil {
		return
	}
	file = File{
		file:   fileStream,
		cursor: 0,
	}
	return
}

// OpenAndReadFile opens & reads a csv file and returns its records
func OpenAndReadFile(filePath string) (records [][]string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	records, err = ReadFile(file)
	return
}

// ReadFile reads a csv file and returns its records
func ReadFile(reader io.Reader) (records [][]string, err error) {
	csvReader := csv.NewReader(reader)
	for {
		record, errTmp := csvReader.Read()
		if errTmp != nil {
			return
		}
		records = append(records, record)
	}
}
