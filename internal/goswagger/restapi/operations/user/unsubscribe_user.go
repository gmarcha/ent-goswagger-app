// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

// UnsubscribeUserHandlerFunc turns a function with the right signature into a unsubscribe user handler
type UnsubscribeUserHandlerFunc func(UnsubscribeUserParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn UnsubscribeUserHandlerFunc) Handle(params UnsubscribeUserParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// UnsubscribeUserHandler interface for that can handle valid unsubscribe user params
type UnsubscribeUserHandler interface {
	Handle(UnsubscribeUserParams, *models.Principal) middleware.Responder
}

// NewUnsubscribeUser creates a new http.Handler for the unsubscribe user operation
func NewUnsubscribeUser(ctx *middleware.Context, handler UnsubscribeUserHandler) *UnsubscribeUser {
	return &UnsubscribeUser{Context: ctx, Handler: handler}
}

/* UnsubscribeUser swagger:route DELETE /users/{userId}/events/{eventId} User unsubscribeUser

Unsubscribe user

Unsubscribe a user to an event.

*/
type UnsubscribeUser struct {
	Context *middleware.Context
	Handler UnsubscribeUserHandler
}

func (o *UnsubscribeUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUnsubscribeUserParams()
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
