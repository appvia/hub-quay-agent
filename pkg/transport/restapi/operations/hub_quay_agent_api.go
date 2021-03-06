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
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	models "github.com/appvia/hub-quay-agent/pkg/transport/models"
)

// NewHubQuayAgentAPI creates a new HubQuayAgent instance
func NewHubQuayAgentAPI(spec *loads.Document) *HubQuayAgentAPI {
	return &HubQuayAgentAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		DeleteRegistryNamespaceNameHandler: DeleteRegistryNamespaceNameHandlerFunc(func(params DeleteRegistryNamespaceNameParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation DeleteRegistryNamespaceName has not yet been implemented")
		}),
		DeleteRobotsNamespaceNameHandler: DeleteRobotsNamespaceNameHandlerFunc(func(params DeleteRobotsNamespaceNameParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation DeleteRobotsNamespaceName has not yet been implemented")
		}),
		DeleteTeamsNamespaceNameHandler: DeleteTeamsNamespaceNameHandlerFunc(func(params DeleteTeamsNamespaceNameParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation DeleteTeamsNamespaceName has not yet been implemented")
		}),
		GetAliveHandler: GetAliveHandlerFunc(func(params GetAliveParams) middleware.Responder {
			return middleware.NotImplemented("operation GetAlive has not yet been implemented")
		}),
		GetHealthzNamespaceHandler: GetHealthzNamespaceHandlerFunc(func(params GetHealthzNamespaceParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetHealthzNamespace has not yet been implemented")
		}),
		GetRegistryNamespaceHandler: GetRegistryNamespaceHandlerFunc(func(params GetRegistryNamespaceParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetRegistryNamespace has not yet been implemented")
		}),
		GetRegistryNamespaceNameHandler: GetRegistryNamespaceNameHandlerFunc(func(params GetRegistryNamespaceNameParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetRegistryNamespaceName has not yet been implemented")
		}),
		GetRegistryNamespaceNameStatusHandler: GetRegistryNamespaceNameStatusHandlerFunc(func(params GetRegistryNamespaceNameStatusParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetRegistryNamespaceNameStatus has not yet been implemented")
		}),
		GetRobotsNamespaceHandler: GetRobotsNamespaceHandlerFunc(func(params GetRobotsNamespaceParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetRobotsNamespace has not yet been implemented")
		}),
		GetRobotsNamespaceNameHandler: GetRobotsNamespaceNameHandlerFunc(func(params GetRobotsNamespaceNameParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetRobotsNamespaceName has not yet been implemented")
		}),
		GetTeamsNamespaceHandler: GetTeamsNamespaceHandlerFunc(func(params GetTeamsNamespaceParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetTeamsNamespace has not yet been implemented")
		}),
		GetTeamsNamespaceNameHandler: GetTeamsNamespaceNameHandlerFunc(func(params GetTeamsNamespaceNameParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetTeamsNamespaceName has not yet been implemented")
		}),
		PutRegistryNamespaceNameHandler: PutRegistryNamespaceNameHandlerFunc(func(params PutRegistryNamespaceNameParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation PutRegistryNamespaceName has not yet been implemented")
		}),
		PutRobotsNamespaceNameHandler: PutRobotsNamespaceNameHandlerFunc(func(params PutRobotsNamespaceNameParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation PutRobotsNamespaceName has not yet been implemented")
		}),
		PutTeamsNamespaceNameHandler: PutTeamsNamespaceNameHandlerFunc(func(params PutTeamsNamespaceNameParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation PutTeamsNamespaceName has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		ApikeyAuth: func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (apikey) Authorization from header param [Authorization] has not yet been implemented")
		},

		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*HubQuayAgentAPI an agent used to provision and configure repositories in quay */
type HubQuayAgentAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// ApikeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	ApikeyAuth func(string) (*models.Principal, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// DeleteRegistryNamespaceNameHandler sets the operation handler for the delete registry namespace name operation
	DeleteRegistryNamespaceNameHandler DeleteRegistryNamespaceNameHandler
	// DeleteRobotsNamespaceNameHandler sets the operation handler for the delete robots namespace name operation
	DeleteRobotsNamespaceNameHandler DeleteRobotsNamespaceNameHandler
	// DeleteTeamsNamespaceNameHandler sets the operation handler for the delete teams namespace name operation
	DeleteTeamsNamespaceNameHandler DeleteTeamsNamespaceNameHandler
	// GetAliveHandler sets the operation handler for the get alive operation
	GetAliveHandler GetAliveHandler
	// GetHealthzNamespaceHandler sets the operation handler for the get healthz namespace operation
	GetHealthzNamespaceHandler GetHealthzNamespaceHandler
	// GetRegistryNamespaceHandler sets the operation handler for the get registry namespace operation
	GetRegistryNamespaceHandler GetRegistryNamespaceHandler
	// GetRegistryNamespaceNameHandler sets the operation handler for the get registry namespace name operation
	GetRegistryNamespaceNameHandler GetRegistryNamespaceNameHandler
	// GetRegistryNamespaceNameStatusHandler sets the operation handler for the get registry namespace name status operation
	GetRegistryNamespaceNameStatusHandler GetRegistryNamespaceNameStatusHandler
	// GetRobotsNamespaceHandler sets the operation handler for the get robots namespace operation
	GetRobotsNamespaceHandler GetRobotsNamespaceHandler
	// GetRobotsNamespaceNameHandler sets the operation handler for the get robots namespace name operation
	GetRobotsNamespaceNameHandler GetRobotsNamespaceNameHandler
	// GetTeamsNamespaceHandler sets the operation handler for the get teams namespace operation
	GetTeamsNamespaceHandler GetTeamsNamespaceHandler
	// GetTeamsNamespaceNameHandler sets the operation handler for the get teams namespace name operation
	GetTeamsNamespaceNameHandler GetTeamsNamespaceNameHandler
	// PutRegistryNamespaceNameHandler sets the operation handler for the put registry namespace name operation
	PutRegistryNamespaceNameHandler PutRegistryNamespaceNameHandler
	// PutRobotsNamespaceNameHandler sets the operation handler for the put robots namespace name operation
	PutRobotsNamespaceNameHandler PutRobotsNamespaceNameHandler
	// PutTeamsNamespaceNameHandler sets the operation handler for the put teams namespace name operation
	PutTeamsNamespaceNameHandler PutTeamsNamespaceNameHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *HubQuayAgentAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *HubQuayAgentAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *HubQuayAgentAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *HubQuayAgentAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *HubQuayAgentAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *HubQuayAgentAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *HubQuayAgentAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the HubQuayAgentAPI
func (o *HubQuayAgentAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.ApikeyAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.DeleteRegistryNamespaceNameHandler == nil {
		unregistered = append(unregistered, "DeleteRegistryNamespaceNameHandler")
	}

	if o.DeleteRobotsNamespaceNameHandler == nil {
		unregistered = append(unregistered, "DeleteRobotsNamespaceNameHandler")
	}

	if o.DeleteTeamsNamespaceNameHandler == nil {
		unregistered = append(unregistered, "DeleteTeamsNamespaceNameHandler")
	}

	if o.GetAliveHandler == nil {
		unregistered = append(unregistered, "GetAliveHandler")
	}

	if o.GetHealthzNamespaceHandler == nil {
		unregistered = append(unregistered, "GetHealthzNamespaceHandler")
	}

	if o.GetRegistryNamespaceHandler == nil {
		unregistered = append(unregistered, "GetRegistryNamespaceHandler")
	}

	if o.GetRegistryNamespaceNameHandler == nil {
		unregistered = append(unregistered, "GetRegistryNamespaceNameHandler")
	}

	if o.GetRegistryNamespaceNameStatusHandler == nil {
		unregistered = append(unregistered, "GetRegistryNamespaceNameStatusHandler")
	}

	if o.GetRobotsNamespaceHandler == nil {
		unregistered = append(unregistered, "GetRobotsNamespaceHandler")
	}

	if o.GetRobotsNamespaceNameHandler == nil {
		unregistered = append(unregistered, "GetRobotsNamespaceNameHandler")
	}

	if o.GetTeamsNamespaceHandler == nil {
		unregistered = append(unregistered, "GetTeamsNamespaceHandler")
	}

	if o.GetTeamsNamespaceNameHandler == nil {
		unregistered = append(unregistered, "GetTeamsNamespaceNameHandler")
	}

	if o.PutRegistryNamespaceNameHandler == nil {
		unregistered = append(unregistered, "PutRegistryNamespaceNameHandler")
	}

	if o.PutRobotsNamespaceNameHandler == nil {
		unregistered = append(unregistered, "PutRobotsNamespaceNameHandler")
	}

	if o.PutTeamsNamespaceNameHandler == nil {
		unregistered = append(unregistered, "PutTeamsNamespaceNameHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *HubQuayAgentAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *HubQuayAgentAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "apikey":

			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.ApikeyAuth(token)
			})

		}
	}
	return result

}

// Authorizer returns the registered authorizer
func (o *HubQuayAgentAPI) Authorizer() runtime.Authorizer {

	return o.APIAuthorizer

}

// ConsumersFor gets the consumers for the specified media types
func (o *HubQuayAgentAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *HubQuayAgentAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *HubQuayAgentAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the hub quay agent API
func (o *HubQuayAgentAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *HubQuayAgentAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/registry/{namespace}/{name}"] = NewDeleteRegistryNamespaceName(o.context, o.DeleteRegistryNamespaceNameHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/robots/{namespace}/{name}"] = NewDeleteRobotsNamespaceName(o.context, o.DeleteRobotsNamespaceNameHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/teams/{namespace}/{name}"] = NewDeleteTeamsNamespaceName(o.context, o.DeleteTeamsNamespaceNameHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/alive"] = NewGetAlive(o.context, o.GetAliveHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/healthz/{namespace}"] = NewGetHealthzNamespace(o.context, o.GetHealthzNamespaceHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/registry/{namespace}"] = NewGetRegistryNamespace(o.context, o.GetRegistryNamespaceHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/registry/{namespace}/{name}"] = NewGetRegistryNamespaceName(o.context, o.GetRegistryNamespaceNameHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/registry/{namespace}/{name}/status"] = NewGetRegistryNamespaceNameStatus(o.context, o.GetRegistryNamespaceNameStatusHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/robots/{namespace}"] = NewGetRobotsNamespace(o.context, o.GetRobotsNamespaceHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/robots/{namespace}/{name}"] = NewGetRobotsNamespaceName(o.context, o.GetRobotsNamespaceNameHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/teams/{namespace}"] = NewGetTeamsNamespace(o.context, o.GetTeamsNamespaceHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/teams/{namespace}/{name}"] = NewGetTeamsNamespaceName(o.context, o.GetTeamsNamespaceNameHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/registry/{namespace}/{name}"] = NewPutRegistryNamespaceName(o.context, o.PutRegistryNamespaceNameHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/robots/{namespace}/{name}"] = NewPutRobotsNamespaceName(o.context, o.PutRobotsNamespaceNameHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/teams/{namespace}/{name}"] = NewPutTeamsNamespaceName(o.context, o.PutTeamsNamespaceNameHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *HubQuayAgentAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *HubQuayAgentAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *HubQuayAgentAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *HubQuayAgentAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
