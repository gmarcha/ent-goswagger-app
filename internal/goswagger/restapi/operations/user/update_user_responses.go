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

// UpdateUserOKCode is the HTTP code returned for type UpdateUserOK
const UpdateUserOKCode int = 200

/*UpdateUserOK OK

swagger:response updateUserOK
*/
type UpdateUserOK struct {

	/*
	  In: Body
	*/
	Payload *ent.User `json:"body,omitempty"`
}

// NewUpdateUserOK creates UpdateUserOK with default headers values
func NewUpdateUserOK() *UpdateUserOK {

	return &UpdateUserOK{}
}

// WithPayload adds the payload to the update user o k response
func (o *UpdateUserOK) WithPayload(payload *ent.User) *UpdateUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user o k response
func (o *UpdateUserOK) SetPayload(payload *ent.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateUserBadRequestCode is the HTTP code returned for type UpdateUserBadRequest
const UpdateUserBadRequestCode int = 400

/*UpdateUserBadRequest Bad request

swagger:response updateUserBadRequest
*/
type UpdateUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateUserBadRequest creates UpdateUserBadRequest with default headers values
func NewUpdateUserBadRequest() *UpdateUserBadRequest {

	return &UpdateUserBadRequest{}
}

// WithPayload adds the payload to the update user bad request response
func (o *UpdateUserBadRequest) WithPayload(payload *models.Error) *UpdateUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user bad request response
func (o *UpdateUserBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateUserUnauthorizedCode is the HTTP code returned for type UpdateUserUnauthorized
const UpdateUserUnauthorizedCode int = 401

/*UpdateUserUnauthorized Unauthorized

swagger:response updateUserUnauthorized
*/
type UpdateUserUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateUserUnauthorized creates UpdateUserUnauthorized with default headers values
func NewUpdateUserUnauthorized() *UpdateUserUnauthorized {

	return &UpdateUserUnauthorized{}
}

// WithPayload adds the payload to the update user unauthorized response
func (o *UpdateUserUnauthorized) WithPayload(payload *models.Error) *UpdateUserUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user unauthorized response
func (o *UpdateUserUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateUserForbiddenCode is the HTTP code returned for type UpdateUserForbidden
const UpdateUserForbiddenCode int = 403

/*UpdateUserForbidden Forbidden

swagger:response updateUserForbidden
*/
type UpdateUserForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateUserForbidden creates UpdateUserForbidden with default headers values
func NewUpdateUserForbidden() *UpdateUserForbidden {

	return &UpdateUserForbidden{}
}

// WithPayload adds the payload to the update user forbidden response
func (o *UpdateUserForbidden) WithPayload(payload *models.Error) *UpdateUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user forbidden response
func (o *UpdateUserForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateUserNotFoundCode is the HTTP code returned for type UpdateUserNotFound
const UpdateUserNotFoundCode int = 404

/*UpdateUserNotFound Not Found

swagger:response updateUserNotFound
*/
type UpdateUserNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateUserNotFound creates UpdateUserNotFound with default headers values
func NewUpdateUserNotFound() *UpdateUserNotFound {

	return &UpdateUserNotFound{}
}

// WithPayload adds the payload to the update user not found response
func (o *UpdateUserNotFound) WithPayload(payload *models.Error) *UpdateUserNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user not found response
func (o *UpdateUserNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateUserInternalServerErrorCode is the HTTP code returned for type UpdateUserInternalServerError
const UpdateUserInternalServerErrorCode int = 500

/*UpdateUserInternalServerError Internal Server Error

swagger:response updateUserInternalServerError
*/
type UpdateUserInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateUserInternalServerError creates UpdateUserInternalServerError with default headers values
func NewUpdateUserInternalServerError() *UpdateUserInternalServerError {

	return &UpdateUserInternalServerError{}
}

// WithPayload adds the payload to the update user internal server error response
func (o *UpdateUserInternalServerError) WithPayload(payload *models.Error) *UpdateUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user internal server error response
func (o *UpdateUserInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
