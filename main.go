package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/isaacp/bookd/entities"
	"github.com/isaacp/collections/stack"
)

func main() {
	calendars := make([]entities.Calendar, 0)
	inflateCalendars("calendars.json", &calendars)
	eventFilter := make(map[string]bool)

	events := make([]entities.Event, 0)

	begin := time.Date(2024, 2, 7, 0, 0, 0, 0, time.Now().Location())
	end := time.Date(2024, 2, 8, 0, 0, 0, 0, time.Now().Location())

	for _, c := range calendars {
		for _, event := range c.Events {
			if !strings.Contains(strings.ToLower(event.Summary), "[canceled]") && !eventFilter[event.Id] && ((event.Start.After(begin) && event.Start.Before(end)) || (event.Start.Before(begin) && event.End.After(begin)) || (event.Start.Before(end) && event.End.After(end))) {
				events = append(events, *event)
				eventFilter[event.Id] = true
			}
		}
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Start.Before(events[j].Start)
	})

	intervals := stack.NewStack[entities.Interval]()

	for _, event := range events {
		fmt.Printf("%s: %s \n", event.Start, event.Summary)
		if intervals.IsEmpty() {
			intervals.Push(event.Interval())
			continue
		}

		top, _ := intervals.Peek()
		if event.Interval().Overlapping(*top) {
			merged := event.Interval().MergeWith(*top)
			intervals.Pop()
			intervals.Push(merged)
		} else {
			intervals.Push(event.Interval())
		}
	}

	start := begin
	finish := end
	freeIntervals := make([]entities.Interval, 0)
	intervalSlice := intervals.ToSlice()
	for index, interval := range intervalSlice {
		freeIntervals = append(freeIntervals, entities.Interval{
			Begin: start,
			End:   interval.Begin,
		})

		if index == len(intervalSlice)-1 {
			freeIntervals = append(freeIntervals, entities.Interval{
				Begin: interval.End,
				End:   finish,
			})
		} else {
			start = interval.End
		}
	}

	fmt.Printf("Height: %d\n", intervals.Height())

	for i := 0; i < len(freeIntervals); i++ {
		fmt.Println(freeIntervals[i])
	}
}

func inflateCalendars(file string, calendars *[]entities.Calendar) {
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
