package eventType

import (
	"context"

	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/restapi/operations/event"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/restapi/operations/event_type"
	u "github.com/gmarcha/ent-goswagger-app/internal/utils"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
)

type listType struct {
	eventType *Service
}

func (t *listType) Handle(params event_type.ListTypeParams, principal *models.Principal) middleware.Responder {
	ctx := context.Background()
	res, err := t.eventType.ListType(ctx)
	if err != nil {
		event.NewListTypeInternalServerError().WithPayload(u.Err(500, err))
	}
	return event.NewListTypeOK().WithPayload(res)
}

type createType struct {
	eventType *Service
}

func (t *createType) Handle(params event_type.CreateTypeParams, principal *models.Principal) middleware.Responder {
	ctx := context.Background()
	res, err := t.eventType.CreateType(ctx, params.Type)
	if err != nil {
		if ent.IsValidationError(err) || ent.IsConstraintError(err) {
			return event.NewCreateTypeBadRequest().WithPayload(u.Err(400, err))
		}
		return event.NewCreateTypeInternalServerError().WithPayload(u.Err(500, err))
	}
	return event.NewCreateTypeCreated().WithPayload(res)
}

type readType struct {
	eventType *Service
}

func (t *readType) Handle(params event_type.ReadTypeParams, principal *models.Principal) middleware.Responder {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return event.NewReadTypeBadRequest().WithPayload(u.Err(400, err))
	}
	ctx := context.Background()
	res, err := t.eventType.ReadTypeByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return event.NewReadTypeNotFound().WithPayload(u.Err(404, err))
		}
		return event.NewReadTypeInternalServerError().WithPayload(u.Err(500, err))
	}
	return event.NewReadTypeOK().WithPayload(res)
}

type updateType struct {
	eventType *Service
}

func (t *updateType) Handle(params event_type.UpdateTypeParams, principal *models.Principal) middleware.Responder {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return event.NewUpdateTypeBadRequest().WithPayload(u.Err(400, err))
	}
	ctx := context.Background()
	res, err := t.eventType.UpdateTypeByID(ctx, id, params.Type)
	if err != nil {
		if ent.IsValidationError(err) || ent.IsConstraintError(err) {
			return event.NewUpdateTypeBadRequest().WithPayload(u.Err(400, err))
		} else if ent.IsNotFound(err) {
			return event.NewUpdateTypeNotFound().WithPayload(u.Err(404, err))
		}
		return event.NewUpdateTypeInternalServerError().WithPayload(u.Err(500, err))
	}
	return event.NewUpdateTypeOK().WithPayload(res)
}

type deleteType struct {
	eventType *Service
}

func (t *deleteType) Handle(params event_type.DeleteTypeParams, principal *models.Principal) middleware.Responder {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return event.NewDeleteTypeBadRequest().WithPayload(u.Err(400, err))
	}
	ctx := context.Background()
	err = t.eventType.DeleteTypeByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return event.NewDeleteTypeNotFound().WithPayload(u.Err(404, err))
		}
		return event.NewDeleteTypeInternalServerError().WithPayload(u.Err(500, err))
	}
	return event.NewDeleteTypeNoContent()
}

type listTypeEvents struct {
	eventType *Service
}

func (t *listTypeEvents) Handle(params event_type.ListTypeEventsParams, principal *models.Principal) middleware.Responder {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return event_type.NewListTypeEventsBadRequest().WithPayload(u.Err(400, err))
	}
	ctx := context.Background()
	res, err := t.eventType.ListTypeEventsByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return event_type.NewListTypeEventsNotFound().WithPayload(u.Err(404, err))
		}
		return event_type.NewListTypeEventsInternalServerError().WithPayload(u.Err(500, err))
	}
	return event_type.NewListTypeEventsOK().WithPayload(res)
}
