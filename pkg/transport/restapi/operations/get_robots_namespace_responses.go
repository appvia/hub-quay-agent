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

// GetRobotsNamespaceOKCode is the HTTP code returned for type GetRobotsNamespaceOK
const GetRobotsNamespaceOKCode int = 200

/*GetRobotsNamespaceOK Returning a list of robot accounts

swagger:response getRobotsNamespaceOK
*/
type GetRobotsNamespaceOK struct {

	/*
	  In: Body
	*/
	Payload *models.RobotList `json:"body,omitempty"`
}

// NewGetRobotsNamespaceOK creates GetRobotsNamespaceOK with default headers values
func NewGetRobotsNamespaceOK() *GetRobotsNamespaceOK {

	return &GetRobotsNamespaceOK{}
}

// WithPayload adds the payload to the get robots namespace o k response
func (o *GetRobotsNamespaceOK) WithPayload(payload *models.RobotList) *GetRobotsNamespaceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get robots namespace o k response
func (o *GetRobotsNamespaceOK) SetPayload(payload *models.RobotList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRobotsNamespaceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetRobotsNamespaceDefault A generic erorr returned by the api

swagger:response getRobotsNamespaceDefault
*/
type GetRobotsNamespaceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetRobotsNamespaceDefault creates GetRobotsNamespaceDefault with default headers values
func NewGetRobotsNamespaceDefault(code int) *GetRobotsNamespaceDefault {
	if code <= 0 {
		code = 500
	}

	return &GetRobotsNamespaceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get robots namespace default response
func (o *GetRobotsNamespaceDefault) WithStatusCode(code int) *GetRobotsNamespaceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get robots namespace default response
func (o *GetRobotsNamespaceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get robots namespace default response
func (o *GetRobotsNamespaceDefault) WithPayload(payload *models.APIError) *GetRobotsNamespaceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get robots namespace default response
func (o *GetRobotsNamespaceDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRobotsNamespaceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
