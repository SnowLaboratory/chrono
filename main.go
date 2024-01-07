package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("event_schedule.csv")
	if err != nil {
		log.Fatal("error opening  csv file", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("error reading csv file", err)
	}

	for _, value := range data {
		fmt.Println(value)
	}
}
