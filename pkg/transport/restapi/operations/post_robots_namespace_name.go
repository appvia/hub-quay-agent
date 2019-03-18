// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright (C) 2019  Rohith Jayawardene <gambol99@gmail.com>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
//

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostRobotsNamespaceNameHandlerFunc turns a function with the right signature into a post robots namespace name handler
type PostRobotsNamespaceNameHandlerFunc func(PostRobotsNamespaceNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostRobotsNamespaceNameHandlerFunc) Handle(params PostRobotsNamespaceNameParams) middleware.Responder {
	return fn(params)
}

// PostRobotsNamespaceNameHandler interface for that can handle valid post robots namespace name params
type PostRobotsNamespaceNameHandler interface {
	Handle(PostRobotsNamespaceNameParams) middleware.Responder
}

// NewPostRobotsNamespaceName creates a new http.Handler for the post robots namespace name operation
func NewPostRobotsNamespaceName(ctx *middleware.Context, handler PostRobotsNamespaceNameHandler) *PostRobotsNamespaceName {
	return &PostRobotsNamespaceName{Context: ctx, Handler: handler}
}

/*PostRobotsNamespaceName swagger:route POST /robots/{namespace}/{name} postRobotsNamespaceName

Retrieves a list of robot accounts from within the registry

Used to retrieve a list of robot accounts and the permissions they
have on the repositories


*/
type PostRobotsNamespaceName struct {
	Context *middleware.Context
	Handler PostRobotsNamespaceNameHandler
}

func (o *PostRobotsNamespaceName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostRobotsNamespaceNameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
