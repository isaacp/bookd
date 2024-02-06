package entities

import (
	"log"
	"regexp"
	"strconv"
	"time"

	ics "github.com/arran4/golang-ical"
)

type (
	Event struct {
		Id           string    `json:"id"`
		Start        time.Time `json:"start"`
		End          time.Time `json:"end"`
		Summary      string    `json:"summary"`
		Location     string    `json:"location"`
		LastModified string    `json:"last_modified"`
	}
)

func Convert(ve *ics.VEvent) Event {
	return Event{
		Id:           getProperty(ve, ics.ComponentPropertyUniqueId),
		Start:        getTime(getProperty(ve, ics.ComponentPropertyDtStart)),
		End:          getTime(getProperty(ve, ics.ComponentPropertyDtEnd)),
		Summary:      getProperty(ve, ics.ComponentPropertySummary),
		Location:     getProperty(ve, ics.ComponentPropertyLocation),
		LastModified: getProperty(ve, ics.ComponentPropertyLastModified),
	}
}

func getProperty(ve *ics.VEvent, prop ics.ComponentProperty) string {
	if ve.GetProperty(prop) != nil {
		return ve.GetProperty(prop).Value
	} else {
		return ""
	}
}

func getTime(timeStr string) time.Time {
	re := regexp.MustCompile(`(?P<year>[0-9]{4})(?P<month>[0-1]{1}[0-9]{1})(?P<date>[0-3]{1}[0-9]{1})(T(?P<hour>[0-5]{1}[0-9]{1})(?P<minutes>[0-5]{1}[0-9]{1})(?P<seconds>[0-5]{1}[0-9]{1}))?`)
	if matches := re.FindAllStringSubmatch(timeStr, -1); matches == nil {
		log.Panicf("substring patterns do not match: %s", timeStr)
	} else {
		year, _ := strconv.Atoi(matches[0][re.SubexpIndex("year")])
		month, _ := strconv.Atoi(matches[0][re.SubexpIndex("month")])
		date, _ := strconv.Atoi(matches[0][re.SubexpIndex("date")])
		hour := 0
		minutes := 0
		seconds := 0
		nano := 0
		if matches[0][re.SubexpIndex("hour")] != "" {
			hour, _ = strconv.Atoi(matches[0][re.SubexpIndex("hour")])
			minutes, _ = strconv.Atoi(matches[0][re.SubexpIndex("minutes")])
			//seconds, _ = strconv.Atoi(matches[0][re.SubexpIndex("seconds")])
		}
		return time.Date(year, time.Month(month), date, hour, minutes, seconds, nano, time.Now().Local().Location())
	}

	return time.Time{}
}
