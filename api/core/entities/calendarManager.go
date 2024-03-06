package entities

import (
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/isaacp/collections/stack"
)

type (
	CalendarManager struct {
	}
)

func (cm *CalendarManager) GetCalendars() []Calendar {
	calendars := []Calendar{
		{
			Owner:    "isaac",
			Name:     "Isaac",
			Location: "https://rest.cozi.com/api/ext/1103/c8bb7533-fbe5-4db8-975f-f7b16a681888/icalendar/feed/feed.ics",
		},
		{
			Owner:    "kateshia",
			Name:     "Kateshia",
			Location: "https://rest.cozi.com/api/ext/1103/b4c67d49-dc7c-40bd-84a6-276473c91584/icalendar/feed/feed.ics",
		},
		{
			Owner:    "asharia",
			Name:     "Asharia",
			Location: "https://rest.cozi.com/api/ext/1103/0f7876ea-9375-44f9-8f1f-febb17d9b007/icalendar/feed/feed.ics",
		},
		{
			Owner:    "asharia",
			Name:     "metros08G",
			Location: "http://ical-cdn.teamsnap.com/team_schedule/47d3df72-1533-4cd1-88a0-f9d0b97c499f.ics",
		},
		{
			Owner:    "ikey",
			Name:     "Ikey",
			Location: "https://rest.cozi.com/api/ext/1103/1ae5172b-7a94-4919-9df2-1108e89abf94/icalendar/feed/feed.ics",
		},
		{
			Owner:    "ikey",
			Name:     "metros13B",
			Location: "http://ical-cdn.teamsnap.com/team_schedule/b04c7bda-684e-4c63-a1e8-a5215729a006.ics",
		},
		{
			Owner:    "ikey",
			Name:     "pressure13B",
			Location: "http://ical-cdn.teamsnap.com/team_schedule/e78d3328-dcb7-498f-870f-8cca74a40c79.ics",
		},
	}

	wg := sync.WaitGroup{}
	wg.Add(len(calendars))
	for index, calendar := range calendars {
		go func(ndx int, cal Calendar) {
			defer wg.Done()
			calendars[ndx] = *cal.Process()
		}(index, calendar)
	}
	wg.Wait()

	return calendars
}

func (cm *CalendarManager) FreeIntervals(_start, _finish string) []Interval {
	calendars := cm.GetCalendars()
	eventFilter := make(map[string]bool)

	events := make([]Event, 0)

	begin, _ := time.Parse(time.RFC3339, _start)
	end, _ := time.Parse(time.RFC3339, _finish)

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

	intervals := stack.NewStack[Interval]()

	for _, event := range events {
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
	freeIntervals := make([]Interval, 0)
	intervalSlice := intervals.ToSlice()

	for index, interval := range intervalSlice {
		freeIntervals = append(freeIntervals, Interval{
			Begin: start,
			End:   interval.Begin,
		})

		if index == len(intervalSlice)-1 {
			freeIntervals = append(freeIntervals, Interval{
				Begin: interval.End,
				End:   finish,
			})
		} else {
			start = interval.End
		}
	}

	return freeIntervals
}

func (cm *CalendarManager) Events(_start, _finish string) []Event {
	calendars := cm.GetCalendars()
	eventFilter := make(map[string]bool)

	events := make([]Event, 0)

	begin, _ := time.Parse(time.RFC3339, _start)
	end, _ := time.Parse(time.RFC3339, _finish)

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

	return events

}
