// Code generated by go-swagger; DO NOT EDIT.

package token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/craguilar/cidm/models"
)

// UpdateUserPermissionOKCode is the HTTP code returned for type UpdateUserPermissionOK
const UpdateUserPermissionOKCode int = 200

/*UpdateUserPermissionOK successful operation

swagger:response updateUserPermissionOK
*/
type UpdateUserPermissionOK struct {
}

// NewUpdateUserPermissionOK creates UpdateUserPermissionOK with default headers values
func NewUpdateUserPermissionOK() *UpdateUserPermissionOK {
	return &UpdateUserPermissionOK{}
}

// WriteResponse to the client
func (o *UpdateUserPermissionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

// UpdateUserPermissionUnauthorizedCode is the HTTP code returned for type UpdateUserPermissionUnauthorized
const UpdateUserPermissionUnauthorizedCode int = 401

/*UpdateUserPermissionUnauthorized Authentication information is missing or invalid

swagger:response updateUserPermissionUnauthorized
*/
type UpdateUserPermissionUnauthorized struct {
	/*
	  Required: true
	*/
	WWWAuthenticate string `json:"WWW_Authenticate"`
}

// NewUpdateUserPermissionUnauthorized creates UpdateUserPermissionUnauthorized with default headers values
func NewUpdateUserPermissionUnauthorized() *UpdateUserPermissionUnauthorized {
	return &UpdateUserPermissionUnauthorized{}
}

// WithWWWAuthenticate adds the wWWAuthenticate to the update user permission unauthorized response
func (o *UpdateUserPermissionUnauthorized) WithWWWAuthenticate(wWWAuthenticate string) *UpdateUserPermissionUnauthorized {
	o.WWWAuthenticate = wWWAuthenticate
	return o
}

// SetWWWAuthenticate sets the wWWAuthenticate to the update user permission unauthorized response
func (o *UpdateUserPermissionUnauthorized) SetWWWAuthenticate(wWWAuthenticate string) {
	o.WWWAuthenticate = wWWAuthenticate
}

// WriteResponse to the client
func (o *UpdateUserPermissionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header WWW_Authenticate

	wWWAuthenticate := o.WWWAuthenticate
	if wWWAuthenticate != "" {
		rw.Header().Set("WWW_Authenticate", wWWAuthenticate)
	}

	rw.WriteHeader(401)
}

/*UpdateUserPermissionDefault Unexpected error

swagger:response updateUserPermissionDefault
*/
type UpdateUserPermissionDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateUserPermissionDefault creates UpdateUserPermissionDefault with default headers values
func NewUpdateUserPermissionDefault(code int) *UpdateUserPermissionDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateUserPermissionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update user permission default response
func (o *UpdateUserPermissionDefault) WithStatusCode(code int) *UpdateUserPermissionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update user permission default response
func (o *UpdateUserPermissionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update user permission default response
func (o *UpdateUserPermissionDefault) WithPayload(payload *models.Error) *UpdateUserPermissionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user permission default response
func (o *UpdateUserPermissionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserPermissionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
