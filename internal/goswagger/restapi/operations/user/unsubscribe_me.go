// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

// UnsubscribeMeHandlerFunc turns a function with the right signature into a unsubscribe me handler
type UnsubscribeMeHandlerFunc func(UnsubscribeMeParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn UnsubscribeMeHandlerFunc) Handle(params UnsubscribeMeParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// UnsubscribeMeHandler interface for that can handle valid unsubscribe me params
type UnsubscribeMeHandler interface {
	Handle(UnsubscribeMeParams, *models.Principal) middleware.Responder
}

// NewUnsubscribeMe creates a new http.Handler for the unsubscribe me operation
func NewUnsubscribeMe(ctx *middleware.Context, handler UnsubscribeMeHandler) *UnsubscribeMe {
	return &UnsubscribeMe{Context: ctx, Handler: handler}
}

/* UnsubscribeMe swagger:route DELETE /users/me/events/{id} User unsubscribeMe

Unsubscribe authenticated user

Unsubscribe an authenticated user to an event.

*/
type UnsubscribeMe struct {
	Context *middleware.Context
	Handler UnsubscribeMeHandler
}

func (o *UnsubscribeMe) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUnsubscribeMeParams()
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
