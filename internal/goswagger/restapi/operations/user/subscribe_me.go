// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

// SubscribeMeHandlerFunc turns a function with the right signature into a subscribe me handler
type SubscribeMeHandlerFunc func(SubscribeMeParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn SubscribeMeHandlerFunc) Handle(params SubscribeMeParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// SubscribeMeHandler interface for that can handle valid subscribe me params
type SubscribeMeHandler interface {
	Handle(SubscribeMeParams, *models.Principal) middleware.Responder
}

// NewSubscribeMe creates a new http.Handler for the subscribe me operation
func NewSubscribeMe(ctx *middleware.Context, handler SubscribeMeHandler) *SubscribeMe {
	return &SubscribeMe{Context: ctx, Handler: handler}
}

/* SubscribeMe swagger:route POST /users/me/events/{id} User subscribeMe

Subscribe authenticated user

Subscribe an authenticated user to an event.

*/
type SubscribeMe struct {
	Context *middleware.Context
	Handler SubscribeMeHandler
}

func (o *SubscribeMe) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSubscribeMeParams()
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
