package events

import (
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isaacp/bookd/entities"
	"github.com/isaacp/bookd/handlers"
)

func List(c *gin.Context) {
	calendars := make([]entities.Calendar, 0)
	handlers.InflateCalendars("calendars.json", &calendars)
	eventFilter := make(map[string]bool)

	events := make([]entities.Event, 0)

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

	c.JSON(http.StatusOK, events)
}
