package csvutils

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

// InputFile is an input file structure with meta data
type InputFile struct {
	file   *os.File
	reader *csv.Reader
}

// OpenFile opens a file by filename path
func OpenFile(path string) (file InputFile, err error) {
	fileStream, err := os.Open(path)
	if err != nil {
		return
	}
	file = InputFile{
		file:   fileStream,
		reader: csv.NewReader(fileStream),
	}
	return
}

// Close closes the file buffer
func (f *InputFile) Close() {
	f.file.Close()
}

// ReadRide returns a ride
func (f *InputFile) ReadRide() (ride *models.Ride, err error) {
	return ReadRide(f.file, f.reader)
}

type previousRecordType struct {
	reader *csv.Reader
	record []string
	// Used for EOF, so the last ride in buffer will be returned
	// and will not output an error
	lastError error
}

var previousRecord previousRecordType = previousRecordType{}

func (p *previousRecordType) Set(reader *csv.Reader, record []string) {
	previousRecord.reader = reader
	previousRecord.record = append([]string{}, record...)
}

// ReadRide returns a ride from a buffer with ride segments
// assuming that the file it's continuous and pre-sorted per-ride
func ReadRide(
	reader io.Reader,
	csvReader *csv.Reader,
) (ride *models.Ride, err error) {
	var rideID string

	if csvReader == nil {
		csvReader = csv.NewReader(reader)
		csvReader.FieldsPerRecord = 4
	}

	jobs := make(chan []string)

	// Read records and parse to ride segments
	rideResult := make(chan *models.Ride)
	go parseCSVRecordsToRide(jobs, rideResult)

	if previousRecord.record != nil &&
		len(previousRecord.record) > 0 &&
		previousRecord.reader == csvReader {
		_previousRecord := append([]string{}, previousRecord.record...)
		jobs <- _previousRecord
		previousRecord.record = []string{}
	}

	// Read and parse to string records
	for {
		record, errTmp := csvReader.Read()
		if previousRecord.lastError == errTmp &&
			errTmp == io.EOF {
			return nil, ErrEOF
		}
		if errTmp != nil {
			previousRecord.lastError = errTmp
			break
		}

		// Reallocate reused pointers
		previousRecord.Set(csvReader, record)
		record = append([]string{}, record...)

		if rideID == "" {
			rideID = record[0]
		}
		if rideID != "" && rideID != record[0] {
			break
		}

		jobs <- record
	}
	close(jobs)

	ride = <-rideResult
	return
}

func parseCSVRecordsToRide(jobs <-chan []string, rideResult chan *models.Ride) {
	segments := []models.RideSegment{}
	var rideID string = ""
	for {
		record, ok := <-jobs
		if !ok {
			break
		}
		if rideID == "" {
			rideID = record[0]
		}
		la, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			panic("Couldn't parse latitude")
		}
		lo, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			panic("Couldn't parse longitude")
		}
		timestamp, err := strconv.ParseInt(record[3], 10, 64)
		if err != nil {
			panic("Coulnd't parse timestamp")
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
	rideResult <- models.MakeRide(
		rideID,
		segments,
	)
	return
}
