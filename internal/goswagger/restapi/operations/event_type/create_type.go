// Code generated by go-swagger; DO NOT EDIT.

package event_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

// CreateTypeHandlerFunc turns a function with the right signature into a create type handler
type CreateTypeHandlerFunc func(CreateTypeParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateTypeHandlerFunc) Handle(params CreateTypeParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// CreateTypeHandler interface for that can handle valid create type params
type CreateTypeHandler interface {
	Handle(CreateTypeParams, *models.Principal) middleware.Responder
}

// NewCreateType creates a new http.Handler for the create type operation
func NewCreateType(ctx *middleware.Context, handler CreateTypeHandler) *CreateType {
	return &CreateType{Context: ctx, Handler: handler}
}

/* CreateType swagger:route POST /events/types EventType createType

Create event type

Create a new event category.

*/
type CreateType struct {
	Context *middleware.Context
	Handler CreateTypeHandler
}

func (o *CreateType) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateTypeParams()
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
