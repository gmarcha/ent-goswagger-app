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

// GetEventTypeOKCode is the HTTP code returned for type GetEventTypeOK
const GetEventTypeOKCode int = 200

/*GetEventTypeOK OK

swagger:response getEventTypeOK
*/
type GetEventTypeOK struct {

	/*
	  In: Body
	*/
	Payload *ent.EventType `json:"body,omitempty"`
}

// NewGetEventTypeOK creates GetEventTypeOK with default headers values
func NewGetEventTypeOK() *GetEventTypeOK {

	return &GetEventTypeOK{}
}

// WithPayload adds the payload to the get event type o k response
func (o *GetEventTypeOK) WithPayload(payload *ent.EventType) *GetEventTypeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get event type o k response
func (o *GetEventTypeOK) SetPayload(payload *ent.EventType) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventTypeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEventTypeBadRequestCode is the HTTP code returned for type GetEventTypeBadRequest
const GetEventTypeBadRequestCode int = 400

/*GetEventTypeBadRequest Bad request

swagger:response getEventTypeBadRequest
*/
type GetEventTypeBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetEventTypeBadRequest creates GetEventTypeBadRequest with default headers values
func NewGetEventTypeBadRequest() *GetEventTypeBadRequest {

	return &GetEventTypeBadRequest{}
}

// WithPayload adds the payload to the get event type bad request response
func (o *GetEventTypeBadRequest) WithPayload(payload *models.Error) *GetEventTypeBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get event type bad request response
func (o *GetEventTypeBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventTypeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEventTypeUnauthorizedCode is the HTTP code returned for type GetEventTypeUnauthorized
const GetEventTypeUnauthorizedCode int = 401

/*GetEventTypeUnauthorized Unauthorized

swagger:response getEventTypeUnauthorized
*/
type GetEventTypeUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetEventTypeUnauthorized creates GetEventTypeUnauthorized with default headers values
func NewGetEventTypeUnauthorized() *GetEventTypeUnauthorized {

	return &GetEventTypeUnauthorized{}
}

// WithPayload adds the payload to the get event type unauthorized response
func (o *GetEventTypeUnauthorized) WithPayload(payload *models.Error) *GetEventTypeUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get event type unauthorized response
func (o *GetEventTypeUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventTypeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEventTypeForbiddenCode is the HTTP code returned for type GetEventTypeForbidden
const GetEventTypeForbiddenCode int = 403

/*GetEventTypeForbidden Forbidden

swagger:response getEventTypeForbidden
*/
type GetEventTypeForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetEventTypeForbidden creates GetEventTypeForbidden with default headers values
func NewGetEventTypeForbidden() *GetEventTypeForbidden {

	return &GetEventTypeForbidden{}
}

// WithPayload adds the payload to the get event type forbidden response
func (o *GetEventTypeForbidden) WithPayload(payload *models.Error) *GetEventTypeForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get event type forbidden response
func (o *GetEventTypeForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventTypeForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEventTypeNotFoundCode is the HTTP code returned for type GetEventTypeNotFound
const GetEventTypeNotFoundCode int = 404

/*GetEventTypeNotFound Not Found

swagger:response getEventTypeNotFound
*/
type GetEventTypeNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetEventTypeNotFound creates GetEventTypeNotFound with default headers values
func NewGetEventTypeNotFound() *GetEventTypeNotFound {

	return &GetEventTypeNotFound{}
}

// WithPayload adds the payload to the get event type not found response
func (o *GetEventTypeNotFound) WithPayload(payload *models.Error) *GetEventTypeNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get event type not found response
func (o *GetEventTypeNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventTypeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEventTypeInternalServerErrorCode is the HTTP code returned for type GetEventTypeInternalServerError
const GetEventTypeInternalServerErrorCode int = 500

/*GetEventTypeInternalServerError Internal Server Error

swagger:response getEventTypeInternalServerError
*/
type GetEventTypeInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetEventTypeInternalServerError creates GetEventTypeInternalServerError with default headers values
func NewGetEventTypeInternalServerError() *GetEventTypeInternalServerError {

	return &GetEventTypeInternalServerError{}
}

// WithPayload adds the payload to the get event type internal server error response
func (o *GetEventTypeInternalServerError) WithPayload(payload *models.Error) *GetEventTypeInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get event type internal server error response
func (o *GetEventTypeInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventTypeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}