// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewListEventParams creates a new ListEventParams object
//
// There are no default values defined in the spec.
func NewListEventParams() ListEventParams {

	return ListEventParams{}
}

// ListEventParams contains all the bound params for the list event operation
// typically these are obtained from a http.Request
//
// swagger:parameters listEvent
type ListEventParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*To end day.
	  In: query
	*/
	End *string
	/*From start day.
	  In: query
	*/
	Start *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListEventParams() beforehand.
func (o *ListEventParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qEnd, qhkEnd, _ := qs.GetOK("end")
	if err := o.bindEnd(qEnd, qhkEnd, route.Formats); err != nil {
		res = append(res, err)
	}

	qStart, qhkStart, _ := qs.GetOK("start")
	if err := o.bindStart(qStart, qhkStart, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindEnd binds and validates parameter End from query.
func (o *ListEventParams) bindEnd(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.End = &raw

	return nil
}

// bindStart binds and validates parameter Start from query.
func (o *ListEventParams) bindStart(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Start = &raw

	return nil
}
