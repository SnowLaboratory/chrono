package main

import (
	"encoding/csv"
	"log"
	"net/http"
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
	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("error reading csv file", err)
	}

	router := gin.Default()
	router.GET("/events", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, data)
	})
	router.Run()
}
