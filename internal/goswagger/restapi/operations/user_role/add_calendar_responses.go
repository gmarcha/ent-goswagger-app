// Code generated by go-swagger; DO NOT EDIT.

package user_role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

// AddCalendarCreatedCode is the HTTP code returned for type AddCalendarCreated
const AddCalendarCreatedCode int = 201

/*AddCalendarCreated Created

swagger:response addCalendarCreated
*/
type AddCalendarCreated struct {

	/*
	  In: Body
	*/
	Payload []*ent.Role `json:"body,omitempty"`
}

// NewAddCalendarCreated creates AddCalendarCreated with default headers values
func NewAddCalendarCreated() *AddCalendarCreated {

	return &AddCalendarCreated{}
}

// WithPayload adds the payload to the add calendar created response
func (o *AddCalendarCreated) WithPayload(payload []*ent.Role) *AddCalendarCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add calendar created response
func (o *AddCalendarCreated) SetPayload(payload []*ent.Role) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddCalendarCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*ent.Role, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// AddCalendarBadRequestCode is the HTTP code returned for type AddCalendarBadRequest
const AddCalendarBadRequestCode int = 400

/*AddCalendarBadRequest Bad request

swagger:response addCalendarBadRequest
*/
type AddCalendarBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddCalendarBadRequest creates AddCalendarBadRequest with default headers values
func NewAddCalendarBadRequest() *AddCalendarBadRequest {

	return &AddCalendarBadRequest{}
}

// WithPayload adds the payload to the add calendar bad request response
func (o *AddCalendarBadRequest) WithPayload(payload *models.Error) *AddCalendarBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add calendar bad request response
func (o *AddCalendarBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddCalendarBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddCalendarUnauthorizedCode is the HTTP code returned for type AddCalendarUnauthorized
const AddCalendarUnauthorizedCode int = 401

/*AddCalendarUnauthorized Unauthorized

swagger:response addCalendarUnauthorized
*/
type AddCalendarUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddCalendarUnauthorized creates AddCalendarUnauthorized with default headers values
func NewAddCalendarUnauthorized() *AddCalendarUnauthorized {

	return &AddCalendarUnauthorized{}
}

// WithPayload adds the payload to the add calendar unauthorized response
func (o *AddCalendarUnauthorized) WithPayload(payload *models.Error) *AddCalendarUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add calendar unauthorized response
func (o *AddCalendarUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddCalendarUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddCalendarForbiddenCode is the HTTP code returned for type AddCalendarForbidden
const AddCalendarForbiddenCode int = 403

/*AddCalendarForbidden Forbidden

swagger:response addCalendarForbidden
*/
type AddCalendarForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddCalendarForbidden creates AddCalendarForbidden with default headers values
func NewAddCalendarForbidden() *AddCalendarForbidden {

	return &AddCalendarForbidden{}
}

// WithPayload adds the payload to the add calendar forbidden response
func (o *AddCalendarForbidden) WithPayload(payload *models.Error) *AddCalendarForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add calendar forbidden response
func (o *AddCalendarForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddCalendarForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddCalendarNotFoundCode is the HTTP code returned for type AddCalendarNotFound
const AddCalendarNotFoundCode int = 404

/*AddCalendarNotFound Not Found

swagger:response addCalendarNotFound
*/
type AddCalendarNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddCalendarNotFound creates AddCalendarNotFound with default headers values
func NewAddCalendarNotFound() *AddCalendarNotFound {

	return &AddCalendarNotFound{}
}

// WithPayload adds the payload to the add calendar not found response
func (o *AddCalendarNotFound) WithPayload(payload *models.Error) *AddCalendarNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add calendar not found response
func (o *AddCalendarNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddCalendarNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddCalendarInternalServerErrorCode is the HTTP code returned for type AddCalendarInternalServerError
const AddCalendarInternalServerErrorCode int = 500

/*AddCalendarInternalServerError Internal Server Error

swagger:response addCalendarInternalServerError
*/
type AddCalendarInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddCalendarInternalServerError creates AddCalendarInternalServerError with default headers values
func NewAddCalendarInternalServerError() *AddCalendarInternalServerError {

	return &AddCalendarInternalServerError{}
}

// WithPayload adds the payload to the add calendar internal server error response
func (o *AddCalendarInternalServerError) WithPayload(payload *models.Error) *AddCalendarInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add calendar internal server error response
func (o *AddCalendarInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddCalendarInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
