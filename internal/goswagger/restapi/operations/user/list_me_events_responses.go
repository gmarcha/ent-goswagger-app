// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

// ListMeEventsOKCode is the HTTP code returned for type ListMeEventsOK
const ListMeEventsOKCode int = 200

/*ListMeEventsOK OK

swagger:response listMeEventsOK
*/
type ListMeEventsOK struct {

	/*
	  In: Body
	*/
	Payload []*ent.Event `json:"body,omitempty"`
}

// NewListMeEventsOK creates ListMeEventsOK with default headers values
func NewListMeEventsOK() *ListMeEventsOK {

	return &ListMeEventsOK{}
}

// WithPayload adds the payload to the list me events o k response
func (o *ListMeEventsOK) WithPayload(payload []*ent.Event) *ListMeEventsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list me events o k response
func (o *ListMeEventsOK) SetPayload(payload []*ent.Event) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListMeEventsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*ent.Event, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListMeEventsUnauthorizedCode is the HTTP code returned for type ListMeEventsUnauthorized
const ListMeEventsUnauthorizedCode int = 401

/*ListMeEventsUnauthorized Unauthorized

swagger:response listMeEventsUnauthorized
*/
type ListMeEventsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListMeEventsUnauthorized creates ListMeEventsUnauthorized with default headers values
func NewListMeEventsUnauthorized() *ListMeEventsUnauthorized {

	return &ListMeEventsUnauthorized{}
}

// WithPayload adds the payload to the list me events unauthorized response
func (o *ListMeEventsUnauthorized) WithPayload(payload *models.Error) *ListMeEventsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list me events unauthorized response
func (o *ListMeEventsUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListMeEventsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListMeEventsForbiddenCode is the HTTP code returned for type ListMeEventsForbidden
const ListMeEventsForbiddenCode int = 403

/*ListMeEventsForbidden Forbidden

swagger:response listMeEventsForbidden
*/
type ListMeEventsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListMeEventsForbidden creates ListMeEventsForbidden with default headers values
func NewListMeEventsForbidden() *ListMeEventsForbidden {

	return &ListMeEventsForbidden{}
}

// WithPayload adds the payload to the list me events forbidden response
func (o *ListMeEventsForbidden) WithPayload(payload *models.Error) *ListMeEventsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list me events forbidden response
func (o *ListMeEventsForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListMeEventsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListMeEventsInternalServerErrorCode is the HTTP code returned for type ListMeEventsInternalServerError
const ListMeEventsInternalServerErrorCode int = 500

/*ListMeEventsInternalServerError Internal Server Error

swagger:response listMeEventsInternalServerError
*/
type ListMeEventsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListMeEventsInternalServerError creates ListMeEventsInternalServerError with default headers values
func NewListMeEventsInternalServerError() *ListMeEventsInternalServerError {

	return &ListMeEventsInternalServerError{}
}

// WithPayload adds the payload to the list me events internal server error response
func (o *ListMeEventsInternalServerError) WithPayload(payload *models.Error) *ListMeEventsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list me events internal server error response
func (o *ListMeEventsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListMeEventsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
