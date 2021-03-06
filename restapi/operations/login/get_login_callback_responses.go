// Code generated by go-swagger; DO NOT EDIT.

package login

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/craguilar/cidm/models"
)

// GetLoginCallbackFoundCode is the HTTP code returned for type GetLoginCallbackFound
const GetLoginCallbackFoundCode int = 302

/*GetLoginCallbackFound Successful operation ok - redirect to OAuth2

swagger:response getLoginCallbackFound
*/
type GetLoginCallbackFound struct {
	/*
	  Required: true
	*/
	Authorization string `json:"Authorization"`
	/*
	  Required: true
	*/
	Location string `json:"Location"`
}

// NewGetLoginCallbackFound creates GetLoginCallbackFound with default headers values
func NewGetLoginCallbackFound() *GetLoginCallbackFound {
	return &GetLoginCallbackFound{}
}

// WithAuthorization adds the authorization to the get login callback found response
func (o *GetLoginCallbackFound) WithAuthorization(authorization string) *GetLoginCallbackFound {
	o.Authorization = authorization
	return o
}

// SetAuthorization sets the authorization to the get login callback found response
func (o *GetLoginCallbackFound) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithLocation adds the location to the get login callback found response
func (o *GetLoginCallbackFound) WithLocation(location string) *GetLoginCallbackFound {
	o.Location = location
	return o
}

// SetLocation sets the location to the get login callback found response
func (o *GetLoginCallbackFound) SetLocation(location string) {
	o.Location = location
}

// WriteResponse to the client
func (o *GetLoginCallbackFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Authorization

	authorization := o.Authorization
	if authorization != "" {
		rw.Header().Set("Authorization", authorization)
	}

	// response header Location

	location := o.Location
	if location != "" {
		rw.Header().Set("Location", location)
	}

	rw.WriteHeader(302)
}

/*GetLoginCallbackDefault Unexpected error

swagger:response getLoginCallbackDefault
*/
type GetLoginCallbackDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLoginCallbackDefault creates GetLoginCallbackDefault with default headers values
func NewGetLoginCallbackDefault(code int) *GetLoginCallbackDefault {
	if code <= 0 {
		code = 500
	}

	return &GetLoginCallbackDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get login callback default response
func (o *GetLoginCallbackDefault) WithStatusCode(code int) *GetLoginCallbackDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get login callback default response
func (o *GetLoginCallbackDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get login callback default response
func (o *GetLoginCallbackDefault) WithPayload(payload *models.Error) *GetLoginCallbackDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get login callback default response
func (o *GetLoginCallbackDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLoginCallbackDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
