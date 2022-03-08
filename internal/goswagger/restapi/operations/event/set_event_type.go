// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

// SetEventTypeHandlerFunc turns a function with the right signature into a set event type handler
type SetEventTypeHandlerFunc func(SetEventTypeParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn SetEventTypeHandlerFunc) Handle(params SetEventTypeParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// SetEventTypeHandler interface for that can handle valid set event type params
type SetEventTypeHandler interface {
	Handle(SetEventTypeParams, *models.Principal) middleware.Responder
}

// NewSetEventType creates a new http.Handler for the set event type operation
func NewSetEventType(ctx *middleware.Context, handler SetEventTypeHandler) *SetEventType {
	return &SetEventType{Context: ctx, Handler: handler}
}

/* SetEventType swagger:route PATCH /events/{eventID}/types/{typeID} Event setEventType

Set event type

Set event type value.

*/
type SetEventType struct {
	Context *middleware.Context
	Handler SetEventTypeHandler
}

func (o *SetEventType) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSetEventTypeParams()
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
