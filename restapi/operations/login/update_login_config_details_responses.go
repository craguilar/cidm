// Code generated by go-swagger; DO NOT EDIT.

package login

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/craguilar/cidm/models"
)

// UpdateLoginConfigDetailsCreatedCode is the HTTP code returned for type UpdateLoginConfigDetailsCreated
const UpdateLoginConfigDetailsCreatedCode int = 201

/*UpdateLoginConfigDetailsCreated Created

swagger:response updateLoginConfigDetailsCreated
*/
type UpdateLoginConfigDetailsCreated struct {
}

// NewUpdateLoginConfigDetailsCreated creates UpdateLoginConfigDetailsCreated with default headers values
func NewUpdateLoginConfigDetailsCreated() *UpdateLoginConfigDetailsCreated {
	return &UpdateLoginConfigDetailsCreated{}
}

// WriteResponse to the client
func (o *UpdateLoginConfigDetailsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
}

// UpdateLoginConfigDetailsUnauthorizedCode is the HTTP code returned for type UpdateLoginConfigDetailsUnauthorized
const UpdateLoginConfigDetailsUnauthorizedCode int = 401

/*UpdateLoginConfigDetailsUnauthorized Authentication information is missing or invalid

swagger:response updateLoginConfigDetailsUnauthorized
*/
type UpdateLoginConfigDetailsUnauthorized struct {
	/*
	  Required: true
	*/
	WWWAuthenticate string `json:"WWW_Authenticate"`
}

// NewUpdateLoginConfigDetailsUnauthorized creates UpdateLoginConfigDetailsUnauthorized with default headers values
func NewUpdateLoginConfigDetailsUnauthorized() *UpdateLoginConfigDetailsUnauthorized {
	return &UpdateLoginConfigDetailsUnauthorized{}
}

// WithWWWAuthenticate adds the wWWAuthenticate to the update login config details unauthorized response
func (o *UpdateLoginConfigDetailsUnauthorized) WithWWWAuthenticate(wWWAuthenticate string) *UpdateLoginConfigDetailsUnauthorized {
	o.WWWAuthenticate = wWWAuthenticate
	return o
}

// SetWWWAuthenticate sets the wWWAuthenticate to the update login config details unauthorized response
func (o *UpdateLoginConfigDetailsUnauthorized) SetWWWAuthenticate(wWWAuthenticate string) {
	o.WWWAuthenticate = wWWAuthenticate
}

// WriteResponse to the client
func (o *UpdateLoginConfigDetailsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header WWW_Authenticate

	wWWAuthenticate := o.WWWAuthenticate
	if wWWAuthenticate != "" {
		rw.Header().Set("WWW_Authenticate", wWWAuthenticate)
	}

	rw.WriteHeader(401)
}

// UpdateLoginConfigDetailsNotFoundCode is the HTTP code returned for type UpdateLoginConfigDetailsNotFound
const UpdateLoginConfigDetailsNotFoundCode int = 404

/*UpdateLoginConfigDetailsNotFound Loging not found

swagger:response updateLoginConfigDetailsNotFound
*/
type UpdateLoginConfigDetailsNotFound struct {
}

// NewUpdateLoginConfigDetailsNotFound creates UpdateLoginConfigDetailsNotFound with default headers values
func NewUpdateLoginConfigDetailsNotFound() *UpdateLoginConfigDetailsNotFound {
	return &UpdateLoginConfigDetailsNotFound{}
}

// WriteResponse to the client
func (o *UpdateLoginConfigDetailsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

/*UpdateLoginConfigDetailsDefault Unexpected error

swagger:response updateLoginConfigDetailsDefault
*/
type UpdateLoginConfigDetailsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateLoginConfigDetailsDefault creates UpdateLoginConfigDetailsDefault with default headers values
func NewUpdateLoginConfigDetailsDefault(code int) *UpdateLoginConfigDetailsDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateLoginConfigDetailsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update login config details default response
func (o *UpdateLoginConfigDetailsDefault) WithStatusCode(code int) *UpdateLoginConfigDetailsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update login config details default response
func (o *UpdateLoginConfigDetailsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update login config details default response
func (o *UpdateLoginConfigDetailsDefault) WithPayload(payload *models.Error) *UpdateLoginConfigDetailsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update login config details default response
func (o *UpdateLoginConfigDetailsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateLoginConfigDetailsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
