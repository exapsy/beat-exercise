package main

import (
	"fmt"
	"log"
	"os"

	"github.com/exapsy/beat-exercise/internal/app/drivers/csvutils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Incorrect usage: app <csvfile> <optional: outputfile>")
		os.Exit(1)
	}
	csvFilePath := os.Args[1]

	file, err := csvutils.OpenFile(csvFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var outputFile *os.File = nil
	if len(os.Args) == 3 {
		fileName := os.Args[2]
		outputFile, err = os.Create(fileName)
		if err != nil {
			panic(err)
		}
		defer outputFile.Close()
	}
	for {
		ride, err := file.ReadRide()
		if err == csvutils.ErrEOF {
			fmt.Println("Done!")
			os.Exit(0)
		}
		if err != nil {
			log.Fatal(err)
		}
		if outputFile != nil {
			csvutils.WriteRide(outputFile, *ride)
		} else {
			ridesOutput := csvutils.GetRideOutputString(*ride)
			fmt.Println(ridesOutput)
		}
	}
}
