// Code generated by go-swagger; DO NOT EDIT.

package token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// TokenValidateHandlerFunc turns a function with the right signature into a token validate handler
type TokenValidateHandlerFunc func(TokenValidateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn TokenValidateHandlerFunc) Handle(params TokenValidateParams) middleware.Responder {
	return fn(params)
}

// TokenValidateHandler interface for that can handle valid token validate params
type TokenValidateHandler interface {
	Handle(TokenValidateParams) middleware.Responder
}

// NewTokenValidate creates a new http.Handler for the token validate operation
func NewTokenValidate(ctx *middleware.Context, handler TokenValidateHandler) *TokenValidate {
	return &TokenValidate{Context: ctx, Handler: handler}
}

/*TokenValidate swagger:route GET /token/authentication token tokenValidate

Validate token in header.

Get Loging screen

*/
type TokenValidate struct {
	Context *middleware.Context
	Handler TokenValidateHandler
}

func (o *TokenValidate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewTokenValidateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
