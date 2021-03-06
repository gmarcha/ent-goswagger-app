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

// ListMeRolesOKCode is the HTTP code returned for type ListMeRolesOK
const ListMeRolesOKCode int = 200

/*ListMeRolesOK OK

swagger:response listMeRolesOK
*/
type ListMeRolesOK struct {

	/*
	  In: Body
	*/
	Payload []*ent.Role `json:"body,omitempty"`
}

// NewListMeRolesOK creates ListMeRolesOK with default headers values
func NewListMeRolesOK() *ListMeRolesOK {

	return &ListMeRolesOK{}
}

// WithPayload adds the payload to the list me roles o k response
func (o *ListMeRolesOK) WithPayload(payload []*ent.Role) *ListMeRolesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list me roles o k response
func (o *ListMeRolesOK) SetPayload(payload []*ent.Role) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListMeRolesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ListMeRolesUnauthorizedCode is the HTTP code returned for type ListMeRolesUnauthorized
const ListMeRolesUnauthorizedCode int = 401

/*ListMeRolesUnauthorized Unauthorized

swagger:response listMeRolesUnauthorized
*/
type ListMeRolesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListMeRolesUnauthorized creates ListMeRolesUnauthorized with default headers values
func NewListMeRolesUnauthorized() *ListMeRolesUnauthorized {

	return &ListMeRolesUnauthorized{}
}

// WithPayload adds the payload to the list me roles unauthorized response
func (o *ListMeRolesUnauthorized) WithPayload(payload *models.Error) *ListMeRolesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list me roles unauthorized response
func (o *ListMeRolesUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListMeRolesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListMeRolesForbiddenCode is the HTTP code returned for type ListMeRolesForbidden
const ListMeRolesForbiddenCode int = 403

/*ListMeRolesForbidden Forbidden

swagger:response listMeRolesForbidden
*/
type ListMeRolesForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListMeRolesForbidden creates ListMeRolesForbidden with default headers values
func NewListMeRolesForbidden() *ListMeRolesForbidden {

	return &ListMeRolesForbidden{}
}

// WithPayload adds the payload to the list me roles forbidden response
func (o *ListMeRolesForbidden) WithPayload(payload *models.Error) *ListMeRolesForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list me roles forbidden response
func (o *ListMeRolesForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListMeRolesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListMeRolesInternalServerErrorCode is the HTTP code returned for type ListMeRolesInternalServerError
const ListMeRolesInternalServerErrorCode int = 500

/*ListMeRolesInternalServerError Internal Server Error

swagger:response listMeRolesInternalServerError
*/
type ListMeRolesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListMeRolesInternalServerError creates ListMeRolesInternalServerError with default headers values
func NewListMeRolesInternalServerError() *ListMeRolesInternalServerError {

	return &ListMeRolesInternalServerError{}
}

// WithPayload adds the payload to the list me roles internal server error response
func (o *ListMeRolesInternalServerError) WithPayload(payload *models.Error) *ListMeRolesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list me roles internal server error response
func (o *ListMeRolesInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListMeRolesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
