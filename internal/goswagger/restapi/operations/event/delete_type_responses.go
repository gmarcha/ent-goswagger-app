// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

// DeleteTypeNoContentCode is the HTTP code returned for type DeleteTypeNoContent
const DeleteTypeNoContentCode int = 204

/*DeleteTypeNoContent No content

swagger:response deleteTypeNoContent
*/
type DeleteTypeNoContent struct {
}

// NewDeleteTypeNoContent creates DeleteTypeNoContent with default headers values
func NewDeleteTypeNoContent() *DeleteTypeNoContent {

	return &DeleteTypeNoContent{}
}

// WriteResponse to the client
func (o *DeleteTypeNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteTypeBadRequestCode is the HTTP code returned for type DeleteTypeBadRequest
const DeleteTypeBadRequestCode int = 400

/*DeleteTypeBadRequest Bad request

swagger:response deleteTypeBadRequest
*/
type DeleteTypeBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTypeBadRequest creates DeleteTypeBadRequest with default headers values
func NewDeleteTypeBadRequest() *DeleteTypeBadRequest {

	return &DeleteTypeBadRequest{}
}

// WithPayload adds the payload to the delete type bad request response
func (o *DeleteTypeBadRequest) WithPayload(payload *models.Error) *DeleteTypeBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete type bad request response
func (o *DeleteTypeBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTypeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteTypeNotFoundCode is the HTTP code returned for type DeleteTypeNotFound
const DeleteTypeNotFoundCode int = 404

/*DeleteTypeNotFound Not Found

swagger:response deleteTypeNotFound
*/
type DeleteTypeNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTypeNotFound creates DeleteTypeNotFound with default headers values
func NewDeleteTypeNotFound() *DeleteTypeNotFound {

	return &DeleteTypeNotFound{}
}

// WithPayload adds the payload to the delete type not found response
func (o *DeleteTypeNotFound) WithPayload(payload *models.Error) *DeleteTypeNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete type not found response
func (o *DeleteTypeNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTypeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteTypeInternalServerErrorCode is the HTTP code returned for type DeleteTypeInternalServerError
const DeleteTypeInternalServerErrorCode int = 500

/*DeleteTypeInternalServerError Internal Server Error

swagger:response deleteTypeInternalServerError
*/
type DeleteTypeInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTypeInternalServerError creates DeleteTypeInternalServerError with default headers values
func NewDeleteTypeInternalServerError() *DeleteTypeInternalServerError {

	return &DeleteTypeInternalServerError{}
}

// WithPayload adds the payload to the delete type internal server error response
func (o *DeleteTypeInternalServerError) WithPayload(payload *models.Error) *DeleteTypeInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete type internal server error response
func (o *DeleteTypeInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTypeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
