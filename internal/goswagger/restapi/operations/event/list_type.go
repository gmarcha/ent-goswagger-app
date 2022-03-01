// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

// ListTypeHandlerFunc turns a function with the right signature into a list type handler
type ListTypeHandlerFunc func(ListTypeParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ListTypeHandlerFunc) Handle(params ListTypeParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ListTypeHandler interface for that can handle valid list type params
type ListTypeHandler interface {
	Handle(ListTypeParams, *models.Principal) middleware.Responder
}

// NewListType creates a new http.Handler for the list type operation
func NewListType(ctx *middleware.Context, handler ListTypeHandler) *ListType {
	return &ListType{Context: ctx, Handler: handler}
}

/* ListType swagger:route GET /events/types Event listType

List event types

List all event categories.

*/
type ListType struct {
	Context *middleware.Context
	Handler ListTypeHandler
}

func (o *ListType) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListTypeParams()
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