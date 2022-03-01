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

// ListEventUsersOKCode is the HTTP code returned for type ListEventUsersOK
const ListEventUsersOKCode int = 200

/*ListEventUsersOK OK

swagger:response listEventUsersOK
*/
type ListEventUsersOK struct {

	/*
	  In: Body
	*/
	Payload []*ent.User `json:"body,omitempty"`
}

// NewListEventUsersOK creates ListEventUsersOK with default headers values
func NewListEventUsersOK() *ListEventUsersOK {

	return &ListEventUsersOK{}
}

// WithPayload adds the payload to the list event users o k response
func (o *ListEventUsersOK) WithPayload(payload []*ent.User) *ListEventUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list event users o k response
func (o *ListEventUsersOK) SetPayload(payload []*ent.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEventUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*ent.User, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListEventUsersBadRequestCode is the HTTP code returned for type ListEventUsersBadRequest
const ListEventUsersBadRequestCode int = 400

/*ListEventUsersBadRequest Bad request

swagger:response listEventUsersBadRequest
*/
type ListEventUsersBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListEventUsersBadRequest creates ListEventUsersBadRequest with default headers values
func NewListEventUsersBadRequest() *ListEventUsersBadRequest {

	return &ListEventUsersBadRequest{}
}

// WithPayload adds the payload to the list event users bad request response
func (o *ListEventUsersBadRequest) WithPayload(payload *models.Error) *ListEventUsersBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list event users bad request response
func (o *ListEventUsersBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEventUsersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListEventUsersUnauthorizedCode is the HTTP code returned for type ListEventUsersUnauthorized
const ListEventUsersUnauthorizedCode int = 401

/*ListEventUsersUnauthorized Unauthorized

swagger:response listEventUsersUnauthorized
*/
type ListEventUsersUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListEventUsersUnauthorized creates ListEventUsersUnauthorized with default headers values
func NewListEventUsersUnauthorized() *ListEventUsersUnauthorized {

	return &ListEventUsersUnauthorized{}
}

// WithPayload adds the payload to the list event users unauthorized response
func (o *ListEventUsersUnauthorized) WithPayload(payload *models.Error) *ListEventUsersUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list event users unauthorized response
func (o *ListEventUsersUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEventUsersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListEventUsersForbiddenCode is the HTTP code returned for type ListEventUsersForbidden
const ListEventUsersForbiddenCode int = 403

/*ListEventUsersForbidden Forbidden

swagger:response listEventUsersForbidden
*/
type ListEventUsersForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListEventUsersForbidden creates ListEventUsersForbidden with default headers values
func NewListEventUsersForbidden() *ListEventUsersForbidden {

	return &ListEventUsersForbidden{}
}

// WithPayload adds the payload to the list event users forbidden response
func (o *ListEventUsersForbidden) WithPayload(payload *models.Error) *ListEventUsersForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list event users forbidden response
func (o *ListEventUsersForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEventUsersForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListEventUsersNotFoundCode is the HTTP code returned for type ListEventUsersNotFound
const ListEventUsersNotFoundCode int = 404

/*ListEventUsersNotFound Not Found

swagger:response listEventUsersNotFound
*/
type ListEventUsersNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListEventUsersNotFound creates ListEventUsersNotFound with default headers values
func NewListEventUsersNotFound() *ListEventUsersNotFound {

	return &ListEventUsersNotFound{}
}

// WithPayload adds the payload to the list event users not found response
func (o *ListEventUsersNotFound) WithPayload(payload *models.Error) *ListEventUsersNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list event users not found response
func (o *ListEventUsersNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEventUsersNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListEventUsersInternalServerErrorCode is the HTTP code returned for type ListEventUsersInternalServerError
const ListEventUsersInternalServerErrorCode int = 500

/*ListEventUsersInternalServerError Internal Server Error

swagger:response listEventUsersInternalServerError
*/
type ListEventUsersInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListEventUsersInternalServerError creates ListEventUsersInternalServerError with default headers values
func NewListEventUsersInternalServerError() *ListEventUsersInternalServerError {

	return &ListEventUsersInternalServerError{}
}

// WithPayload adds the payload to the list event users internal server error response
func (o *ListEventUsersInternalServerError) WithPayload(payload *models.Error) *ListEventUsersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list event users internal server error response
func (o *ListEventUsersInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEventUsersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
