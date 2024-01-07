package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	//open csv
	file, err := os.Open("event_schedule.csv")
	if err != nil {
		log.Fatal("error opening  csv file", err)
	}
	defer file.Close() //make sure to close file, no matter how the func exits

	// create new csv reader and read entire file
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("error reading csv file", err)
	}

	//loop over csv
	for _, value := range data {
		fmt.Println(value)
	}
}
