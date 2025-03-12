package helpers

import (
	"io"
	"time"

	ics "github.com/arran4/golang-ical"
	"middleware/example/internal/models"
)

// ParseICS lit un flux ICS et renvoie la liste des événements.
func ParseICS(r io.Reader) ([]models.Event, error) {
	cal, err := ics.ParseCalendar(r)
	if err != nil {
		return nil, err
	}

	var events []models.Event

	// Utiliser la méthode Events() qui renvoie une slice de *ics.VEvent
	for _, ev := range cal.Events() {
		evt := models.Event{
			UID:      ev.GetProperty("UID").Value,
			Summary:  ev.GetProperty("SUMMARY").Value,
			Location: ev.GetProperty("LOCATION").Value,
			// GetStartAt() et GetEndAt() renvoient des time.Time
			StartTime: func() time.Time {
				start, _ := ev.GetStartAt()
				return start
			}(),
			EndTime: func() time.Time {
				end, _ := ev.GetEndAt()
				return end
			}(),
		}
		events = append(events, evt)
	}

	return events, nil
}
