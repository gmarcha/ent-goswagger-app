// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gamarcha/ent-goswagger-app/goswagger/models"
)

// GetAuthCallbackOKCode is the HTTP code returned for type GetAuthCallbackOK
const GetAuthCallbackOKCode int = 200

/*GetAuthCallbackOK OK

swagger:response getAuthCallbackOK
*/
type GetAuthCallbackOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetAuthCallbackOK creates GetAuthCallbackOK with default headers values
func NewGetAuthCallbackOK() *GetAuthCallbackOK {

	return &GetAuthCallbackOK{}
}

// WithPayload adds the payload to the get auth callback o k response
func (o *GetAuthCallbackOK) WithPayload(payload string) *GetAuthCallbackOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get auth callback o k response
func (o *GetAuthCallbackOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAuthCallbackOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAuthCallbackInternalServerErrorCode is the HTTP code returned for type GetAuthCallbackInternalServerError
const GetAuthCallbackInternalServerErrorCode int = 500

/*GetAuthCallbackInternalServerError Internal Server Error

swagger:response getAuthCallbackInternalServerError
*/
type GetAuthCallbackInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAuthCallbackInternalServerError creates GetAuthCallbackInternalServerError with default headers values
func NewGetAuthCallbackInternalServerError() *GetAuthCallbackInternalServerError {

	return &GetAuthCallbackInternalServerError{}
}

// WithPayload adds the payload to the get auth callback internal server error response
func (o *GetAuthCallbackInternalServerError) WithPayload(payload *models.Error) *GetAuthCallbackInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get auth callback internal server error response
func (o *GetAuthCallbackInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAuthCallbackInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}