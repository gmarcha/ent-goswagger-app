package user

import (
	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/restapi/operations"
)

func Init(api *operations.TutorAPI, db *ent.Client) {

	userService := &Service{User: db.User}

	api.UserListUserHandler = &listUser{user: userService}
	api.UserCreateUserHandler = &createUser{user: userService}
	api.UserReadUserHandler = &readUser{user: userService}
	api.UserUpdateUserHandler = &updateUser{user: userService}
	api.UserDeleteUserHandler = &deleteUser{user: userService}
	// api.UserListUserEventsHandler =
	// api.UserSubscribeUserHandler =
	// api.UserUnsubscribeUserHandler =

	// api.UserReadMeHandler =
	// api.UserUpdateMeHandler =
	// api.UserDeleteMeHandler =
	// api.UserListMeEventsHandler =
	// api.UserSubscribeMeHandler =
	// api.UserUnsubscribeMeHandler =
}
