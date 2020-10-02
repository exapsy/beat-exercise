package csvfilereader_test

import (
	"bytes"
	"strconv"
	"testing"

	csvfilereader "github.com/exapsy/beat-exercise/internal/app/drivers/csvFileReader"
)

func TestReadFile(t *testing.T) {
	t.Run("TwoRecords", func(t *testing.T) {
		var buffer bytes.Buffer
		buffer.WriteString("1,37.966660,23.728308,1405594957\r\n")
		buffer.WriteString("1,37.966627,23.728263,1405594966\r\n")

		records, err := csvfilereader.ReadFile(&buffer)
		if err != nil {
			t.Fatal(err)
		}
		if len(records) < 2 {
			t.Fatal("Expected length of 2 records but got ", strconv.Itoa(len(records)))
		}
		// Todo: Not too practical to test everything individually
		if records[0][0] != "1" {
			t.Error("Expected 1 ID but got", records[0][0])
		}
		if records[0][1] != "37.966660" {
			t.Error("Expected 37.966660 latitude but got", records[0][1])
		}
		if records[0][2] != "23.728308" {
			t.Error("Expected 23.728308 longitude but got", records[0][1])
		}
		if records[0][3] != "1405594957" {
			t.Error("Expected 1405594957 timestamp but got", records[0][1])
		}
		if records[1][0] != "1" {
			t.Error("Expected 1 ID but got", records[0][0])
		}
		if records[1][1] != "37.966627" {
			t.Error("Expected 37.966627 latitude but got", records[0][1])
		}
		if records[1][2] != "23.728263" {
			t.Error("Expected 23.728263 longitude but got", records[0][1])
		}
		if records[1][3] != "1405594966" {
			t.Error("Expected 1405594957 timestamp but got", records[0][1])
		}
	})
}
