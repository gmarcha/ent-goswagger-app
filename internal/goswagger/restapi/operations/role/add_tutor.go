// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

// AddTutorHandlerFunc turns a function with the right signature into a add tutor handler
type AddTutorHandlerFunc func(AddTutorParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn AddTutorHandlerFunc) Handle(params AddTutorParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// AddTutorHandler interface for that can handle valid add tutor params
type AddTutorHandler interface {
	Handle(AddTutorParams, *models.Principal) middleware.Responder
}

// NewAddTutor creates a new http.Handler for the add tutor operation
func NewAddTutor(ctx *middleware.Context, handler AddTutorHandler) *AddTutor {
	return &AddTutor{Context: ctx, Handler: handler}
}

/* AddTutor swagger:route POST /users/{id}/role/tutor Role addTutor

Add tutor

Add tutor role to a user.

*/
type AddTutor struct {
	Context *middleware.Context
	Handler AddTutorHandler
}

func (o *AddTutor) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAddTutorParams()
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
