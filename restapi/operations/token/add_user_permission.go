// Code generated by go-swagger; DO NOT EDIT.

package token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddUserPermissionHandlerFunc turns a function with the right signature into a add user permission handler
type AddUserPermissionHandlerFunc func(AddUserPermissionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AddUserPermissionHandlerFunc) Handle(params AddUserPermissionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AddUserPermissionHandler interface for that can handle valid add user permission params
type AddUserPermissionHandler interface {
	Handle(AddUserPermissionParams, interface{}) middleware.Responder
}

// NewAddUserPermission creates a new http.Handler for the add user permission operation
func NewAddUserPermission(ctx *middleware.Context, handler AddUserPermissionHandler) *AddUserPermission {
	return &AddUserPermission{Context: ctx, Handler: handler}
}

/*AddUserPermission swagger:route POST /token/authorization/ token addUserPermission

Get all Permissions for a particular token .

Add a permission.

*/
type AddUserPermission struct {
	Context *middleware.Context
	Handler AddUserPermissionHandler
}

func (o *AddUserPermission) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddUserPermissionParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}