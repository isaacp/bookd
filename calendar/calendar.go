package calendar

import ics "github.com/arran4/golang-ical"

type Calendar struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Data     *ics.Calendar
}
