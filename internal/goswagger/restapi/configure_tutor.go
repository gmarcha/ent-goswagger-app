// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/gamarcha/ent-goswagger-app/internal/clients/database"
	"github.com/gamarcha/ent-goswagger-app/internal/goswagger/models"
	"github.com/gamarcha/ent-goswagger-app/internal/goswagger/restapi/operations"
	"github.com/gamarcha/ent-goswagger-app/internal/goswagger/restapi/operations/authentication"
	"github.com/gamarcha/ent-goswagger-app/internal/goswagger/restapi/operations/event"
	"github.com/gamarcha/ent-goswagger-app/internal/goswagger/restapi/operations/user"
)

//go:generate swagger generate server --target ../../goswagger --name Tutor --spec ../../../docs/swagger.yaml --principal models.Principal --exclude-main

func configureFlags(api *operations.TutorAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TutorAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	db := database.Init()

	if api.OauthSecurityAuth == nil {
		api.OauthSecurityAuth = func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (OauthSecurity) has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.AuthenticationGetAuthCallbackHandler == nil {
		api.AuthenticationGetAuthCallbackHandler = authentication.GetAuthCallbackHandlerFunc(func(params authentication.GetAuthCallbackParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication.GetAuthCallback has not yet been implemented")
		})
	}
	if api.AuthenticationGetLoginHandler == nil {
		api.AuthenticationGetLoginHandler = authentication.GetLoginHandlerFunc(func(params authentication.GetLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication.GetLogin has not yet been implemented")
		})
	}
	if api.EventCreateEventHandler == nil {
		api.EventCreateEventHandler = event.CreateEventHandlerFunc(func(params event.CreateEventParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation event.CreateEvent has not yet been implemented")
		})
	}
	if api.UserCreateUserHandler == nil {
		api.UserCreateUserHandler = user.CreateUserHandlerFunc(func(params user.CreateUserParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.CreateUser has not yet been implemented")
		})
	}
	if api.EventDeleteEventHandler == nil {
		api.EventDeleteEventHandler = event.DeleteEventHandlerFunc(func(params event.DeleteEventParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation event.DeleteEvent has not yet been implemented")
		})
	}
	if api.UserDeleteUserHandler == nil {
		api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(func(params user.DeleteUserParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteUser has not yet been implemented")
		})
	}
	if api.EventListEventHandler == nil {
		api.EventListEventHandler = event.ListEventHandlerFunc(func(params event.ListEventParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation event.ListEvent has not yet been implemented")
		})
	}
	if api.EventListEventUsersHandler == nil {
		api.EventListEventUsersHandler = event.ListEventUsersHandlerFunc(func(params event.ListEventUsersParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation event.ListEventUsers has not yet been implemented")
		})
	}
	if api.UserListUserHandler == nil {
		api.UserListUserHandler = user.ListUserHandlerFunc(func(params user.ListUserParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.ListUser has not yet been implemented")
		})
	}
	if api.UserListUserEventsHandler == nil {
		api.UserListUserEventsHandler = user.ListUserEventsHandlerFunc(func(params user.ListUserEventsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.ListUserEvents has not yet been implemented")
		})
	}
	if api.EventReadEventHandler == nil {
		api.EventReadEventHandler = event.ReadEventHandlerFunc(func(params event.ReadEventParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation event.ReadEvent has not yet been implemented")
		})
	}
	if api.UserReadUserHandler == nil {
		api.UserReadUserHandler = user.ReadUserHandlerFunc(func(params user.ReadUserParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.ReadUser has not yet been implemented")
		})
	}
	if api.UserSubscribeUserHandler == nil {
		api.UserSubscribeUserHandler = user.SubscribeUserHandlerFunc(func(params user.SubscribeUserParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.SubscribeUser has not yet been implemented")
		})
	}
	if api.UserUnsubscribeUserHandler == nil {
		api.UserUnsubscribeUserHandler = user.UnsubscribeUserHandlerFunc(func(params user.UnsubscribeUserParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.UnsubscribeUser has not yet been implemented")
		})
	}
	if api.EventUpdateEventHandler == nil {
		api.EventUpdateEventHandler = event.UpdateEventHandlerFunc(func(params event.UpdateEventParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation event.UpdateEvent has not yet been implemented")
		})
	}
	if api.UserUpdateUserHandler == nil {
		api.UserUpdateUserHandler = user.UpdateUserHandlerFunc(func(params user.UpdateUserParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.UpdateUser has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
