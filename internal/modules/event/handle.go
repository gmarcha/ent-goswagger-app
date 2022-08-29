package event

import (
	"context"

	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/restapi/operations/event"
	u "github.com/gmarcha/ent-goswagger-app/internal/utils"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
)

type listEvent struct {
	event *Service
}

func (e *listEvent) Handle(params event.ListEventParams, principal *models.Principal) middleware.Responder {
	ctx := context.Background()
	res, err := e.event.ListEvent(ctx, params.Start, params.End)
	if err != nil {
		return event.NewListEventInternalServerError().WithPayload(u.Err(500, err))
	}
	return event.NewListEventOK().WithPayload(res)
}

type listEventWithUsers struct {
	event *Service
}

func (e *listEventWithUsers) Handle(params event.ListEventWithUsersParams, principal *models.Principal) middleware.Responder {
	ctx := context.Background()
	res, err := e.event.ListEventWithUsers(ctx, params.Start, params.End)
	if err != nil {
		return event.NewListEventWithUsersInternalServerError().WithPayload(u.Err(500, err))
	}
	return event.NewListEventWithUsersOK().WithPayload(res)
}

type createEvent struct {
	event *Service
}

func (e *createEvent) Handle(params event.CreateEventParams, principal *models.Principal) middleware.Responder {
	ctx := context.Background()
	res, err := e.event.CreateEvent(ctx, params.Event)
	if err != nil {
		if ent.IsValidationError(err) || ent.IsConstraintError(err) {
			return event.NewCreateEventBadRequest().WithPayload(u.Err(400, err))
		}
		return event.NewCreateEventInternalServerError().WithPayload(u.Err(500, err))
	}
	return event.NewCreateEventCreated().WithPayload(res)
}

type readEvent struct {
	event *Service
}

func (e *readEvent) Handle(params event.ReadEventParams, principal *models.Principal) middleware.Responder {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return event.NewReadEventBadRequest().WithPayload(u.Err(400, err))
	}
	ctx := context.Background()
	res, err := e.event.ReadEventByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return event.NewReadEventNotFound().WithPayload(u.Err(404, err))
		}
		return event.NewReadEventInternalServerError().WithPayload(u.Err(500, err))
	}
	return event.NewReadEventOK().WithPayload(res)
}

type updateEvent struct {
	event *Service
}

func (e *updateEvent) Handle(params event.UpdateEventParams, principal *models.Principal) middleware.Responder {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return event.NewUpdateEventBadRequest().WithPayload(u.Err(400, err))
	}
	ctx := context.Background()
	res, err := e.event.UpdateEventByID(ctx, id, params.Event)
	if err != nil {
		if ent.IsValidationError(err) || ent.IsConstraintError(err) {
			return event.NewUpdateEventBadRequest().WithPayload(u.Err(400, err))
		} else if ent.IsNotFound(err) {
			return event.NewUpdateEventNotFound().WithPayload(u.Err(404, err))
		}
		return event.NewUpdateEventInternalServerError().WithPayload(u.Err(500, err))
	}
	return event.NewUpdateEventOK().WithPayload(res)
}

type deleteEvent struct {
	event *Service
}

func (e *deleteEvent) Handle(params event.DeleteEventParams, principal *models.Principal) middleware.Responder {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return event.NewDeleteEventBadRequest().WithPayload(u.Err(400, err))
	}
	ctx := context.Background()
	err = e.event.DeleteEventByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return event.NewDeleteEventNotFound().WithPayload(u.Err(404, err))
		}
		return event.NewDeleteEventInternalServerError().WithPayload(u.Err(500, err))
	}
	return event.NewDeleteEventNoContent()
}

type listEventUsers struct {
	event *Service
}

func (e *listEventUsers) Handle(params event.ListEventUsersParams, principal *models.Principal) middleware.Responder {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return event.NewListEventUsersBadRequest().WithPayload(u.Err(400, err))
	}
	ctx := context.Background()
	res, err := e.event.ListEventUsersByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return event.NewListEventUsersNotFound().WithPayload(u.Err(404, err))
		}
		return event.NewListEventUsersInternalServerError().WithPayload(u.Err(500, err))
	}
	return event.NewListEventUsersOK().WithPayload(res)
}

type getEventType struct {
	event *Service
}

func (e *getEventType) Handle(params event.GetEventTypeParams, principal *models.Principal) middleware.Responder {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return event.NewGetEventTypeBadRequest().WithPayload(u.Err(400, err))
	}
	ctx := context.Background()
	res, err := e.event.ReadEventTypeByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return event.NewGetEventTypeNotFound().WithPayload(u.Err(404, err))
		}
		return event.NewGetEventTypeInternalServerError().WithPayload(u.Err(500, err))
	}
	return event.NewGetEventTypeOK().WithPayload(res)
}

type setEventType struct {
	event *Service
}

func (e *setEventType) Handle(params event.SetEventTypeParams, principal *models.Principal) middleware.Responder {
	eventID, err := uuid.Parse(params.EventID)
	if err != nil {
		return event.NewSetEventTypeBadRequest().WithPayload(u.Err(400, err))
	}
	typeID, err := uuid.Parse(params.TypeID)
	if err != nil {
		return event.NewSetEventTypeBadRequest().WithPayload(u.Err(400, err))
	}
	ctx := context.Background()
	res, err := e.event.UpdateEventTypeByID(ctx, eventID, typeID)
	if err != nil {
		if ent.IsValidationError(err) || ent.IsConstraintError(err) {
			return event.NewSetEventTypeBadRequest().WithPayload(u.Err(400, err))
		} else if ent.IsNotFound(err) {
			return event.NewSetEventTypeNotFound().WithPayload(u.Err(404, err))
		}
		return event.NewSetEventTypeInternalServerError().WithPayload(u.Err(500, err))
	}
	return event.NewSetEventTypeOK().WithPayload(res)
}
