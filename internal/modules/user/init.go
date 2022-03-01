package user

import (
	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/restapi/operations"
)

func Init(api *operations.TutorAPI, db *ent.Client) {

	userService := &Service{User: db.User}

	api.UserListUserHandler = &listUser{user: userService}
	api.UserCreateUserHandler = &createUser{user: userService}

	api.UserReadMeHandler = &readMe{user: userService}
	api.UserUpdateMeHandler = &updateMe{user: userService}
	api.UserDeleteMeHandler = &deleteMe{user: userService}
	api.UserListMeEventsHandler = &listMeEvents{user: userService}
	api.UserSubscribeMeHandler = &subscribeMe{user: userService}
	api.UserUnsubscribeMeHandler = &unsuscribeMe{user: userService}

	api.UserReadUserHandler = &readUser{user: userService}
	api.UserUpdateUserHandler = &updateUser{user: userService}
	api.UserDeleteUserHandler = &deleteUser{user: userService}
	api.UserListUserEventsHandler = &listUserEvents{user: userService}
	api.UserSubscribeUserHandler = &subscribeUser{user: userService}
	api.UserUnsubscribeUserHandler = &unsubscribeUser{user: userService}
}
