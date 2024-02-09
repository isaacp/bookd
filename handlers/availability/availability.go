package availability

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isaacp/bookd/entities"
	"github.com/isaacp/bookd/handlers"
	"github.com/isaacp/collections/stack"
)

func List(c *gin.Context) {
	calendars := make([]entities.Calendar, 0)
	handlers.InflateCalendars("calendars.json", &calendars)
	eventFilter := make(map[string]bool)

	events := make([]entities.Event, 0)

	fmt.Println(c.Query("start"))
	begin, _ := time.Parse(time.RFC3339, c.Query("start"))
	end, _ := time.Parse(time.RFC3339, c.Query("end"))

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

	for i := 0; i < len(freeIntervals); i++ {
		fmt.Println(freeIntervals[i])
	}

	c.JSON(http.StatusOK, freeIntervals)
}
