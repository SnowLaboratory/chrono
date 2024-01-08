package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	file, err := os.Open("event_schedule.csv")
	if err != nil {
		log.Fatal("error opening  csv file", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

	rows := []map[string]string{}
	var header []string
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if header == nil {
			header = record
		} else {
			dict := map[string]string{}
			for i := range header {
				dict[header[i]] = record[i]
			}
			rows = append(rows, dict)
		}
	}

	newFile, err := os.Create("event_data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	jsonString, err := json.Marshal(rows)

	if err != nil {
		log.Fatal(err)
	}

	newFile.Write(jsonString)

	router := gin.Default()
	router.GET("/events", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.File("event_data.json")
	})
	router.Run()
}
