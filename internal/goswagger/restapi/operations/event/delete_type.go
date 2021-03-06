// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

// DeleteTypeHandlerFunc turns a function with the right signature into a delete type handler
type DeleteTypeHandlerFunc func(DeleteTypeParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteTypeHandlerFunc) Handle(params DeleteTypeParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DeleteTypeHandler interface for that can handle valid delete type params
type DeleteTypeHandler interface {
	Handle(DeleteTypeParams, *models.Principal) middleware.Responder
}

// NewDeleteType creates a new http.Handler for the delete type operation
func NewDeleteType(ctx *middleware.Context, handler DeleteTypeHandler) *DeleteType {
	return &DeleteType{Context: ctx, Handler: handler}
}

/* DeleteType swagger:route DELETE /events/types/{id} Event deleteType

Delete event type

Update an event category.

*/
type DeleteType struct {
	Context *middleware.Context
	Handler DeleteTypeHandler
}

func (o *DeleteType) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteTypeParams()
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
