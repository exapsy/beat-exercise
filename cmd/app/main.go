package main

import (
	"fmt"
	"log"
	"os"

	csvfilereader "github.com/exapsy/beat-exercise/internal/app/drivers/csvFileReader"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Incorrect usage: app <csvfile>")
		os.Exit(1)
	}
	csvFilePath := os.Args[1]

	csvfilereader.OpenFile(csvFilePath)
	file, err := csvfilereader.OpenFile(csvFilePath)
	ride, err := file.ReadRide()
	if err == csvfilereader.ErrEOF {
		os.Exit(0)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ride)
}
