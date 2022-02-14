// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListEventHandlerFunc turns a function with the right signature into a list event handler
type ListEventHandlerFunc func(ListEventParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListEventHandlerFunc) Handle(params ListEventParams) middleware.Responder {
	return fn(params)
}

// ListEventHandler interface for that can handle valid list event params
type ListEventHandler interface {
	Handle(ListEventParams) middleware.Responder
}

// NewListEvent creates a new http.Handler for the list event operation
func NewListEvent(ctx *middleware.Context, handler ListEventHandler) *ListEvent {
	return &ListEvent{Context: ctx, Handler: handler}
}

/* ListEvent swagger:route GET /events Event listEvent

List events

List all events.

*/
type ListEvent struct {
	Context *middleware.Context
	Handler ListEventHandler
}

func (o *ListEvent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListEventParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
