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

// SetEventTypeOKCode is the HTTP code returned for type SetEventTypeOK
const SetEventTypeOKCode int = 200

/*SetEventTypeOK OK

swagger:response setEventTypeOK
*/
type SetEventTypeOK struct {

	/*
	  In: Body
	*/
	Payload *ent.EventType `json:"body,omitempty"`
}

// NewSetEventTypeOK creates SetEventTypeOK with default headers values
func NewSetEventTypeOK() *SetEventTypeOK {

	return &SetEventTypeOK{}
}

// WithPayload adds the payload to the set event type o k response
func (o *SetEventTypeOK) WithPayload(payload *ent.EventType) *SetEventTypeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set event type o k response
func (o *SetEventTypeOK) SetPayload(payload *ent.EventType) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetEventTypeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SetEventTypeBadRequestCode is the HTTP code returned for type SetEventTypeBadRequest
const SetEventTypeBadRequestCode int = 400

/*SetEventTypeBadRequest Bad request

swagger:response setEventTypeBadRequest
*/
type SetEventTypeBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSetEventTypeBadRequest creates SetEventTypeBadRequest with default headers values
func NewSetEventTypeBadRequest() *SetEventTypeBadRequest {

	return &SetEventTypeBadRequest{}
}

// WithPayload adds the payload to the set event type bad request response
func (o *SetEventTypeBadRequest) WithPayload(payload *models.Error) *SetEventTypeBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set event type bad request response
func (o *SetEventTypeBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetEventTypeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SetEventTypeUnauthorizedCode is the HTTP code returned for type SetEventTypeUnauthorized
const SetEventTypeUnauthorizedCode int = 401

/*SetEventTypeUnauthorized Unauthorized

swagger:response setEventTypeUnauthorized
*/
type SetEventTypeUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSetEventTypeUnauthorized creates SetEventTypeUnauthorized with default headers values
func NewSetEventTypeUnauthorized() *SetEventTypeUnauthorized {

	return &SetEventTypeUnauthorized{}
}

// WithPayload adds the payload to the set event type unauthorized response
func (o *SetEventTypeUnauthorized) WithPayload(payload *models.Error) *SetEventTypeUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set event type unauthorized response
func (o *SetEventTypeUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetEventTypeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SetEventTypeForbiddenCode is the HTTP code returned for type SetEventTypeForbidden
const SetEventTypeForbiddenCode int = 403

/*SetEventTypeForbidden Forbidden

swagger:response setEventTypeForbidden
*/
type SetEventTypeForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSetEventTypeForbidden creates SetEventTypeForbidden with default headers values
func NewSetEventTypeForbidden() *SetEventTypeForbidden {

	return &SetEventTypeForbidden{}
}

// WithPayload adds the payload to the set event type forbidden response
func (o *SetEventTypeForbidden) WithPayload(payload *models.Error) *SetEventTypeForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set event type forbidden response
func (o *SetEventTypeForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetEventTypeForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SetEventTypeNotFoundCode is the HTTP code returned for type SetEventTypeNotFound
const SetEventTypeNotFoundCode int = 404

/*SetEventTypeNotFound Not Found

swagger:response setEventTypeNotFound
*/
type SetEventTypeNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSetEventTypeNotFound creates SetEventTypeNotFound with default headers values
func NewSetEventTypeNotFound() *SetEventTypeNotFound {

	return &SetEventTypeNotFound{}
}

// WithPayload adds the payload to the set event type not found response
func (o *SetEventTypeNotFound) WithPayload(payload *models.Error) *SetEventTypeNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set event type not found response
func (o *SetEventTypeNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetEventTypeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SetEventTypeInternalServerErrorCode is the HTTP code returned for type SetEventTypeInternalServerError
const SetEventTypeInternalServerErrorCode int = 500

/*SetEventTypeInternalServerError Internal Server Error

swagger:response setEventTypeInternalServerError
*/
type SetEventTypeInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSetEventTypeInternalServerError creates SetEventTypeInternalServerError with default headers values
func NewSetEventTypeInternalServerError() *SetEventTypeInternalServerError {

	return &SetEventTypeInternalServerError{}
}

// WithPayload adds the payload to the set event type internal server error response
func (o *SetEventTypeInternalServerError) WithPayload(payload *models.Error) *SetEventTypeInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set event type internal server error response
func (o *SetEventTypeInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetEventTypeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}