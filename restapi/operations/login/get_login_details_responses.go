// Code generated by go-swagger; DO NOT EDIT.

package login

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/craguilar/cidm/models"
)

// GetLoginDetailsOKCode is the HTTP code returned for type GetLoginDetailsOK
const GetLoginDetailsOKCode int = 200

/*GetLoginDetailsOK Successful operation

swagger:response getLoginDetailsOK
*/
type GetLoginDetailsOK struct {

	/*
	  In: Body
	*/
	Payload *models.Login `json:"body,omitempty"`
}

// NewGetLoginDetailsOK creates GetLoginDetailsOK with default headers values
func NewGetLoginDetailsOK() *GetLoginDetailsOK {
	return &GetLoginDetailsOK{}
}

// WithPayload adds the payload to the get login details o k response
func (o *GetLoginDetailsOK) WithPayload(payload *models.Login) *GetLoginDetailsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get login details o k response
func (o *GetLoginDetailsOK) SetPayload(payload *models.Login) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLoginDetailsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetLoginDetailsNotFoundCode is the HTTP code returned for type GetLoginDetailsNotFound
const GetLoginDetailsNotFoundCode int = 404

/*GetLoginDetailsNotFound Loging not found

swagger:response getLoginDetailsNotFound
*/
type GetLoginDetailsNotFound struct {
}

// NewGetLoginDetailsNotFound creates GetLoginDetailsNotFound with default headers values
func NewGetLoginDetailsNotFound() *GetLoginDetailsNotFound {
	return &GetLoginDetailsNotFound{}
}

// WriteResponse to the client
func (o *GetLoginDetailsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

/*GetLoginDetailsDefault Unexpected error

swagger:response getLoginDetailsDefault
*/
type GetLoginDetailsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLoginDetailsDefault creates GetLoginDetailsDefault with default headers values
func NewGetLoginDetailsDefault(code int) *GetLoginDetailsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetLoginDetailsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get login details default response
func (o *GetLoginDetailsDefault) WithStatusCode(code int) *GetLoginDetailsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get login details default response
func (o *GetLoginDetailsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get login details default response
func (o *GetLoginDetailsDefault) WithPayload(payload *models.Error) *GetLoginDetailsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get login details default response
func (o *GetLoginDetailsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLoginDetailsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
