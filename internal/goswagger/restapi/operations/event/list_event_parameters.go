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

	/*Day filter.
	  In: query
	*/
	Day *string
	/*Month filter.
	  In: query
	*/
	Month *string
	/*Week filter.
	  In: query
	*/
	Week *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListEventParams() beforehand.
func (o *ListEventParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qDay, qhkDay, _ := qs.GetOK("day")
	if err := o.bindDay(qDay, qhkDay, route.Formats); err != nil {
		res = append(res, err)
	}

	qMonth, qhkMonth, _ := qs.GetOK("month")
	if err := o.bindMonth(qMonth, qhkMonth, route.Formats); err != nil {
		res = append(res, err)
	}

	qWeek, qhkWeek, _ := qs.GetOK("week")
	if err := o.bindWeek(qWeek, qhkWeek, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDay binds and validates parameter Day from query.
func (o *ListEventParams) bindDay(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Day = &raw

	return nil
}

// bindMonth binds and validates parameter Month from query.
func (o *ListEventParams) bindMonth(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Month = &raw

	return nil
}

// bindWeek binds and validates parameter Week from query.
func (o *ListEventParams) bindWeek(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Week = &raw

	return nil
}
