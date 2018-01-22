// Code generated by go-swagger; DO NOT EDIT.

package login

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetLoginHandlerFunc turns a function with the right signature into a get login handler
type GetLoginHandlerFunc func(GetLoginParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetLoginHandlerFunc) Handle(params GetLoginParams) middleware.Responder {
	return fn(params)
}

// GetLoginHandler interface for that can handle valid get login params
type GetLoginHandler interface {
	Handle(GetLoginParams) middleware.Responder
}

// NewGetLogin creates a new http.Handler for the get login operation
func NewGetLogin(ctx *middleware.Context, handler GetLoginHandler) *GetLogin {
	return &GetLogin{Context: ctx, Handler: handler}
}

/*GetLogin swagger:route POST /cidmlogin login getLogin

Execute action for  Loging handler.

Get Loging handler

*/
type GetLogin struct {
	Context *middleware.Context
	Handler GetLoginHandler
}

func (o *GetLogin) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetLoginParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
