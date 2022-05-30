package role

import (
	"context"

	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/restapi/operations/role"
	"github.com/gmarcha/ent-goswagger-app/internal/modules/user"
	e "github.com/gmarcha/ent-goswagger-app/internal/utils"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
)

type listRole struct {
	role *Service
}

func (r *listRole) Handle(params role.ListRoleParams, principal *models.Principal) middleware.Responder {
	ctx := context.Background()
	res, err := r.role.ListRole(ctx)
	if err != nil {
		return role.NewListRoleInternalServerError().WithPayload(e.Err(500, err))
	}
	return role.NewListRoleOK().WithPayload(res)
}

type addTutor struct {
	user *user.Service
	role *Service
}

func (r *addTutor) Handle(params role.AddTutorParams, principal *models.Principal) middleware.Responder {
	userId, err := uuid.Parse(params.ID)
	if err != nil {
		return role.NewAddTutorBadRequest().WithPayload(e.Err(400, err))
	}
	ctx := context.Background()
	res, err := r.role.AddRole(ctx, r.user, userId, "tutor")
	if err != nil {
		if ent.IsNotFound(err) {
			return role.NewAddTutorNotFound().WithPayload(e.Err(404, err))
		}
		return role.NewAddTutorInternalServerError().WithPayload(e.Err(500, err))
	}
	return role.NewAddTutorCreated().WithPayload(res)
}

type removeTutor struct {
	user *user.Service
	role *Service
}

func (r *removeTutor) Handle(params role.RemoveTutorParams, principal *models.Principal) middleware.Responder {
	userId, err := uuid.Parse(params.ID)
	if err != nil {
		return role.NewRemoveTutorBadRequest().WithPayload(e.Err(400, err))
	}
	ctx := context.Background()
	err = r.role.RemoveRole(ctx, r.user, userId, "tutor")
	if err != nil {
		if ent.IsNotFound(err) {
			return role.NewRemoveTutorNotFound().WithPayload(e.Err(404, err))
		}
		return role.NewRemoveTutorInternalServerError().WithPayload(e.Err(500, err))
	}
	return role.NewRemoveTutorNoContent()
}

type addCalendar struct {
	user *user.Service
	role *Service
}

func (r *addCalendar) Handle(params role.AddCalendarParams, principal *models.Principal) middleware.Responder {
	userId, err := uuid.Parse(params.ID)
	if err != nil {
		return role.NewAddCalendarBadRequest().WithPayload(e.Err(400, err))
	}
	ctx := context.Background()
	res, err := r.role.AddRole(ctx, r.user, userId, "calendar")
	if err != nil {
		if ent.IsNotFound(err) {
			return role.NewAddCalendarNotFound().WithPayload(e.Err(404, err))
		}
		return role.NewAddCalendarInternalServerError().WithPayload(e.Err(500, err))
	}
	return role.NewAddCalendarCreated().WithPayload(res)
}

type removeCalendar struct {
	user *user.Service
	role *Service
}

func (r *removeCalendar) Handle(params role.RemoveCalendarParams, principal *models.Principal) middleware.Responder {
	userId, err := uuid.Parse(params.ID)
	if err != nil {
		return role.NewRemoveCalendarBadRequest().WithPayload(e.Err(400, err))
	}
	ctx := context.Background()
	err = r.role.RemoveRole(ctx, r.user, userId, "calendar")
	if err != nil {
		if ent.IsNotFound(err) {
			return role.NewRemoveCalendarNotFound().WithPayload(e.Err(404, err))
		}
		return role.NewRemoveCalendarInternalServerError().WithPayload(e.Err(500, err))
	}
	return role.NewRemoveCalendarNoContent()
}

type addAdmin struct {
	user *user.Service
	role *Service
}

func (r *addAdmin) Handle(params role.AddAdminParams, principal *models.Principal) middleware.Responder {
	userId, err := uuid.Parse(params.ID)
	if err != nil {
		return role.NewAddAdminBadRequest().WithPayload(e.Err(400, err))
	}
	ctx := context.Background()
	res, err := r.role.AddRole(ctx, r.user, userId, "admin")
	if err != nil {
		if ent.IsNotFound(err) {
			return role.NewAddAdminNotFound().WithPayload(e.Err(404, err))
		}
		return role.NewAddAdminInternalServerError().WithPayload(e.Err(500, err))
	}
	return role.NewAddAdminCreated().WithPayload(res)
}

type removeAdmin struct {
	user *user.Service
	role *Service
}

func (r *removeAdmin) Handle(params role.RemoveAdminParams, principal *models.Principal) middleware.Responder {
	userId, err := uuid.Parse(params.ID)
	if err != nil {
		return role.NewRemoveAdminBadRequest().WithPayload(e.Err(400, err))
	}
	ctx := context.Background()
	err = r.role.RemoveRole(ctx, r.user, userId, "admin")
	if err != nil {
		if ent.IsNotFound(err) {
			return role.NewRemoveAdminNotFound().WithPayload(e.Err(404, err))
		}
		return role.NewRemoveAdminInternalServerError().WithPayload(e.Err(500, err))
	}
	return role.NewRemoveAdminNoContent()
}
