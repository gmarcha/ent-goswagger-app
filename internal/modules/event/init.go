package event

import (
	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/restapi/operations"
)

func Init(api *operations.TutorAPI, db *ent.Client) {

	eventService := &Service{Event: db.Event}

	api.EventListEventHandler = &listEvent{event: eventService}
	api.EventCreateEventHandler = &createEvent{event: eventService}

	api.EventReadEventHandler = &readEvent{event: eventService}
	api.EventUpdateEventHandler = &updateEvent{event: eventService}
	api.EventDeleteEventHandler = &deleteEvent{event: eventService}
	api.EventListEventUsersHandler = &listEventUsers{event: eventService}
	api.EventGetEventTypeHandler = &getEventType{event: eventService}
	api.EventSetEventTypeHandler = &setEventType{event: eventService}
}
