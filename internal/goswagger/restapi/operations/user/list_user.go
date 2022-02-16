// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/gamarcha/ent-goswagger-app/internal/goswagger/models"
)

// ListUserHandlerFunc turns a function with the right signature into a list user handler
type ListUserHandlerFunc func(ListUserParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ListUserHandlerFunc) Handle(params ListUserParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ListUserHandler interface for that can handle valid list user params
type ListUserHandler interface {
	Handle(ListUserParams, *models.Principal) middleware.Responder
}

// NewListUser creates a new http.Handler for the list user operation
func NewListUser(ctx *middleware.Context, handler ListUserHandler) *ListUser {
	return &ListUser{Context: ctx, Handler: handler}
}

/* ListUser swagger:route GET /users User listUser

List users

List all users.

*/
type ListUser struct {
	Context *middleware.Context
	Handler ListUserHandler
}

func (o *ListUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListUserParams()
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