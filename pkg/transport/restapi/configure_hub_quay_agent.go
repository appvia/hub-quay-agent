// This file is safe to edit. Once it exists it will not be overwritten

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

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/appvia/hub-quay-agent/pkg/transport/restapi/operations"
	"github.com/appvia/hub-quay-agent/pkg/transport/restapi/operations/health"
)

//go:generate swagger generate server --target ../../transport --name HubQuayAgent --spec ../../../swagger.yml --exclude-main

func configureFlags(api *operations.HubQuayAgentAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.HubQuayAgentAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.DeleteRegistryNamespaceNameHandler = operations.DeleteRegistryNamespaceNameHandlerFunc(func(params operations.DeleteRegistryNamespaceNameParams) middleware.Responder {
		return middleware.NotImplemented("operation .DeleteRegistryNamespaceName has not yet been implemented")
	})
	api.DeleteRobotsNamespaceNameHandler = operations.DeleteRobotsNamespaceNameHandlerFunc(func(params operations.DeleteRobotsNamespaceNameParams) middleware.Responder {
		return middleware.NotImplemented("operation .DeleteRobotsNamespaceName has not yet been implemented")
	})
	api.HealthGetHealthzHandler = health.GetHealthzHandlerFunc(func(params health.GetHealthzParams) middleware.Responder {
		return middleware.NotImplemented("operation health.GetHealthz has not yet been implemented")
	})
	api.GetRegistryNamespaceHandler = operations.GetRegistryNamespaceHandlerFunc(func(params operations.GetRegistryNamespaceParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetRegistryNamespace has not yet been implemented")
	})
	api.GetRegistryNamespaceNameHandler = operations.GetRegistryNamespaceNameHandlerFunc(func(params operations.GetRegistryNamespaceNameParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetRegistryNamespaceName has not yet been implemented")
	})
	api.GetRobotsNamespaceHandler = operations.GetRobotsNamespaceHandlerFunc(func(params operations.GetRobotsNamespaceParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetRobotsNamespace has not yet been implemented")
	})
	api.GetRobotsNamespaceNameHandler = operations.GetRobotsNamespaceNameHandlerFunc(func(params operations.GetRobotsNamespaceNameParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetRobotsNamespaceName has not yet been implemented")
	})
	api.PostRegistryNamespaceNameHandler = operations.PostRegistryNamespaceNameHandlerFunc(func(params operations.PostRegistryNamespaceNameParams) middleware.Responder {
		return middleware.NotImplemented("operation .PostRegistryNamespaceName has not yet been implemented")
	})
	api.PostRobotsNamespaceNameHandler = operations.PostRobotsNamespaceNameHandlerFunc(func(params operations.PostRobotsNamespaceNameParams) middleware.Responder {
		return middleware.NotImplemented("operation .PostRobotsNamespaceName has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
