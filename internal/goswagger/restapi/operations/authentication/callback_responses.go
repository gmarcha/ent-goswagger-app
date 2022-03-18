// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

// CallbackOKCode is the HTTP code returned for type CallbackOK
const CallbackOKCode int = 200

/*CallbackOK OK

swagger:response callbackOK
*/
type CallbackOK struct {

	/*
	  In: Body
	*/
	Payload *models.Token `json:"body,omitempty"`
}

// NewCallbackOK creates CallbackOK with default headers values
func NewCallbackOK() *CallbackOK {

	return &CallbackOK{}
}

// WithPayload adds the payload to the callback o k response
func (o *CallbackOK) WithPayload(payload *models.Token) *CallbackOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the callback o k response
func (o *CallbackOK) SetPayload(payload *models.Token) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CallbackOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CallbackUnauthorizedCode is the HTTP code returned for type CallbackUnauthorized
const CallbackUnauthorizedCode int = 401

/*CallbackUnauthorized Unprocessable Entity

swagger:response callbackUnauthorized
*/
type CallbackUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCallbackUnauthorized creates CallbackUnauthorized with default headers values
func NewCallbackUnauthorized() *CallbackUnauthorized {

	return &CallbackUnauthorized{}
}

// WithPayload adds the payload to the callback unauthorized response
func (o *CallbackUnauthorized) WithPayload(payload *models.Error) *CallbackUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the callback unauthorized response
func (o *CallbackUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CallbackUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CallbackUnprocessableEntityCode is the HTTP code returned for type CallbackUnprocessableEntity
const CallbackUnprocessableEntityCode int = 422

/*CallbackUnprocessableEntity Unprocessable Entity

swagger:response callbackUnprocessableEntity
*/
type CallbackUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCallbackUnprocessableEntity creates CallbackUnprocessableEntity with default headers values
func NewCallbackUnprocessableEntity() *CallbackUnprocessableEntity {

	return &CallbackUnprocessableEntity{}
}

// WithPayload adds the payload to the callback unprocessable entity response
func (o *CallbackUnprocessableEntity) WithPayload(payload *models.Error) *CallbackUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the callback unprocessable entity response
func (o *CallbackUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CallbackUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CallbackInternalServerErrorCode is the HTTP code returned for type CallbackInternalServerError
const CallbackInternalServerErrorCode int = 500

/*CallbackInternalServerError Internal Server Error

swagger:response callbackInternalServerError
*/
type CallbackInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCallbackInternalServerError creates CallbackInternalServerError with default headers values
func NewCallbackInternalServerError() *CallbackInternalServerError {

	return &CallbackInternalServerError{}
}

// WithPayload adds the payload to the callback internal server error response
func (o *CallbackInternalServerError) WithPayload(payload *models.Error) *CallbackInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the callback internal server error response
func (o *CallbackInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CallbackInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
