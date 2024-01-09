package main

import (
	"log"
	"snowlabs/chrono/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	_, err := lib.ExportCsvToJson()
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	router.GET("/events", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.File("event_data.json")
	})
	router.Run()
}
