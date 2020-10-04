package csvutils_test

import (
	"bytes"
	"encoding/csv"
	"strconv"
	"testing"
	"time"

	"github.com/exapsy/beat-exercise/internal/app/drivers/csvutils"
)

func TestReadFile(t *testing.T) {
	t.Run("read_ride_values", func(t *testing.T) {
		var buffer bytes.Buffer
		buffer.WriteString("1,37.966660,23.728308,1405594957\r\n")
		buffer.WriteString("1,37.966627,23.728263,1405594966\r\n")
		buffer.WriteString("2,38.966223,22.728269,1405594968\r\n")

		ride, err := csvutils.ReadRide(&buffer, nil)
		if err != nil {
			t.Fatal(err)
		}
		if ride.ID != "1" {
			t.Error("Expected id of 1 but got", ride.ID)
		}
		if len(ride.Segments) != 2 {
			t.Error("Expected segments array with length of 2 but got", strconv.Itoa(len(ride.Segments)))
		}
		if ride.Segments[0].Point.Latitude != 37.966660 {
			t.Error("Expected 37.966660 latitude but got", ride.Segments[0].Point.Latitude)
		}
		if ride.Segments[0].Point.Longitude != 23.728308 {
			t.Error("Expected 23.728308 longitude but got", ride.Segments[0].Point.Longitude)
		}
		if ride.Segments[0].Timestamp.Unix() != 1405594957 {
			t.Error("Expected 1405594957 timestamp but got", ride.Segments[0].Timestamp.Unix())
		}
		if ride.Segments[1].Point.Latitude != 37.966627 {
			t.Error("Expected 37.966627 latitude but got", ride.Segments[1].Point.Latitude)
		}
		if ride.Segments[1].Point.Longitude != 23.728263 {
			t.Error("Expected 23.728263 longitude but got", ride.Segments[1].Point.Longitude)
		}
		if ride.Segments[1].Timestamp.Unix() != 1405594966 {
			t.Error("Expected 1405594966 timestamp but got", ride.Segments[1].Timestamp.Unix())
		}
	})
	t.Run("read_3_rides", func(t *testing.T) {
		// Reading 3 rides prevents nextRideFirstRecord from being stuck to the previous value
		var buffer bytes.Buffer
		buffer.WriteString("1,37.966660,23.728308,1405594957\r\n")
		buffer.WriteString("1,37.966627,23.728263,1405594966\r\n")
		buffer.WriteString("2,37.966627,23.728263,1405594968\r\n")
		buffer.WriteString("2,37.966627,23.728263,1405594969\r\n")
		buffer.WriteString("3,37.966627,23.728263,1405594973\r\n")
		buffer.WriteString("3,37.966627,23.728263,1405594975\r\n")
		buffer.WriteString("3,37.966627,23.728263,1405594976\r\n")

		// Keep the cursor
		csvReader := csv.NewReader(&buffer)
		ride, err := csvutils.ReadRide(&buffer, csvReader)
		if err != nil {
			t.Fatal(err)
		}
		if ride.ID != "1" {
			t.Error("Expected id of 1 but got", ride.ID)
		}
		if len(ride.Segments) != 2 {
			t.Error("Expected segments array with length of 2 but got", strconv.Itoa(len(ride.Segments)))
		}
		ride, err = csvutils.ReadRide(&buffer, csvReader)
		if err != nil {
			t.Fatal(err)
		}
		if ride.ID != "2" {
			t.Error("Expected id of 2 but got", ride.ID)
		}
		if len(ride.Segments) != 2 {
			t.Error("Expected segments array with length of 2 but got", strconv.Itoa(len(ride.Segments)))
		}
		ride, err = csvutils.ReadRide(&buffer, csvReader)
		if err != nil {
			t.Fatal(err)
		}
		if ride.ID != "3" {
			t.Error("Expected id of 3 but got", ride.ID)
		}
		if len(ride.Segments) != 3 {
			t.Error("Expected segments array with length of 3 but got", strconv.Itoa(len(ride.Segments)))
		}
	})
	t.Run("read_500_segments_of_4_rides", func(t *testing.T) {
		var buffer bytes.Buffer
		timestamp := time.Time{}
		for i := 0; i < 120; i++ {
			timestampStr := strconv.FormatInt(timestamp.Unix(), 10)
			buffer.WriteString("1,37.966660,23.728308," + timestampStr + "\r\n")
			timestamp = timestamp.Add(time.Second * 1)
		}
		for i := 0; i < 280; i++ {
			timestampStr := strconv.FormatInt(timestamp.Unix(), 10)
			buffer.WriteString("2,37.966660,23.728308," + timestampStr + "\r\n")
			timestamp = timestamp.Add(time.Second * 1)
		}
		for i := 0; i < 100; i++ {
			timestampStr := strconv.FormatInt(timestamp.Unix(), 10)
			buffer.WriteString("3,37.966660,23.728308," + timestampStr + "\r\n")
			timestamp = timestamp.Add(time.Second * 1)
		}
		for i := 0; i < 200; i++ {
			timestampStr := strconv.FormatInt(timestamp.Unix(), 10)
			buffer.WriteString("4,37.966660,23.728308," + timestampStr + "\r\n")
			timestamp = timestamp.Add(time.Second * 1)
		}

		// Keep the cursor
		csvReader := csv.NewReader(&buffer)

		ride, err := csvutils.ReadRide(&buffer, csvReader)
		if err != nil {
			t.Fatal(err)
		}
		if ride.ID != "1" {
			t.Error("Expected id of 1 but got", ride.ID)
		}
		if len(ride.Segments) != 120 {
			t.Error("Expected segments array with length of 30 but got", strconv.Itoa(len(ride.Segments)))
		}

		ride, err = csvutils.ReadRide(&buffer, csvReader)
		if err != nil {
			t.Fatal(err)
		}
		if ride.ID != "2" {
			t.Error("Expected id of 2 but got", ride.ID)
		}
		if len(ride.Segments) != 280 {
			t.Error("Expected segments array with length of 70 but got", strconv.Itoa(len(ride.Segments)))
		}

		ride, err = csvutils.ReadRide(&buffer, csvReader)
		if err != nil {
			t.Fatal(err)
		}
		if ride.ID != "3" {
			t.Error("Expected id of 3 but got", ride.ID)
		}
		if len(ride.Segments) != 100 {
			t.Error("Expected segments array with length of 100 but got", strconv.Itoa(len(ride.Segments)))
		}

		ride, err = csvutils.ReadRide(&buffer, csvReader)
		if err != nil {
			t.Fatal(err)
		}
		if ride.ID != "4" {
			t.Error("Expected id of 4 but got", ride.ID)
		}
		if len(ride.Segments) != 200 {
			t.Error("Expected segments array with length of 200 but got", strconv.Itoa(len(ride.Segments)))
		}
	})
}
