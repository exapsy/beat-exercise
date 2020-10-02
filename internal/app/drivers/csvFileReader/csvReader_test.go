package csvfilereader_test

import (
	"bytes"
	"strconv"
	"testing"

	csvfilereader "github.com/exapsy/beat-exercise/internal/app/drivers/csvFileReader"
)

func TestReadFile(t *testing.T) {
	t.Run("read_ride_with_2_records", func(t *testing.T) {
		var buffer bytes.Buffer
		buffer.WriteString("1,37.966660,23.728308,1405594957\r\n")
		buffer.WriteString("1,37.966627,23.728263,1405594966\r\n")
		buffer.WriteString("2,37.966627,23.728263,1405594966\r\n")

		ride, err := csvfilereader.ReadRide(&buffer)
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
}
