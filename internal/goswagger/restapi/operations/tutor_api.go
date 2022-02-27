// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/restapi/operations/authentication"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/restapi/operations/event"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/restapi/operations/user"
)

// NewTutorAPI creates a new Tutor instance
func NewTutorAPI(spec *loads.Document) *TutorAPI {
	return &TutorAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		UserAddAdminHandler: user.AddAdminHandlerFunc(func(params user.AddAdminParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.AddAdmin has not yet been implemented")
		}),
		UserAddCalendarHandler: user.AddCalendarHandlerFunc(func(params user.AddCalendarParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.AddCalendar has not yet been implemented")
		}),
		UserAddTutorHandler: user.AddTutorHandlerFunc(func(params user.AddTutorParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.AddTutor has not yet been implemented")
		}),
		AuthenticationCallbackHandler: authentication.CallbackHandlerFunc(func(params authentication.CallbackParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication.Callback has not yet been implemented")
		}),
		EventCreateEventHandler: event.CreateEventHandlerFunc(func(params event.CreateEventParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation event.CreateEvent has not yet been implemented")
		}),
		UserCreateUserHandler: user.CreateUserHandlerFunc(func(params user.CreateUserParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.CreateUser has not yet been implemented")
		}),
		EventDeleteEventHandler: event.DeleteEventHandlerFunc(func(params event.DeleteEventParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation event.DeleteEvent has not yet been implemented")
		}),
		UserDeleteMeHandler: user.DeleteMeHandlerFunc(func(params user.DeleteMeParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteMe has not yet been implemented")
		}),
		UserDeleteUserHandler: user.DeleteUserHandlerFunc(func(params user.DeleteUserParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteUser has not yet been implemented")
		}),
		EventListEventHandler: event.ListEventHandlerFunc(func(params event.ListEventParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation event.ListEvent has not yet been implemented")
		}),
		EventListEventUsersHandler: event.ListEventUsersHandlerFunc(func(params event.ListEventUsersParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation event.ListEventUsers has not yet been implemented")
		}),
		UserListMeEventsHandler: user.ListMeEventsHandlerFunc(func(params user.ListMeEventsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.ListMeEvents has not yet been implemented")
		}),
		UserListUserHandler: user.ListUserHandlerFunc(func(params user.ListUserParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.ListUser has not yet been implemented")
		}),
		UserListUserEventsHandler: user.ListUserEventsHandlerFunc(func(params user.ListUserEventsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.ListUserEvents has not yet been implemented")
		}),
		AuthenticationLoginHandler: authentication.LoginHandlerFunc(func(params authentication.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication.Login has not yet been implemented")
		}),
		EventReadEventHandler: event.ReadEventHandlerFunc(func(params event.ReadEventParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation event.ReadEvent has not yet been implemented")
		}),
		UserReadMeHandler: user.ReadMeHandlerFunc(func(params user.ReadMeParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.ReadMe has not yet been implemented")
		}),
		UserReadUserHandler: user.ReadUserHandlerFunc(func(params user.ReadUserParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.ReadUser has not yet been implemented")
		}),
		UserRemoveAdminHandler: user.RemoveAdminHandlerFunc(func(params user.RemoveAdminParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.RemoveAdmin has not yet been implemented")
		}),
		UserRemoveCalendarHandler: user.RemoveCalendarHandlerFunc(func(params user.RemoveCalendarParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.RemoveCalendar has not yet been implemented")
		}),
		UserRemoveTutorHandler: user.RemoveTutorHandlerFunc(func(params user.RemoveTutorParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.RemoveTutor has not yet been implemented")
		}),
		UserSubscribeMeHandler: user.SubscribeMeHandlerFunc(func(params user.SubscribeMeParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.SubscribeMe has not yet been implemented")
		}),
		UserSubscribeUserHandler: user.SubscribeUserHandlerFunc(func(params user.SubscribeUserParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.SubscribeUser has not yet been implemented")
		}),
		AuthenticationTokenInfoHandler: authentication.TokenInfoHandlerFunc(func(params authentication.TokenInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication.TokenInfo has not yet been implemented")
		}),
		AuthenticationTokenRefreshHandler: authentication.TokenRefreshHandlerFunc(func(params authentication.TokenRefreshParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication.TokenRefresh has not yet been implemented")
		}),
		UserUnsubscribeMeHandler: user.UnsubscribeMeHandlerFunc(func(params user.UnsubscribeMeParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.UnsubscribeMe has not yet been implemented")
		}),
		UserUnsubscribeUserHandler: user.UnsubscribeUserHandlerFunc(func(params user.UnsubscribeUserParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.UnsubscribeUser has not yet been implemented")
		}),
		EventUpdateEventHandler: event.UpdateEventHandlerFunc(func(params event.UpdateEventParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation event.UpdateEvent has not yet been implemented")
		}),
		UserUpdateUserHandler: user.UpdateUserHandlerFunc(func(params user.UpdateUserParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.UpdateUser has not yet been implemented")
		}),

		OAuth2Auth: func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (OAuth2) has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*TutorAPI An API for 42 tutors. */
type TutorAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// OAuth2Auth registers a function that takes an access token and a collection of required scopes and returns a principal
	// it performs authentication based on an oauth2 bearer token provided in the request
	OAuth2Auth func(string, []string) (*models.Principal, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// UserAddAdminHandler sets the operation handler for the add admin operation
	UserAddAdminHandler user.AddAdminHandler
	// UserAddCalendarHandler sets the operation handler for the add calendar operation
	UserAddCalendarHandler user.AddCalendarHandler
	// UserAddTutorHandler sets the operation handler for the add tutor operation
	UserAddTutorHandler user.AddTutorHandler
	// AuthenticationCallbackHandler sets the operation handler for the callback operation
	AuthenticationCallbackHandler authentication.CallbackHandler
	// EventCreateEventHandler sets the operation handler for the create event operation
	EventCreateEventHandler event.CreateEventHandler
	// UserCreateUserHandler sets the operation handler for the create user operation
	UserCreateUserHandler user.CreateUserHandler
	// EventDeleteEventHandler sets the operation handler for the delete event operation
	EventDeleteEventHandler event.DeleteEventHandler
	// UserDeleteMeHandler sets the operation handler for the delete me operation
	UserDeleteMeHandler user.DeleteMeHandler
	// UserDeleteUserHandler sets the operation handler for the delete user operation
	UserDeleteUserHandler user.DeleteUserHandler
	// EventListEventHandler sets the operation handler for the list event operation
	EventListEventHandler event.ListEventHandler
	// EventListEventUsersHandler sets the operation handler for the list event users operation
	EventListEventUsersHandler event.ListEventUsersHandler
	// UserListMeEventsHandler sets the operation handler for the list me events operation
	UserListMeEventsHandler user.ListMeEventsHandler
	// UserListUserHandler sets the operation handler for the list user operation
	UserListUserHandler user.ListUserHandler
	// UserListUserEventsHandler sets the operation handler for the list user events operation
	UserListUserEventsHandler user.ListUserEventsHandler
	// AuthenticationLoginHandler sets the operation handler for the login operation
	AuthenticationLoginHandler authentication.LoginHandler
	// EventReadEventHandler sets the operation handler for the read event operation
	EventReadEventHandler event.ReadEventHandler
	// UserReadMeHandler sets the operation handler for the read me operation
	UserReadMeHandler user.ReadMeHandler
	// UserReadUserHandler sets the operation handler for the read user operation
	UserReadUserHandler user.ReadUserHandler
	// UserRemoveAdminHandler sets the operation handler for the remove admin operation
	UserRemoveAdminHandler user.RemoveAdminHandler
	// UserRemoveCalendarHandler sets the operation handler for the remove calendar operation
	UserRemoveCalendarHandler user.RemoveCalendarHandler
	// UserRemoveTutorHandler sets the operation handler for the remove tutor operation
	UserRemoveTutorHandler user.RemoveTutorHandler
	// UserSubscribeMeHandler sets the operation handler for the subscribe me operation
	UserSubscribeMeHandler user.SubscribeMeHandler
	// UserSubscribeUserHandler sets the operation handler for the subscribe user operation
	UserSubscribeUserHandler user.SubscribeUserHandler
	// AuthenticationTokenInfoHandler sets the operation handler for the token info operation
	AuthenticationTokenInfoHandler authentication.TokenInfoHandler
	// AuthenticationTokenRefreshHandler sets the operation handler for the token refresh operation
	AuthenticationTokenRefreshHandler authentication.TokenRefreshHandler
	// UserUnsubscribeMeHandler sets the operation handler for the unsubscribe me operation
	UserUnsubscribeMeHandler user.UnsubscribeMeHandler
	// UserUnsubscribeUserHandler sets the operation handler for the unsubscribe user operation
	UserUnsubscribeUserHandler user.UnsubscribeUserHandler
	// EventUpdateEventHandler sets the operation handler for the update event operation
	EventUpdateEventHandler event.UpdateEventHandler
	// UserUpdateUserHandler sets the operation handler for the update user operation
	UserUpdateUserHandler user.UpdateUserHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *TutorAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *TutorAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *TutorAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *TutorAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *TutorAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *TutorAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *TutorAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *TutorAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *TutorAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the TutorAPI
func (o *TutorAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.OAuth2Auth == nil {
		unregistered = append(unregistered, "OAuth2Auth")
	}

	if o.UserAddAdminHandler == nil {
		unregistered = append(unregistered, "user.AddAdminHandler")
	}
	if o.UserAddCalendarHandler == nil {
		unregistered = append(unregistered, "user.AddCalendarHandler")
	}
	if o.UserAddTutorHandler == nil {
		unregistered = append(unregistered, "user.AddTutorHandler")
	}
	if o.AuthenticationCallbackHandler == nil {
		unregistered = append(unregistered, "authentication.CallbackHandler")
	}
	if o.EventCreateEventHandler == nil {
		unregistered = append(unregistered, "event.CreateEventHandler")
	}
	if o.UserCreateUserHandler == nil {
		unregistered = append(unregistered, "user.CreateUserHandler")
	}
	if o.EventDeleteEventHandler == nil {
		unregistered = append(unregistered, "event.DeleteEventHandler")
	}
	if o.UserDeleteMeHandler == nil {
		unregistered = append(unregistered, "user.DeleteMeHandler")
	}
	if o.UserDeleteUserHandler == nil {
		unregistered = append(unregistered, "user.DeleteUserHandler")
	}
	if o.EventListEventHandler == nil {
		unregistered = append(unregistered, "event.ListEventHandler")
	}
	if o.EventListEventUsersHandler == nil {
		unregistered = append(unregistered, "event.ListEventUsersHandler")
	}
	if o.UserListMeEventsHandler == nil {
		unregistered = append(unregistered, "user.ListMeEventsHandler")
	}
	if o.UserListUserHandler == nil {
		unregistered = append(unregistered, "user.ListUserHandler")
	}
	if o.UserListUserEventsHandler == nil {
		unregistered = append(unregistered, "user.ListUserEventsHandler")
	}
	if o.AuthenticationLoginHandler == nil {
		unregistered = append(unregistered, "authentication.LoginHandler")
	}
	if o.EventReadEventHandler == nil {
		unregistered = append(unregistered, "event.ReadEventHandler")
	}
	if o.UserReadMeHandler == nil {
		unregistered = append(unregistered, "user.ReadMeHandler")
	}
	if o.UserReadUserHandler == nil {
		unregistered = append(unregistered, "user.ReadUserHandler")
	}
	if o.UserRemoveAdminHandler == nil {
		unregistered = append(unregistered, "user.RemoveAdminHandler")
	}
	if o.UserRemoveCalendarHandler == nil {
		unregistered = append(unregistered, "user.RemoveCalendarHandler")
	}
	if o.UserRemoveTutorHandler == nil {
		unregistered = append(unregistered, "user.RemoveTutorHandler")
	}
	if o.UserSubscribeMeHandler == nil {
		unregistered = append(unregistered, "user.SubscribeMeHandler")
	}
	if o.UserSubscribeUserHandler == nil {
		unregistered = append(unregistered, "user.SubscribeUserHandler")
	}
	if o.AuthenticationTokenInfoHandler == nil {
		unregistered = append(unregistered, "authentication.TokenInfoHandler")
	}
	if o.AuthenticationTokenRefreshHandler == nil {
		unregistered = append(unregistered, "authentication.TokenRefreshHandler")
	}
	if o.UserUnsubscribeMeHandler == nil {
		unregistered = append(unregistered, "user.UnsubscribeMeHandler")
	}
	if o.UserUnsubscribeUserHandler == nil {
		unregistered = append(unregistered, "user.UnsubscribeUserHandler")
	}
	if o.EventUpdateEventHandler == nil {
		unregistered = append(unregistered, "event.UpdateEventHandler")
	}
	if o.UserUpdateUserHandler == nil {
		unregistered = append(unregistered, "user.UpdateUserHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *TutorAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *TutorAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "OAuth2":
			result[name] = o.BearerAuthenticator(name, func(token string, scopes []string) (interface{}, error) {
				return o.OAuth2Auth(token, scopes)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *TutorAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *TutorAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *TutorAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *TutorAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the tutor API
func (o *TutorAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *TutorAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/{id}/role/admin"] = user.NewAddAdmin(o.context, o.UserAddAdminHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/{id}/role/calendar"] = user.NewAddCalendar(o.context, o.UserAddCalendarHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/{id}/role/tutor"] = user.NewAddTutor(o.context, o.UserAddTutorHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/auth/callback"] = authentication.NewCallback(o.context, o.AuthenticationCallbackHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/events"] = event.NewCreateEvent(o.context, o.EventCreateEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users"] = user.NewCreateUser(o.context, o.UserCreateUserHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/events/{id}"] = event.NewDeleteEvent(o.context, o.EventDeleteEventHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/me"] = user.NewDeleteMe(o.context, o.UserDeleteMeHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/{id}"] = user.NewDeleteUser(o.context, o.UserDeleteUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/events"] = event.NewListEvent(o.context, o.EventListEventHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/events/{id}/users"] = event.NewListEventUsers(o.context, o.EventListEventUsersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/me/events"] = user.NewListMeEvents(o.context, o.UserListMeEventsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users"] = user.NewListUser(o.context, o.UserListUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/{id}/events"] = user.NewListUserEvents(o.context, o.UserListUserEventsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/auth/login"] = authentication.NewLogin(o.context, o.AuthenticationLoginHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/events/{id}"] = event.NewReadEvent(o.context, o.EventReadEventHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/me"] = user.NewReadMe(o.context, o.UserReadMeHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/{id}"] = user.NewReadUser(o.context, o.UserReadUserHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/{id}/role/admin"] = user.NewRemoveAdmin(o.context, o.UserRemoveAdminHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/{id}/role/calendar"] = user.NewRemoveCalendar(o.context, o.UserRemoveCalendarHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/{id}/role/tutor"] = user.NewRemoveTutor(o.context, o.UserRemoveTutorHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/me/events/{id}"] = user.NewSubscribeMe(o.context, o.UserSubscribeMeHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/{userId}/events/{eventId}"] = user.NewSubscribeUser(o.context, o.UserSubscribeUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/auth/token/info"] = authentication.NewTokenInfo(o.context, o.AuthenticationTokenInfoHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/auth/token/refresh"] = authentication.NewTokenRefresh(o.context, o.AuthenticationTokenRefreshHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/me/events/{id}"] = user.NewUnsubscribeMe(o.context, o.UserUnsubscribeMeHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/{userId}/events/{eventId}"] = user.NewUnsubscribeUser(o.context, o.UserUnsubscribeUserHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/events/{id}"] = event.NewUpdateEvent(o.context, o.EventUpdateEventHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/users/{id}"] = user.NewUpdateUser(o.context, o.UserUpdateUserHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *TutorAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *TutorAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *TutorAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *TutorAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *TutorAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
