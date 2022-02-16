// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/gamarcha/ent-goswagger-app/internal/goswagger/models"
)

// SubscribeUserHandlerFunc turns a function with the right signature into a subscribe user handler
type SubscribeUserHandlerFunc func(SubscribeUserParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn SubscribeUserHandlerFunc) Handle(params SubscribeUserParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// SubscribeUserHandler interface for that can handle valid subscribe user params
type SubscribeUserHandler interface {
	Handle(SubscribeUserParams, *models.Principal) middleware.Responder
}

// NewSubscribeUser creates a new http.Handler for the subscribe user operation
func NewSubscribeUser(ctx *middleware.Context, handler SubscribeUserHandler) *SubscribeUser {
	return &SubscribeUser{Context: ctx, Handler: handler}
}

/* SubscribeUser swagger:route POST /users/{userId}/events/{eventId} User subscribeUser

Subscribe user

Subscribe user to an event.

*/
type SubscribeUser struct {
	Context *middleware.Context
	Handler SubscribeUserHandler
}

func (o *SubscribeUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSubscribeUserParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}