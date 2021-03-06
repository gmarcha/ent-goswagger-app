// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

// ListMeRolesHandlerFunc turns a function with the right signature into a list me roles handler
type ListMeRolesHandlerFunc func(ListMeRolesParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ListMeRolesHandlerFunc) Handle(params ListMeRolesParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ListMeRolesHandler interface for that can handle valid list me roles params
type ListMeRolesHandler interface {
	Handle(ListMeRolesParams, *models.Principal) middleware.Responder
}

// NewListMeRoles creates a new http.Handler for the list me roles operation
func NewListMeRoles(ctx *middleware.Context, handler ListMeRolesHandler) *ListMeRoles {
	return &ListMeRoles{Context: ctx, Handler: handler}
}

/* ListMeRoles swagger:route GET /users/me/roles User listMeRoles

List authenticated user roles

List the authenticated user's roles.

*/
type ListMeRoles struct {
	Context *middleware.Context
	Handler ListMeRolesHandler
}

func (o *ListMeRoles) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListMeRolesParams()
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
