// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

// RemoveAdminNoContentCode is the HTTP code returned for type RemoveAdminNoContent
const RemoveAdminNoContentCode int = 204

/*RemoveAdminNoContent No Content

swagger:response removeAdminNoContent
*/
type RemoveAdminNoContent struct {
}

// NewRemoveAdminNoContent creates RemoveAdminNoContent with default headers values
func NewRemoveAdminNoContent() *RemoveAdminNoContent {

	return &RemoveAdminNoContent{}
}

// WriteResponse to the client
func (o *RemoveAdminNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// RemoveAdminBadRequestCode is the HTTP code returned for type RemoveAdminBadRequest
const RemoveAdminBadRequestCode int = 400

/*RemoveAdminBadRequest Bad request

swagger:response removeAdminBadRequest
*/
type RemoveAdminBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRemoveAdminBadRequest creates RemoveAdminBadRequest with default headers values
func NewRemoveAdminBadRequest() *RemoveAdminBadRequest {

	return &RemoveAdminBadRequest{}
}

// WithPayload adds the payload to the remove admin bad request response
func (o *RemoveAdminBadRequest) WithPayload(payload *models.Error) *RemoveAdminBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove admin bad request response
func (o *RemoveAdminBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveAdminBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveAdminNotFoundCode is the HTTP code returned for type RemoveAdminNotFound
const RemoveAdminNotFoundCode int = 404

/*RemoveAdminNotFound Not Found

swagger:response removeAdminNotFound
*/
type RemoveAdminNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRemoveAdminNotFound creates RemoveAdminNotFound with default headers values
func NewRemoveAdminNotFound() *RemoveAdminNotFound {

	return &RemoveAdminNotFound{}
}

// WithPayload adds the payload to the remove admin not found response
func (o *RemoveAdminNotFound) WithPayload(payload *models.Error) *RemoveAdminNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove admin not found response
func (o *RemoveAdminNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveAdminNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveAdminInternalServerErrorCode is the HTTP code returned for type RemoveAdminInternalServerError
const RemoveAdminInternalServerErrorCode int = 500

/*RemoveAdminInternalServerError Internal Server Error

swagger:response removeAdminInternalServerError
*/
type RemoveAdminInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRemoveAdminInternalServerError creates RemoveAdminInternalServerError with default headers values
func NewRemoveAdminInternalServerError() *RemoveAdminInternalServerError {

	return &RemoveAdminInternalServerError{}
}

// WithPayload adds the payload to the remove admin internal server error response
func (o *RemoveAdminInternalServerError) WithPayload(payload *models.Error) *RemoveAdminInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove admin internal server error response
func (o *RemoveAdminInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveAdminInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
