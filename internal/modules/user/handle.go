package user

import (
	"context"

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
		return user.NewCreateUserInternalServerError().WithPayload(e.Err(500, err))
	}
	return user.NewCreateUserCreated().WithPayload(res)
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
	res, err := u.user.ReadUser(ctx, id)
	if err != nil {
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
	res, err := u.user.UpdateUser(ctx, id, params.User)
	if err != nil {
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
	err = u.user.DeleteUser(ctx, id)
	if err != nil {
		return user.NewDeleteUserInternalServerError().WithPayload(e.Err(500, err))
	}
	return user.NewDeleteUserNoContent()
}

type readMe struct {
	user *Service
}

func (u *readMe) Handle(params user.ReadMeParams, principal *models.Principal) middleware.Responder {
	ctx := context.Background()
	res, err := u.user.ReadMe(ctx, string(*principal))
	if err != nil {
		return user.NewReadMeInternalServerError().WithPayload(e.Err(500, err))
	}
	return user.NewReadMeOK().WithPayload(res)
}
