// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

// AddAdminCreatedCode is the HTTP code returned for type AddAdminCreated
const AddAdminCreatedCode int = 201

/*AddAdminCreated Created

swagger:response addAdminCreated
*/
type AddAdminCreated struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewAddAdminCreated creates AddAdminCreated with default headers values
func NewAddAdminCreated() *AddAdminCreated {

	return &AddAdminCreated{}
}

// WithPayload adds the payload to the add admin created response
func (o *AddAdminCreated) WithPayload(payload string) *AddAdminCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add admin created response
func (o *AddAdminCreated) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAdminCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// AddAdminBadRequestCode is the HTTP code returned for type AddAdminBadRequest
const AddAdminBadRequestCode int = 400

/*AddAdminBadRequest Bad request

swagger:response addAdminBadRequest
*/
type AddAdminBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddAdminBadRequest creates AddAdminBadRequest with default headers values
func NewAddAdminBadRequest() *AddAdminBadRequest {

	return &AddAdminBadRequest{}
}

// WithPayload adds the payload to the add admin bad request response
func (o *AddAdminBadRequest) WithPayload(payload *models.Error) *AddAdminBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add admin bad request response
func (o *AddAdminBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAdminBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddAdminNotFoundCode is the HTTP code returned for type AddAdminNotFound
const AddAdminNotFoundCode int = 404

/*AddAdminNotFound Not Found

swagger:response addAdminNotFound
*/
type AddAdminNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddAdminNotFound creates AddAdminNotFound with default headers values
func NewAddAdminNotFound() *AddAdminNotFound {

	return &AddAdminNotFound{}
}

// WithPayload adds the payload to the add admin not found response
func (o *AddAdminNotFound) WithPayload(payload *models.Error) *AddAdminNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add admin not found response
func (o *AddAdminNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAdminNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddAdminInternalServerErrorCode is the HTTP code returned for type AddAdminInternalServerError
const AddAdminInternalServerErrorCode int = 500

/*AddAdminInternalServerError Internal Server Error

swagger:response addAdminInternalServerError
*/
type AddAdminInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddAdminInternalServerError creates AddAdminInternalServerError with default headers values
func NewAddAdminInternalServerError() *AddAdminInternalServerError {

	return &AddAdminInternalServerError{}
}

// WithPayload adds the payload to the add admin internal server error response
func (o *AddAdminInternalServerError) WithPayload(payload *models.Error) *AddAdminInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add admin internal server error response
func (o *AddAdminInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAdminInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
