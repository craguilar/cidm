// Code generated by go-swagger; DO NOT EDIT.

package login

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetLoginDetailsByIDHandlerFunc turns a function with the right signature into a get login details by Id handler
type GetLoginDetailsByIDHandlerFunc func(GetLoginDetailsByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetLoginDetailsByIDHandlerFunc) Handle(params GetLoginDetailsByIDParams) middleware.Responder {
	return fn(params)
}

// GetLoginDetailsByIDHandler interface for that can handle valid get login details by Id params
type GetLoginDetailsByIDHandler interface {
	Handle(GetLoginDetailsByIDParams) middleware.Responder
}

// NewGetLoginDetailsByID creates a new http.Handler for the get login details by Id operation
func NewGetLoginDetailsByID(ctx *middleware.Context, handler GetLoginDetailsByIDHandler) *GetLoginDetailsByID {
	return &GetLoginDetailsByID{Context: ctx, Handler: handler}
}

/*GetLoginDetailsByID swagger:route GET /login/{id} login getLoginDetailsById

Get Login details by Id to display in login screen.

Get Loging details it returns an object which contains basic info to
initiate the login this is a non secured way this allows the user to get
all the login details available.


*/
type GetLoginDetailsByID struct {
	Context *middleware.Context
	Handler GetLoginDetailsByIDHandler
}

func (o *GetLoginDetailsByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetLoginDetailsByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
