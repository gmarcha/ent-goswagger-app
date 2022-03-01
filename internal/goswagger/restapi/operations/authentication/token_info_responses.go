// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

// TokenInfoOKCode is the HTTP code returned for type TokenInfoOK
const TokenInfoOKCode int = 200

/*TokenInfoOK OK

swagger:response tokenInfoOK
*/
type TokenInfoOK struct {

	/*
	  In: Body
	*/
	Payload *TokenInfoOKBody `json:"body,omitempty"`
}

// NewTokenInfoOK creates TokenInfoOK with default headers values
func NewTokenInfoOK() *TokenInfoOK {

	return &TokenInfoOK{}
}

// WithPayload adds the payload to the token info o k response
func (o *TokenInfoOK) WithPayload(payload *TokenInfoOKBody) *TokenInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the token info o k response
func (o *TokenInfoOK) SetPayload(payload *TokenInfoOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TokenInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TokenInfoUnauthorizedCode is the HTTP code returned for type TokenInfoUnauthorized
const TokenInfoUnauthorizedCode int = 401

/*TokenInfoUnauthorized Unauthorized

swagger:response tokenInfoUnauthorized
*/
type TokenInfoUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewTokenInfoUnauthorized creates TokenInfoUnauthorized with default headers values
func NewTokenInfoUnauthorized() *TokenInfoUnauthorized {

	return &TokenInfoUnauthorized{}
}

// WithPayload adds the payload to the token info unauthorized response
func (o *TokenInfoUnauthorized) WithPayload(payload *models.Error) *TokenInfoUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the token info unauthorized response
func (o *TokenInfoUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TokenInfoUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TokenInfoInternalServerErrorCode is the HTTP code returned for type TokenInfoInternalServerError
const TokenInfoInternalServerErrorCode int = 500

/*TokenInfoInternalServerError Internal Server Error

swagger:response tokenInfoInternalServerError
*/
type TokenInfoInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewTokenInfoInternalServerError creates TokenInfoInternalServerError with default headers values
func NewTokenInfoInternalServerError() *TokenInfoInternalServerError {

	return &TokenInfoInternalServerError{}
}

// WithPayload adds the payload to the token info internal server error response
func (o *TokenInfoInternalServerError) WithPayload(payload *models.Error) *TokenInfoInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the token info internal server error response
func (o *TokenInfoInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TokenInfoInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
