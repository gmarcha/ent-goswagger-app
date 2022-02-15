// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gamarcha/ent-goswagger-app/internal/goswagger/models"
)

// DeleteUserNoContentCode is the HTTP code returned for type DeleteUserNoContent
const DeleteUserNoContentCode int = 204

/*DeleteUserNoContent No Content

swagger:response deleteUserNoContent
*/
type DeleteUserNoContent struct {
}

// NewDeleteUserNoContent creates DeleteUserNoContent with default headers values
func NewDeleteUserNoContent() *DeleteUserNoContent {

	return &DeleteUserNoContent{}
}

// WriteResponse to the client
func (o *DeleteUserNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteUserBadRequestCode is the HTTP code returned for type DeleteUserBadRequest
const DeleteUserBadRequestCode int = 400

/*DeleteUserBadRequest Bad request

swagger:response deleteUserBadRequest
*/
type DeleteUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteUserBadRequest creates DeleteUserBadRequest with default headers values
func NewDeleteUserBadRequest() *DeleteUserBadRequest {

	return &DeleteUserBadRequest{}
}

// WithPayload adds the payload to the delete user bad request response
func (o *DeleteUserBadRequest) WithPayload(payload *models.Error) *DeleteUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user bad request response
func (o *DeleteUserBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteUserNotFoundCode is the HTTP code returned for type DeleteUserNotFound
const DeleteUserNotFoundCode int = 404

/*DeleteUserNotFound Not Found

swagger:response deleteUserNotFound
*/
type DeleteUserNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteUserNotFound creates DeleteUserNotFound with default headers values
func NewDeleteUserNotFound() *DeleteUserNotFound {

	return &DeleteUserNotFound{}
}

// WithPayload adds the payload to the delete user not found response
func (o *DeleteUserNotFound) WithPayload(payload *models.Error) *DeleteUserNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user not found response
func (o *DeleteUserNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteUserInternalServerErrorCode is the HTTP code returned for type DeleteUserInternalServerError
const DeleteUserInternalServerErrorCode int = 500

/*DeleteUserInternalServerError Internal Server Error

swagger:response deleteUserInternalServerError
*/
type DeleteUserInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteUserInternalServerError creates DeleteUserInternalServerError with default headers values
func NewDeleteUserInternalServerError() *DeleteUserInternalServerError {

	return &DeleteUserInternalServerError{}
}

// WithPayload adds the payload to the delete user internal server error response
func (o *DeleteUserInternalServerError) WithPayload(payload *models.Error) *DeleteUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user internal server error response
func (o *DeleteUserInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
