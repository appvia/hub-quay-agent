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

// DeleteRegistryNamespaceNameOKCode is the HTTP code returned for type DeleteRegistryNamespaceNameOK
const DeleteRegistryNamespaceNameOKCode int = 200

/*DeleteRegistryNamespaceNameOK Successfully deleted the repository from the organization

swagger:response deleteRegistryNamespaceNameOK
*/
type DeleteRegistryNamespaceNameOK struct {
}

// NewDeleteRegistryNamespaceNameOK creates DeleteRegistryNamespaceNameOK with default headers values
func NewDeleteRegistryNamespaceNameOK() *DeleteRegistryNamespaceNameOK {

	return &DeleteRegistryNamespaceNameOK{}
}

// WriteResponse to the client
func (o *DeleteRegistryNamespaceNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*DeleteRegistryNamespaceNameDefault A generic erorr returned by the api

swagger:response deleteRegistryNamespaceNameDefault
*/
type DeleteRegistryNamespaceNameDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewDeleteRegistryNamespaceNameDefault creates DeleteRegistryNamespaceNameDefault with default headers values
func NewDeleteRegistryNamespaceNameDefault(code int) *DeleteRegistryNamespaceNameDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteRegistryNamespaceNameDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete registry namespace name default response
func (o *DeleteRegistryNamespaceNameDefault) WithStatusCode(code int) *DeleteRegistryNamespaceNameDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete registry namespace name default response
func (o *DeleteRegistryNamespaceNameDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete registry namespace name default response
func (o *DeleteRegistryNamespaceNameDefault) WithPayload(payload *models.APIError) *DeleteRegistryNamespaceNameDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete registry namespace name default response
func (o *DeleteRegistryNamespaceNameDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRegistryNamespaceNameDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
