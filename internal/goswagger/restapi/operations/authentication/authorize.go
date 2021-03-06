// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AuthorizeHandlerFunc turns a function with the right signature into a authorize handler
type AuthorizeHandlerFunc func(AuthorizeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AuthorizeHandlerFunc) Handle(params AuthorizeParams) middleware.Responder {
	return fn(params)
}

// AuthorizeHandler interface for that can handle valid authorize params
type AuthorizeHandler interface {
	Handle(AuthorizeParams) middleware.Responder
}

// NewAuthorize creates a new http.Handler for the authorize operation
func NewAuthorize(ctx *middleware.Context, handler AuthorizeHandler) *Authorize {
	return &Authorize{Context: ctx, Handler: handler}
}

/* Authorize swagger:route GET /auth/authorize Authentication authorize

Handle authentication

Handle authorization API response, authenticate user and redirect to client.

*/
type Authorize struct {
	Context *middleware.Context
	Handler AuthorizeHandler
}

func (o *Authorize) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAuthorizeParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
