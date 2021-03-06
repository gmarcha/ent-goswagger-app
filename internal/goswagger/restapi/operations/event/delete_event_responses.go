// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

// DeleteEventNoContentCode is the HTTP code returned for type DeleteEventNoContent
const DeleteEventNoContentCode int = 204

/*DeleteEventNoContent No Content

swagger:response deleteEventNoContent
*/
type DeleteEventNoContent struct {
}

// NewDeleteEventNoContent creates DeleteEventNoContent with default headers values
func NewDeleteEventNoContent() *DeleteEventNoContent {

	return &DeleteEventNoContent{}
}

// WriteResponse to the client
func (o *DeleteEventNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteEventBadRequestCode is the HTTP code returned for type DeleteEventBadRequest
const DeleteEventBadRequestCode int = 400

/*DeleteEventBadRequest Bad request

swagger:response deleteEventBadRequest
*/
type DeleteEventBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteEventBadRequest creates DeleteEventBadRequest with default headers values
func NewDeleteEventBadRequest() *DeleteEventBadRequest {

	return &DeleteEventBadRequest{}
}

// WithPayload adds the payload to the delete event bad request response
func (o *DeleteEventBadRequest) WithPayload(payload *models.Error) *DeleteEventBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete event bad request response
func (o *DeleteEventBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteEventBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteEventUnauthorizedCode is the HTTP code returned for type DeleteEventUnauthorized
const DeleteEventUnauthorizedCode int = 401

/*DeleteEventUnauthorized Unauthorized

swagger:response deleteEventUnauthorized
*/
type DeleteEventUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteEventUnauthorized creates DeleteEventUnauthorized with default headers values
func NewDeleteEventUnauthorized() *DeleteEventUnauthorized {

	return &DeleteEventUnauthorized{}
}

// WithPayload adds the payload to the delete event unauthorized response
func (o *DeleteEventUnauthorized) WithPayload(payload *models.Error) *DeleteEventUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete event unauthorized response
func (o *DeleteEventUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteEventUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteEventForbiddenCode is the HTTP code returned for type DeleteEventForbidden
const DeleteEventForbiddenCode int = 403

/*DeleteEventForbidden Forbidden

swagger:response deleteEventForbidden
*/
type DeleteEventForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteEventForbidden creates DeleteEventForbidden with default headers values
func NewDeleteEventForbidden() *DeleteEventForbidden {

	return &DeleteEventForbidden{}
}

// WithPayload adds the payload to the delete event forbidden response
func (o *DeleteEventForbidden) WithPayload(payload *models.Error) *DeleteEventForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete event forbidden response
func (o *DeleteEventForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteEventForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteEventNotFoundCode is the HTTP code returned for type DeleteEventNotFound
const DeleteEventNotFoundCode int = 404

/*DeleteEventNotFound Not Found

swagger:response deleteEventNotFound
*/
type DeleteEventNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteEventNotFound creates DeleteEventNotFound with default headers values
func NewDeleteEventNotFound() *DeleteEventNotFound {

	return &DeleteEventNotFound{}
}

// WithPayload adds the payload to the delete event not found response
func (o *DeleteEventNotFound) WithPayload(payload *models.Error) *DeleteEventNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete event not found response
func (o *DeleteEventNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteEventNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteEventInternalServerErrorCode is the HTTP code returned for type DeleteEventInternalServerError
const DeleteEventInternalServerErrorCode int = 500

/*DeleteEventInternalServerError Internal Server Error

swagger:response deleteEventInternalServerError
*/
type DeleteEventInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteEventInternalServerError creates DeleteEventInternalServerError with default headers values
func NewDeleteEventInternalServerError() *DeleteEventInternalServerError {

	return &DeleteEventInternalServerError{}
}

// WithPayload adds the payload to the delete event internal server error response
func (o *DeleteEventInternalServerError) WithPayload(payload *models.Error) *DeleteEventInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete event internal server error response
func (o *DeleteEventInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteEventInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
