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
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/appvia/hub-quay-agent/pkg/transport/models"
)

// GetHealthzNamespaceOKCode is the HTTP code returned for type GetHealthzNamespaceOK
const GetHealthzNamespaceOKCode int = 200

/*GetHealthzNamespaceOK Success

swagger:response getHealthzNamespaceOK
*/
type GetHealthzNamespaceOK struct {
}

// NewGetHealthzNamespaceOK creates GetHealthzNamespaceOK with default headers values
func NewGetHealthzNamespaceOK() *GetHealthzNamespaceOK {

	return &GetHealthzNamespaceOK{}
}

// WriteResponse to the client
func (o *GetHealthzNamespaceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*GetHealthzNamespaceDefault A generic erorr returned by the api

swagger:response getHealthzNamespaceDefault
*/
type GetHealthzNamespaceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetHealthzNamespaceDefault creates GetHealthzNamespaceDefault with default headers values
func NewGetHealthzNamespaceDefault(code int) *GetHealthzNamespaceDefault {
	if code <= 0 {
		code = 500
	}

	return &GetHealthzNamespaceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get healthz namespace default response
func (o *GetHealthzNamespaceDefault) WithStatusCode(code int) *GetHealthzNamespaceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get healthz namespace default response
func (o *GetHealthzNamespaceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get healthz namespace default response
func (o *GetHealthzNamespaceDefault) WithPayload(payload *models.APIError) *GetHealthzNamespaceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get healthz namespace default response
func (o *GetHealthzNamespaceDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHealthzNamespaceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
