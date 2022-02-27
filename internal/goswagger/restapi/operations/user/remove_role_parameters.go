// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewRemoveRoleParams creates a new RemoveRoleParams object
//
// There are no default values defined in the spec.
func NewRemoveRoleParams() RemoveRoleParams {

	return RemoveRoleParams{}
}

// RemoveRoleParams contains all the bound params for the remove role operation
// typically these are obtained from a http.Request
//
// swagger:parameters removeRole
type RemoveRoleParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*User ID.
	  Required: true
	  In: path
	*/
	ID string
	/*Role name.
	  Required: true
	  In: path
	*/
	Role string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRemoveRoleParams() beforehand.
func (o *RemoveRoleParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	rRole, rhkRole, _ := route.Params.GetOK("role")
	if err := o.bindRole(rRole, rhkRole, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *RemoveRoleParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ID = raw

	return nil
}

// bindRole binds and validates parameter Role from path.
func (o *RemoveRoleParams) bindRole(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Role = raw

	return nil
}
