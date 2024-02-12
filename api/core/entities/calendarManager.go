package entities

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

	for index, calendar := range calendars {
		calendars[index] = *calendar.Process()
	}

	return calendars
}
