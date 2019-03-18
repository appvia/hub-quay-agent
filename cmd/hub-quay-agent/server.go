/*
 * Copyright (C) 2019  Rohith Jayawardene <gambol99@gmail.com>
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/appvia/hub-quay-agent/pkg/server"
	"github.com/appvia/hub-quay-agent/pkg/server/middleware/authinfo"
	"github.com/appvia/hub-quay-agent/pkg/server/middleware/keyauth"
	"github.com/appvia/hub-quay-agent/pkg/transport/restapi"
	"github.com/appvia/hub-quay-agent/pkg/transport/restapi/operations"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	flags "github.com/jessevdk/go-flags"
	"github.com/justinas/alice"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// invokeServerAction handles the command line server action
func invokeServerAction(ctx *cli.Context) error {
	// @step: validate the options

	// @step: create the agent service
	svc, err := server.New(&server.Options{
		AccessToken: ctx.String("quay-api-token"),
		HostnameAPI: ctx.String("quay-endpoint-url"),
	})
	if err != nil {
		return err
	}

	// @step: load the swagger api spec
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return err
	}

	// @step: create the transport api
	api := operations.NewHubQuayAgentAPI(swaggerSpec)

	// @step: configure the restapi
	s := restapi.NewServer(api)
	s.EnabledListeners = []string{"http"}
	s.Host = ctx.String("listen")
	s.Port = ctx.Int("http-port")
	s.TLSPort = ctx.Int("https-port")
	s.TLSHost = s.Host
	if ctx.String("tls-cert") != "" && ctx.String("tls-key") != "" {
		s.EnabledListeners = []string{"https"}
		s.TLSCertificate = flags.Filename(ctx.String("tls-cert"))
		s.TLSCertificateKey = flags.Filename(ctx.String("tls-key"))
	}
	defer s.Shutdown()

	api.DeleteRegistryNamespaceNameHandler = operations.DeleteRegistryNamespaceNameHandlerFunc(func(params operations.DeleteRegistryNamespaceNameParams) middleware.Responder {
		if err := svc.Delete(params.HTTPRequest.Context(), params.Namespace, params.Name); err != nil {
			return operations.NewDeleteRegistryNamespaceNameDefault(http.StatusServiceUnavailable).WithPayload(err)
		}
		return operations.NewDeleteRegistryNamespaceNameOK()
	})

	api.DeleteRobotsNamespaceNameHandler = operations.DeleteRobotsNamespaceNameHandlerFunc(func(params operations.DeleteRobotsNamespaceNameParams) middleware.Responder {
		if err := svc.DeleteRobot(params.HTTPRequest.Context(), params.Namespace, params.Name); err != nil {
			return operations.NewDeleteRobotsNamespaceNameDefault(http.StatusServiceUnavailable).WithPayload(err)
		}
		return operations.NewDeleteRobotsNamespaceNameOK()
	})

	api.GetRegistryNamespaceHandler = operations.GetRegistryNamespaceHandlerFunc(func(params operations.GetRegistryNamespaceParams) middleware.Responder {
		resp, err := svc.List(params.HTTPRequest.Context(), params.Namespace)
		if err != nil {
			return operations.NewGetRegistryNamespaceNameDefault(http.StatusServiceUnavailable).WithPayload(err)
		}
		return operations.NewGetRegistryNamespaceOK().WithPayload(resp)
	})

	api.GetRegistryNamespaceNameHandler = operations.GetRegistryNamespaceNameHandlerFunc(func(params operations.GetRegistryNamespaceNameParams) middleware.Responder {
		resp, err := svc.Get(params.HTTPRequest.Context(), params.Namespace, params.Name)
		if err != nil {
			return operations.NewGetRegistryNamespaceNameDefault(http.StatusServiceUnavailable).WithPayload(err)
		}
		return operations.NewGetRegistryNamespaceNameOK().WithPayload(resp)
	})

	api.GetRobotsNamespaceHandler = operations.GetRobotsNamespaceHandlerFunc(func(params operations.GetRobotsNamespaceParams) middleware.Responder {
		resp, err := svc.ListRobots(params.HTTPRequest.Context(), params.Namespace)
		if err != nil {
			return operations.NewGetRobotsNamespaceDefault(http.StatusInternalServerError)
		}
		return operations.NewGetRobotsNamespaceOK().WithPayload(resp)
	})

	api.GetRobotsNamespaceNameHandler = operations.GetRobotsNamespaceNameHandlerFunc(func(params operations.GetRobotsNamespaceNameParams) middleware.Responder {
		resp, err := svc.GetRobot(params.HTTPRequest.Context(), params.Namespace, params.Name)
		if err != nil {
			return operations.NewGetRobotsNamespaceDefault(http.StatusServiceUnavailable).WithPayload(err)
		}
		return operations.NewGetRobotsNamespaceNameOK().WithPayload(resp)
	})

	api.PostRegistryNamespaceNameHandler = operations.PostRegistryNamespaceNameHandlerFunc(func(params operations.PostRegistryNamespaceNameParams) middleware.Responder {
		resp, err := svc.Create(params.HTTPRequest.Context(), params.Repository)
		if err != nil {
			return operations.NewDeleteRegistryNamespaceNameDefault(http.StatusInternalServerError).WithPayload(err)
		}
		return operations.NewPostRegistryNamespaceNameOK().WithPayload(resp)
	})

	api.PostRobotsNamespaceNameHandler = operations.PostRobotsNamespaceNameHandlerFunc(func(params operations.PostRobotsNamespaceNameParams) middleware.Responder {
		resp, err := svc.CreateRobot(params.HTTPRequest.Context(), params.Robot)
		if err != nil {
			return operations.NewPostRobotsNamespaceNameDefault(http.StatusInternalServerError).WithPayload(err)
		}
		return operations.NewPostRobotsNamespaceNameOK().WithPayload(resp)
	})

	handler := alice.New(keyauth.New(ctx.String("auth-token")), authinfo.New).Then(api.Serve(nil))

	s.SetHandler(handler)

	// @step: start up the service
	go func() {
		if err := s.Serve(); err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Fatal("failed to start the api service")
		}
	}()

	// @step: wait for a signal to terminat
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-signalChannel

	return nil
}
