// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/gamarcha/ent-goswagger-app/internal/goswagger/models"
)

// ListEventUsersHandlerFunc turns a function with the right signature into a list event users handler
type ListEventUsersHandlerFunc func(ListEventUsersParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ListEventUsersHandlerFunc) Handle(params ListEventUsersParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ListEventUsersHandler interface for that can handle valid list event users params
type ListEventUsersHandler interface {
	Handle(ListEventUsersParams, *models.Principal) middleware.Responder
}

// NewListEventUsers creates a new http.Handler for the list event users operation
func NewListEventUsers(ctx *middleware.Context, handler ListEventUsersHandler) *ListEventUsers {
	return &ListEventUsers{Context: ctx, Handler: handler}
}

/* ListEventUsers swagger:route GET /events/{id}/users Event listEventUsers

List event users

List users subscribed to an event.

*/
type ListEventUsers struct {
	Context *middleware.Context
	Handler ListEventUsersHandler
}

func (o *ListEventUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListEventUsersParams()
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