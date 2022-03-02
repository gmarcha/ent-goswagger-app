// Code generated by go-swagger; DO NOT EDIT.

package event_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

// ListTypeEventsOKCode is the HTTP code returned for type ListTypeEventsOK
const ListTypeEventsOKCode int = 200

/*ListTypeEventsOK OK

swagger:response listTypeEventsOK
*/
type ListTypeEventsOK struct {

	/*
	  In: Body
	*/
	Payload []*ent.Event `json:"body,omitempty"`
}

// NewListTypeEventsOK creates ListTypeEventsOK with default headers values
func NewListTypeEventsOK() *ListTypeEventsOK {

	return &ListTypeEventsOK{}
}

// WithPayload adds the payload to the list type events o k response
func (o *ListTypeEventsOK) WithPayload(payload []*ent.Event) *ListTypeEventsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list type events o k response
func (o *ListTypeEventsOK) SetPayload(payload []*ent.Event) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListTypeEventsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ListTypeEventsBadRequestCode is the HTTP code returned for type ListTypeEventsBadRequest
const ListTypeEventsBadRequestCode int = 400

/*ListTypeEventsBadRequest Bad request

swagger:response listTypeEventsBadRequest
*/
type ListTypeEventsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListTypeEventsBadRequest creates ListTypeEventsBadRequest with default headers values
func NewListTypeEventsBadRequest() *ListTypeEventsBadRequest {

	return &ListTypeEventsBadRequest{}
}

// WithPayload adds the payload to the list type events bad request response
func (o *ListTypeEventsBadRequest) WithPayload(payload *models.Error) *ListTypeEventsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list type events bad request response
func (o *ListTypeEventsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListTypeEventsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListTypeEventsUnauthorizedCode is the HTTP code returned for type ListTypeEventsUnauthorized
const ListTypeEventsUnauthorizedCode int = 401

/*ListTypeEventsUnauthorized Unauthorized

swagger:response listTypeEventsUnauthorized
*/
type ListTypeEventsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListTypeEventsUnauthorized creates ListTypeEventsUnauthorized with default headers values
func NewListTypeEventsUnauthorized() *ListTypeEventsUnauthorized {

	return &ListTypeEventsUnauthorized{}
}

// WithPayload adds the payload to the list type events unauthorized response
func (o *ListTypeEventsUnauthorized) WithPayload(payload *models.Error) *ListTypeEventsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list type events unauthorized response
func (o *ListTypeEventsUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListTypeEventsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListTypeEventsForbiddenCode is the HTTP code returned for type ListTypeEventsForbidden
const ListTypeEventsForbiddenCode int = 403

/*ListTypeEventsForbidden Forbidden

swagger:response listTypeEventsForbidden
*/
type ListTypeEventsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListTypeEventsForbidden creates ListTypeEventsForbidden with default headers values
func NewListTypeEventsForbidden() *ListTypeEventsForbidden {

	return &ListTypeEventsForbidden{}
}

// WithPayload adds the payload to the list type events forbidden response
func (o *ListTypeEventsForbidden) WithPayload(payload *models.Error) *ListTypeEventsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list type events forbidden response
func (o *ListTypeEventsForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListTypeEventsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListTypeEventsNotFoundCode is the HTTP code returned for type ListTypeEventsNotFound
const ListTypeEventsNotFoundCode int = 404

/*ListTypeEventsNotFound Not Found

swagger:response listTypeEventsNotFound
*/
type ListTypeEventsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListTypeEventsNotFound creates ListTypeEventsNotFound with default headers values
func NewListTypeEventsNotFound() *ListTypeEventsNotFound {

	return &ListTypeEventsNotFound{}
}

// WithPayload adds the payload to the list type events not found response
func (o *ListTypeEventsNotFound) WithPayload(payload *models.Error) *ListTypeEventsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list type events not found response
func (o *ListTypeEventsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListTypeEventsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListTypeEventsInternalServerErrorCode is the HTTP code returned for type ListTypeEventsInternalServerError
const ListTypeEventsInternalServerErrorCode int = 500

/*ListTypeEventsInternalServerError Internal Server Error

swagger:response listTypeEventsInternalServerError
*/
type ListTypeEventsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListTypeEventsInternalServerError creates ListTypeEventsInternalServerError with default headers values
func NewListTypeEventsInternalServerError() *ListTypeEventsInternalServerError {

	return &ListTypeEventsInternalServerError{}
}

// WithPayload adds the payload to the list type events internal server error response
func (o *ListTypeEventsInternalServerError) WithPayload(payload *models.Error) *ListTypeEventsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list type events internal server error response
func (o *ListTypeEventsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListTypeEventsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}