// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateEventHandlerFunc turns a function with the right signature into a update event handler
type UpdateEventHandlerFunc func(UpdateEventParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateEventHandlerFunc) Handle(params UpdateEventParams) middleware.Responder {
	return fn(params)
}

// UpdateEventHandler interface for that can handle valid update event params
type UpdateEventHandler interface {
	Handle(UpdateEventParams) middleware.Responder
}

// NewUpdateEvent creates a new http.Handler for the update event operation
func NewUpdateEvent(ctx *middleware.Context, handler UpdateEventHandler) *UpdateEvent {
	return &UpdateEvent{Context: ctx, Handler: handler}
}

/* UpdateEvent swagger:route PUT /events/{id} Event updateEvent

Update event

Update an event by ID.

*/
type UpdateEvent struct {
	Context *middleware.Context
	Handler UpdateEventHandler
}

func (o *UpdateEvent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateEventParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
