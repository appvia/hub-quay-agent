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

	models "github.com/appvia/hub-quay-agent/pkg/transport/models"
)

// PutTeamsNamespaceNameHandlerFunc turns a function with the right signature into a put teams namespace name handler
type PutTeamsNamespaceNameHandlerFunc func(PutTeamsNamespaceNameParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PutTeamsNamespaceNameHandlerFunc) Handle(params PutTeamsNamespaceNameParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PutTeamsNamespaceNameHandler interface for that can handle valid put teams namespace name params
type PutTeamsNamespaceNameHandler interface {
	Handle(PutTeamsNamespaceNameParams, *models.Principal) middleware.Responder
}

// NewPutTeamsNamespaceName creates a new http.Handler for the put teams namespace name operation
func NewPutTeamsNamespaceName(ctx *middleware.Context, handler PutTeamsNamespaceNameHandler) *PutTeamsNamespaceName {
	return &PutTeamsNamespaceName{Context: ctx, Handler: handler}
}

/*PutTeamsNamespaceName swagger:route PUT /teams/{namespace}/{name} putTeamsNamespaceName

Update the team membership

Used to perform updates or creations of teams in Quay


*/
type PutTeamsNamespaceName struct {
	Context *middleware.Context
	Handler PutTeamsNamespaceNameHandler
}

func (o *PutTeamsNamespaceName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutTeamsNamespaceNameParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
