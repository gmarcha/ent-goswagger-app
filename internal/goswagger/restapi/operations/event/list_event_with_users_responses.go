// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

// ListEventWithUsersOKCode is the HTTP code returned for type ListEventWithUsersOK
const ListEventWithUsersOKCode int = 200

/*ListEventWithUsersOK OK

swagger:response listEventWithUsersOK
*/
type ListEventWithUsersOK struct {

	/*
	  In: Body
	*/
	Payload []*ent.Event `json:"body,omitempty"`
}

// NewListEventWithUsersOK creates ListEventWithUsersOK with default headers values
func NewListEventWithUsersOK() *ListEventWithUsersOK {

	return &ListEventWithUsersOK{}
}

// WithPayload adds the payload to the list event with users o k response
func (o *ListEventWithUsersOK) WithPayload(payload []*ent.Event) *ListEventWithUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list event with users o k response
func (o *ListEventWithUsersOK) SetPayload(payload []*ent.Event) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEventWithUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ListEventWithUsersUnauthorizedCode is the HTTP code returned for type ListEventWithUsersUnauthorized
const ListEventWithUsersUnauthorizedCode int = 401

/*ListEventWithUsersUnauthorized Unauthorized

swagger:response listEventWithUsersUnauthorized
*/
type ListEventWithUsersUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListEventWithUsersUnauthorized creates ListEventWithUsersUnauthorized with default headers values
func NewListEventWithUsersUnauthorized() *ListEventWithUsersUnauthorized {

	return &ListEventWithUsersUnauthorized{}
}

// WithPayload adds the payload to the list event with users unauthorized response
func (o *ListEventWithUsersUnauthorized) WithPayload(payload *models.Error) *ListEventWithUsersUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list event with users unauthorized response
func (o *ListEventWithUsersUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEventWithUsersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListEventWithUsersForbiddenCode is the HTTP code returned for type ListEventWithUsersForbidden
const ListEventWithUsersForbiddenCode int = 403

/*ListEventWithUsersForbidden Forbidden

swagger:response listEventWithUsersForbidden
*/
type ListEventWithUsersForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListEventWithUsersForbidden creates ListEventWithUsersForbidden with default headers values
func NewListEventWithUsersForbidden() *ListEventWithUsersForbidden {

	return &ListEventWithUsersForbidden{}
}

// WithPayload adds the payload to the list event with users forbidden response
func (o *ListEventWithUsersForbidden) WithPayload(payload *models.Error) *ListEventWithUsersForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list event with users forbidden response
func (o *ListEventWithUsersForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEventWithUsersForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListEventWithUsersInternalServerErrorCode is the HTTP code returned for type ListEventWithUsersInternalServerError
const ListEventWithUsersInternalServerErrorCode int = 500

/*ListEventWithUsersInternalServerError Internal Server Error

swagger:response listEventWithUsersInternalServerError
*/
type ListEventWithUsersInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListEventWithUsersInternalServerError creates ListEventWithUsersInternalServerError with default headers values
func NewListEventWithUsersInternalServerError() *ListEventWithUsersInternalServerError {

	return &ListEventWithUsersInternalServerError{}
}

// WithPayload adds the payload to the list event with users internal server error response
func (o *ListEventWithUsersInternalServerError) WithPayload(payload *models.Error) *ListEventWithUsersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list event with users internal server error response
func (o *ListEventWithUsersInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEventWithUsersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
