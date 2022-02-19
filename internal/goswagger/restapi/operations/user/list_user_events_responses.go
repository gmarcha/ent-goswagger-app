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

// ListUserEventsOKCode is the HTTP code returned for type ListUserEventsOK
const ListUserEventsOKCode int = 200

/*ListUserEventsOK OK

swagger:response listUserEventsOK
*/
type ListUserEventsOK struct {

	/*
	  In: Body
	*/
	Payload []*ent.Event `json:"body,omitempty"`
}

// NewListUserEventsOK creates ListUserEventsOK with default headers values
func NewListUserEventsOK() *ListUserEventsOK {

	return &ListUserEventsOK{}
}

// WithPayload adds the payload to the list user events o k response
func (o *ListUserEventsOK) WithPayload(payload []*ent.Event) *ListUserEventsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list user events o k response
func (o *ListUserEventsOK) SetPayload(payload []*ent.Event) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUserEventsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ListUserEventsBadRequestCode is the HTTP code returned for type ListUserEventsBadRequest
const ListUserEventsBadRequestCode int = 400

/*ListUserEventsBadRequest Bad request

swagger:response listUserEventsBadRequest
*/
type ListUserEventsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListUserEventsBadRequest creates ListUserEventsBadRequest with default headers values
func NewListUserEventsBadRequest() *ListUserEventsBadRequest {

	return &ListUserEventsBadRequest{}
}

// WithPayload adds the payload to the list user events bad request response
func (o *ListUserEventsBadRequest) WithPayload(payload *models.Error) *ListUserEventsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list user events bad request response
func (o *ListUserEventsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUserEventsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListUserEventsNotFoundCode is the HTTP code returned for type ListUserEventsNotFound
const ListUserEventsNotFoundCode int = 404

/*ListUserEventsNotFound Not Found

swagger:response listUserEventsNotFound
*/
type ListUserEventsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListUserEventsNotFound creates ListUserEventsNotFound with default headers values
func NewListUserEventsNotFound() *ListUserEventsNotFound {

	return &ListUserEventsNotFound{}
}

// WithPayload adds the payload to the list user events not found response
func (o *ListUserEventsNotFound) WithPayload(payload *models.Error) *ListUserEventsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list user events not found response
func (o *ListUserEventsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUserEventsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListUserEventsInternalServerErrorCode is the HTTP code returned for type ListUserEventsInternalServerError
const ListUserEventsInternalServerErrorCode int = 500

/*ListUserEventsInternalServerError Internal Server Error

swagger:response listUserEventsInternalServerError
*/
type ListUserEventsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListUserEventsInternalServerError creates ListUserEventsInternalServerError with default headers values
func NewListUserEventsInternalServerError() *ListUserEventsInternalServerError {

	return &ListUserEventsInternalServerError{}
}

// WithPayload adds the payload to the list user events internal server error response
func (o *ListUserEventsInternalServerError) WithPayload(payload *models.Error) *ListUserEventsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list user events internal server error response
func (o *ListUserEventsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUserEventsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
