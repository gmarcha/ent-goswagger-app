// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SubscribeUserHandlerFunc turns a function with the right signature into a subscribe user handler
type SubscribeUserHandlerFunc func(SubscribeUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SubscribeUserHandlerFunc) Handle(params SubscribeUserParams) middleware.Responder {
	return fn(params)
}

// SubscribeUserHandler interface for that can handle valid subscribe user params
type SubscribeUserHandler interface {
	Handle(SubscribeUserParams) middleware.Responder
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
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}