// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gamarcha/ent-goswagger-app/goswagger/models"
)

// ReadEventOKCode is the HTTP code returned for type ReadEventOK
const ReadEventOKCode int = 200

/*ReadEventOK OK

swagger:response readEventOK
*/
type ReadEventOK struct {

	/*
	  In: Body
	*/
	Payload *models.Event `json:"body,omitempty"`
}

// NewReadEventOK creates ReadEventOK with default headers values
func NewReadEventOK() *ReadEventOK {

	return &ReadEventOK{}
}

// WithPayload adds the payload to the read event o k response
func (o *ReadEventOK) WithPayload(payload *models.Event) *ReadEventOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read event o k response
func (o *ReadEventOK) SetPayload(payload *models.Event) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadEventBadRequestCode is the HTTP code returned for type ReadEventBadRequest
const ReadEventBadRequestCode int = 400

/*ReadEventBadRequest Bad request

swagger:response readEventBadRequest
*/
type ReadEventBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReadEventBadRequest creates ReadEventBadRequest with default headers values
func NewReadEventBadRequest() *ReadEventBadRequest {

	return &ReadEventBadRequest{}
}

// WithPayload adds the payload to the read event bad request response
func (o *ReadEventBadRequest) WithPayload(payload *models.Error) *ReadEventBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read event bad request response
func (o *ReadEventBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadEventBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadEventNotFoundCode is the HTTP code returned for type ReadEventNotFound
const ReadEventNotFoundCode int = 404

/*ReadEventNotFound Not Found

swagger:response readEventNotFound
*/
type ReadEventNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReadEventNotFound creates ReadEventNotFound with default headers values
func NewReadEventNotFound() *ReadEventNotFound {

	return &ReadEventNotFound{}
}

// WithPayload adds the payload to the read event not found response
func (o *ReadEventNotFound) WithPayload(payload *models.Error) *ReadEventNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read event not found response
func (o *ReadEventNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadEventNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadEventInternalServerErrorCode is the HTTP code returned for type ReadEventInternalServerError
const ReadEventInternalServerErrorCode int = 500

/*ReadEventInternalServerError Internal Server Error

swagger:response readEventInternalServerError
*/
type ReadEventInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReadEventInternalServerError creates ReadEventInternalServerError with default headers values
func NewReadEventInternalServerError() *ReadEventInternalServerError {

	return &ReadEventInternalServerError{}
}

// WithPayload adds the payload to the read event internal server error response
func (o *ReadEventInternalServerError) WithPayload(payload *models.Error) *ReadEventInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read event internal server error response
func (o *ReadEventInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadEventInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}