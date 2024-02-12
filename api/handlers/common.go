package handlers

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/isaacp/bookd/api/core/entities"
)

func InflateCalendars(file string, calendars *[]entities.Calendar) {
	jsonFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	bytes, _ := io.ReadAll(jsonFile)

	err = json.Unmarshal(bytes, calendars)
	if err != nil {
		log.Fatal(err)
	}

	for index, calendar := range *calendars {
		(*calendars)[index] = *calendar.Process()
	}
}
