package lib

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
)

func ExportCsvToJson() (bool, error) {
	reader, file, err := Reader()
	if err != nil {
		return false, err
	}

	defer file.Close()
	rows, err := CsvMapper(reader)
	if err != nil {
		return false, nil
	}
	_, err = CsvToJson(rows)

	if err != nil {
		return false, nil
	}

	return true, nil
}

func Reader() (*csv.Reader, *os.File, error) {
	file, err := os.Open("event_schedule.csv")
	if err != nil {
		log.Fatal("error opening  csv file", err)
		return nil, nil, err
	}

	csvReader := csv.NewReader(file)

	return csvReader, file, nil
}

func CsvMapper(reader *csv.Reader) ([]map[string]string, error) {
	rows := []map[string]string{}
	var header []string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
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
	return rows, nil
}

func CsvToJson(rows []map[string]string) (bool, error) {
	newFile, err := os.Create("event_data.json")
	if err != nil {
		return false, err
	}
	defer newFile.Close()

	jsonString, err := json.Marshal(rows)

	if err != nil {
		return false, err
	}

	_, err = newFile.Write(jsonString)

	if err != nil {
		return false, err
	}
	return true, nil
}
