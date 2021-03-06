// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

// ListUserRolesHandlerFunc turns a function with the right signature into a list user roles handler
type ListUserRolesHandlerFunc func(ListUserRolesParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ListUserRolesHandlerFunc) Handle(params ListUserRolesParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ListUserRolesHandler interface for that can handle valid list user roles params
type ListUserRolesHandler interface {
	Handle(ListUserRolesParams, *models.Principal) middleware.Responder
}

// NewListUserRoles creates a new http.Handler for the list user roles operation
func NewListUserRoles(ctx *middleware.Context, handler ListUserRolesHandler) *ListUserRoles {
	return &ListUserRoles{Context: ctx, Handler: handler}
}

/* ListUserRoles swagger:route GET /users/{id}/roles User listUserRoles

List user roles

List the user's roles.

*/
type ListUserRoles struct {
	Context *middleware.Context
	Handler ListUserRolesHandler
}

func (o *ListUserRoles) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListUserRolesParams()
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
