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

// GetAliveOKCode is the HTTP code returned for type GetAliveOK
const GetAliveOKCode int = 200

/*GetAliveOK Success

swagger:response getAliveOK
*/
type GetAliveOK struct {
}

// NewGetAliveOK creates GetAliveOK with default headers values
func NewGetAliveOK() *GetAliveOK {

	return &GetAliveOK{}
}

// WriteResponse to the client
func (o *GetAliveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*GetAliveDefault A generic erorr returned by the api

swagger:response getAliveDefault
*/
type GetAliveDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetAliveDefault creates GetAliveDefault with default headers values
func NewGetAliveDefault(code int) *GetAliveDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAliveDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get alive default response
func (o *GetAliveDefault) WithStatusCode(code int) *GetAliveDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get alive default response
func (o *GetAliveDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get alive default response
func (o *GetAliveDefault) WithPayload(payload *models.APIError) *GetAliveDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get alive default response
func (o *GetAliveDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAliveDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
