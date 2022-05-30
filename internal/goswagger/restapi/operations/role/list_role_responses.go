// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

// ListRoleOKCode is the HTTP code returned for type ListRoleOK
const ListRoleOKCode int = 200

/*ListRoleOK OK

swagger:response listRoleOK
*/
type ListRoleOK struct {

	/*
	  In: Body
	*/
	Payload []*ent.Role `json:"body,omitempty"`
}

// NewListRoleOK creates ListRoleOK with default headers values
func NewListRoleOK() *ListRoleOK {

	return &ListRoleOK{}
}

// WithPayload adds the payload to the list role o k response
func (o *ListRoleOK) WithPayload(payload []*ent.Role) *ListRoleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list role o k response
func (o *ListRoleOK) SetPayload(payload []*ent.Role) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListRoleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*ent.Role, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListRoleUnauthorizedCode is the HTTP code returned for type ListRoleUnauthorized
const ListRoleUnauthorizedCode int = 401

/*ListRoleUnauthorized Unauthorized

swagger:response listRoleUnauthorized
*/
type ListRoleUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListRoleUnauthorized creates ListRoleUnauthorized with default headers values
func NewListRoleUnauthorized() *ListRoleUnauthorized {

	return &ListRoleUnauthorized{}
}

// WithPayload adds the payload to the list role unauthorized response
func (o *ListRoleUnauthorized) WithPayload(payload *models.Error) *ListRoleUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list role unauthorized response
func (o *ListRoleUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListRoleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListRoleForbiddenCode is the HTTP code returned for type ListRoleForbidden
const ListRoleForbiddenCode int = 403

/*ListRoleForbidden Unauthorized

swagger:response listRoleForbidden
*/
type ListRoleForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListRoleForbidden creates ListRoleForbidden with default headers values
func NewListRoleForbidden() *ListRoleForbidden {

	return &ListRoleForbidden{}
}

// WithPayload adds the payload to the list role forbidden response
func (o *ListRoleForbidden) WithPayload(payload *models.Error) *ListRoleForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list role forbidden response
func (o *ListRoleForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListRoleForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListRoleInternalServerErrorCode is the HTTP code returned for type ListRoleInternalServerError
const ListRoleInternalServerErrorCode int = 500

/*ListRoleInternalServerError Internal Server Error

swagger:response listRoleInternalServerError
*/
type ListRoleInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListRoleInternalServerError creates ListRoleInternalServerError with default headers values
func NewListRoleInternalServerError() *ListRoleInternalServerError {

	return &ListRoleInternalServerError{}
}

// WithPayload adds the payload to the list role internal server error response
func (o *ListRoleInternalServerError) WithPayload(payload *models.Error) *ListRoleInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list role internal server error response
func (o *ListRoleInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListRoleInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
