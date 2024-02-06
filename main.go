package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/isaacp/bookd/entities"
)

func main() {
	calendars := make([]entities.Calendar, 0)

	jsonFile, err := os.Open("calendars.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	bytes, _ := io.ReadAll(jsonFile)

	err = json.Unmarshal(bytes, &calendars)
	if err != nil {
		log.Fatal(err)
	}

	for index, calendar := range calendars {
		calendars[index] = *calendar.Process()
	}

	for _, c := range calendars {
		fmt.Println(c.Name)
		for _, event := range c.Events {
			fmt.Printf("%s: %s | %s - %s\n", c.Name, event.Summary, event.Start, event.End)
		}
	}
}
