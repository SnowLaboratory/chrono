package helpers

import (
	"log"
	"strconv"
	"strings"
	"time"
)

func RemoveUnderscores(name string) string {
	words := strings.Split(name, "_")
	return strings.Join(words, " ")
}

func UnixTime(timeString string) string {
	i, err := strconv.ParseInt(timeString, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	tm := time.Unix(i, 0)
	return tm.Format("2006-01-02 15:04:05 MST")
}
