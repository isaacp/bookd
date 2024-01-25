package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	ics "github.com/arran4/golang-ical"
	"github.com/isaacp/bookd/calendar"
)

func main() {

	re := regexp.MustCompile(`(?P<year>[0-9]{4})(?P<month>[0-1]{1}[0-9]{1})(?P<date>[0-3]{1}[0-9]{1})T(?P<hour>[0-5]{1}[0-9]{1})(?P<minute>[0-5]{1}[0-9]{1})(?P<seconds>[0-5]{1}[0-9]{1})`)
	calendars := make([]calendar.Calendar, 0)
	jsonFile, err := os.Open("calendars.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	bytes, _ := io.ReadAll(jsonFile)

	err = json.Unmarshal(bytes, &calendars)
	if err != nil {
		fmt.Println(err)
	}

	for i, c := range calendars {
		resp, err := http.Get(c.Location)
		if err != nil {
			fmt.Println(err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		calendars[i].Data, err = ics.ParseCalendar(strings.NewReader(string(body)))
		if err != nil {
			fmt.Println(err)
		}
	}

	for _, c := range calendars {
		fmt.Println(c.Name)
		for _, event := range c.Data.Events() {
			start := event.GetProperty(ics.ComponentPropertyDtStart).Value
			if re.FindAllStringSubmatch(start, -1) != nil {
				fmt.Println("Okay, now we're talking!")
			}
			end := event.GetProperty(ics.ComponentPropertyDtEnd).Value

			fmt.Println(re.FindStringSubmatch(start))
			fmt.Println(re.SubexpNames())

			fmt.Printf("%s: %s | %s - %s\n", c.Name, event.GetProperty(ics.ComponentPropertySummary).Value, start, end)
			// if r, _ := event.GetEndAt(); r.After(time.Now()) {
			// 	fmt.Printf("%s: %s @ %s\n", c.Name, event.GetProperty(ics.ComponentPropertySummary).Value, r)
			// }
		}
	}

	// for _, c := range calendars {
	// 	fmt.Printf("%s: %d\n", c.Name, len(c.Data.Events()))
	// }
}
