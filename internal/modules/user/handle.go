package user

import (
	"context"

	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/restapi/operations/user"
	e "github.com/gmarcha/ent-goswagger-app/internal/utils"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
)

type listUser struct {
	user *Service
}

func (u *listUser) Handle(params user.ListUserParams, principal *models.Principal) middleware.Responder {
	ctx := context.Background()
	res, err := u.user.ListUser(ctx, params.Tutor)
	if err != nil {
		return user.NewListUserInternalServerError().WithPayload(e.Err(500, err))
	}
	return user.NewListUserOK().WithPayload(res)
}

type createUser struct {
	user *Service
}

func (u *createUser) Handle(params user.CreateUserParams, principal *models.Principal) middleware.Responder {
	ctx := context.Background()
	res, err := u.user.CreateUser(ctx, params.User)
	if err != nil {
		if ent.IsValidationError(err) || ent.IsConstraintError(err) {
			return user.NewCreateUserBadRequest().WithPayload(e.Err(400, err))
		}
		return user.NewCreateUserInternalServerError().WithPayload(e.Err(500, err))
	}
	return user.NewCreateUserCreated().WithPayload(res)
}

type readMe struct {
	user *Service
}

func (u *readMe) Handle(params user.ReadMeParams, principal *models.Principal) middleware.Responder {
	ctx := context.Background()
	res, err := u.user.ReadUserByLogin(ctx, string(*principal))
	if err != nil {
		return user.NewReadMeInternalServerError().WithPayload(e.Err(500, err))
	}
	return user.NewReadMeOK().WithPayload(res)
}

type updateMe struct {
	user *Service
}

func (u *updateMe) Handle(params user.UpdateMeParams, principal *models.Principal) middleware.Responder {
	ctx := context.Background()
	res, err := u.user.UpdateUserByLogin(ctx, string(*principal), params.User)
	if err != nil {
		if ent.IsValidationError(err) || ent.IsConstraintError(err) {
			return user.NewUpdateMeBadRequest().WithPayload(e.Err(400, err))
		}
		return user.NewUpdateMeInternalServerError().WithPayload(e.Err(500, err))
	}
	return user.NewReadMeOK().WithPayload(res)
}

type deleteMe struct {
	user *Service
}

func (u *deleteMe) Handle(params user.DeleteMeParams, principal *models.Principal) middleware.Responder {
	ctx := context.Background()
	err := u.user.DeleteUserByLogin(ctx, string(*principal))
	if err != nil {
		return user.NewDeleteMeInternalServerError().WithPayload(e.Err(500, err))
	}
	return user.NewDeleteMeNoContent()
}

type listMeEvents struct {
	user *Service
}

func (u *listMeEvents) Handle(params user.ListMeEventsParams, principal *models.Principal) middleware.Responder {
	ctx := context.Background()
	res, err := u.user.ListUserEventsByLogin(ctx, string(*principal))
	if err != nil {
		return user.NewListMeEventsInternalServerError().WithPayload(e.Err(500, err))
	}
	return user.NewListMeEventsOK().WithPayload(res)
}

type subscribeMe struct {
	user *Service
}

func (u *subscribeMe) Handle(params user.SubscribeMeParams, principal *models.Principal) middleware.Responder {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return user.NewReadUserBadRequest().WithPayload(e.Err(400, err))
	}
	ctx := context.Background()
	res, err := u.user.SubscribeUserByLogin(ctx, string(*principal), id)
	if err != nil {
		if ent.IsNotFound(err) {
			return user.NewSubscribeMeNotFound().WithPayload(e.Err(404, err))
		}
		return user.NewSubscribeMeInternalServerError().WithPayload(e.Err(500, err))
	}
	return user.NewSubscribeMeCreated().WithPayload(res)
}

type unsuscribeMe struct {
	user *Service
}

func (u *unsuscribeMe) Handle(params user.UnsubscribeMeParams, principal *models.Principal) middleware.Responder {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return user.NewReadUserBadRequest().WithPayload(e.Err(400, err))
	}
	ctx := context.Background()
	err = u.user.UnsubscribeUserByLogin(ctx, string(*principal), id)
	if err != nil {
		if ent.IsNotFound(err) {
			return user.NewUnsubscribeMeNotFound().WithPayload(e.Err(404, err))
		}
		return user.NewUnsubscribeMeInternalServerError().WithPayload(e.Err(500, err))
	}
	return user.NewUnsubscribeMeNoContent()
}

type readUser struct {
	user *Service
}

func (u *readUser) Handle(params user.ReadUserParams, principal *models.Principal) middleware.Responder {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return user.NewReadUserBadRequest().WithPayload(e.Err(400, err))
	}
	ctx := context.Background()
	res, err := u.user.ReadUserByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return user.NewReadUserNotFound().WithPayload(e.Err(404, err))
		}
		return user.NewReadUserInternalServerError().WithPayload(e.Err(500, err))
	}
	return user.NewReadUserOK().WithPayload(res)
}

type updateUser struct {
	user *Service
}

func (u *updateUser) Handle(params user.UpdateUserParams, principal *models.Principal) middleware.Responder {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return user.NewUpdateUserBadRequest().WithPayload(e.Err(400, err))
	}
	ctx := context.Background()
	res, err := u.user.UpdateUserByID(ctx, id, params.User)
	if err != nil {
		if ent.IsNotFound(err) {
			return user.NewUpdateUserNotFound().WithPayload(e.Err(404, err))
		} else if ent.IsValidationError(err) || ent.IsConstraintError(err) {
			return user.NewUpdateUserBadRequest().WithPayload(e.Err(400, err))
		}
		return user.NewUpdateUserInternalServerError().WithPayload(e.Err(500, err))
	}
	return user.NewUpdateUserOK().WithPayload(res)
}

type deleteUser struct {
	user *Service
}

func (u *deleteUser) Handle(params user.DeleteUserParams, principal *models.Principal) middleware.Responder {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return user.NewDeleteUserBadRequest().WithPayload(e.Err(400, err))
	}
	ctx := context.Background()
	err = u.user.DeleteUserByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return user.NewDeleteUserNotFound().WithPayload(e.Err(404, err))
		}
		return user.NewDeleteUserInternalServerError().WithPayload(e.Err(500, err))
	}
	return user.NewDeleteUserNoContent()
}

type listUserEvents struct {
	user *Service
}

func (u *listUserEvents) Handle(params user.ListUserEventsParams, principal *models.Principal) middleware.Responder {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return user.NewListUserEventsBadRequest().WithPayload(e.Err(400, err))
	}
	ctx := context.Background()
	res, err := u.user.ListUserEventsByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return user.NewListUserEventsNotFound().WithPayload(e.Err(404, err))
		}
		return user.NewListUserEventsInternalServerError().WithPayload(e.Err(500, err))
	}
	return user.NewListUserEventsOK().WithPayload(res)
}

type subscribeUser struct {
	user *Service
}

func (u *subscribeUser) Handle(params user.SubscribeUserParams, principal *models.Principal) middleware.Responder {
	userId, err := uuid.Parse(params.UserID)
	if err != nil {
		return user.NewSubscribeUserBadRequest().WithPayload(e.Err(400, err))
	}
	eventId, err := uuid.Parse(params.EventID)
	if err != nil {
		return user.NewSubscribeUserBadRequest().WithPayload(e.Err(400, err))
	}
	ctx := context.Background()
	res, err := u.user.SubscribeUserByID(ctx, userId, eventId)
	if err != nil {
		if ent.IsNotFound(err) {
			return user.NewSubscribeUserNotFound().WithPayload(e.Err(404, err))
		}
		return user.NewSubscribeUserInternalServerError().WithPayload(e.Err(500, err))
	}
	return user.NewSubscribeUserCreated().WithPayload(res)
}

type unsubscribeUser struct {
	user *Service
}

func (u *unsubscribeUser) Handle(params user.UnsubscribeUserParams, principal *models.Principal) middleware.Responder {
	userId, err := uuid.Parse(params.UserID)
	if err != nil {
		return user.NewUnsubscribeUserBadRequest().WithPayload(e.Err(400, err))
	}
	eventId, err := uuid.Parse(params.EventID)
	if err != nil {
		return user.NewUnsubscribeUserBadRequest().WithPayload(e.Err(400, err))
	}
	ctx := context.Background()
	err = u.user.UnsubscribeUserByID(ctx, userId, eventId)
	if err != nil {
		if ent.IsNotFound(err) {
			return user.NewUnsubscribeUserNotFound().WithPayload(e.Err(404, err))
		}
		return user.NewUnsubscribeUserInternalServerError().WithPayload(e.Err(500, err))
	}
	return user.NewUnsubscribeUserNoContent()
}
