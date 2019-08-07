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

// GetHealthzOKCode is the HTTP code returned for type GetHealthzOK
const GetHealthzOKCode int = 200

/*GetHealthzOK Success

swagger:response getHealthzOK
*/
type GetHealthzOK struct {
}

// NewGetHealthzOK creates GetHealthzOK with default headers values
func NewGetHealthzOK() *GetHealthzOK {

	return &GetHealthzOK{}
}

// WriteResponse to the client
func (o *GetHealthzOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*GetHealthzDefault A generic erorr returned by the api

swagger:response getHealthzDefault
*/
type GetHealthzDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetHealthzDefault creates GetHealthzDefault with default headers values
func NewGetHealthzDefault(code int) *GetHealthzDefault {
	if code <= 0 {
		code = 500
	}

	return &GetHealthzDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get healthz default response
func (o *GetHealthzDefault) WithStatusCode(code int) *GetHealthzDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get healthz default response
func (o *GetHealthzDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get healthz default response
func (o *GetHealthzDefault) WithPayload(payload *models.APIError) *GetHealthzDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get healthz default response
func (o *GetHealthzDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHealthzDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}