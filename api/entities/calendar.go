package entities

import (
	"io"
	"log"
	"net/http"
	"strings"

	ics "github.com/arran4/golang-ical"
)

type Calendar struct {
	Owner    string `json:"owner"`
	Name     string `json:"name"`
	Location string `json:"location"`
	data     *ics.Calendar
	Events   []*Event `json:"events"`
}

func (cal *Calendar) Process() *Calendar {

	resp, err := http.Get(cal.Location)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	cal.data, err = ics.ParseCalendar(strings.NewReader(string(body)))
	if err != nil {
		log.Fatal(err)
	}

	for _, ev := range cal.data.Events() {
		event := Convert(ev)
		cal.Events = append(cal.Events, &event)
	}

	return cal
}
