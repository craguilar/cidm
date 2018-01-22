// Code generated by go-swagger; DO NOT EDIT.

package token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateUserPermissionHandlerFunc turns a function with the right signature into a update user permission handler
type UpdateUserPermissionHandlerFunc func(UpdateUserPermissionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateUserPermissionHandlerFunc) Handle(params UpdateUserPermissionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateUserPermissionHandler interface for that can handle valid update user permission params
type UpdateUserPermissionHandler interface {
	Handle(UpdateUserPermissionParams, interface{}) middleware.Responder
}

// NewUpdateUserPermission creates a new http.Handler for the update user permission operation
func NewUpdateUserPermission(ctx *middleware.Context, handler UpdateUserPermissionHandler) *UpdateUserPermission {
	return &UpdateUserPermission{Context: ctx, Handler: handler}
}

/*UpdateUserPermission swagger:route PUT /token/authorization/ token updateUserPermission

Get all Permissions for a particular token .

Update a permission.

*/
type UpdateUserPermission struct {
	Context *middleware.Context
	Handler UpdateUserPermissionHandler
}

func (o *UpdateUserPermission) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateUserPermissionParams()

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