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
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/appvia/hub-quay-agent/pkg/server"
	"github.com/appvia/hub-quay-agent/pkg/server/middleware/authinfo"
	"github.com/appvia/hub-quay-agent/pkg/transport/models"
	"github.com/appvia/hub-quay-agent/pkg/transport/restapi"
	"github.com/appvia/hub-quay-agent/pkg/transport/restapi/operations"

	errors "github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	flags "github.com/jessevdk/go-flags"
	"github.com/justinas/alice"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// invokeServerAction handles the command line server action
func invokeServerAction(ctx *cli.Context) error {
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

	apiToken := ""
	if ctx.String("auth-token") != "" {
		apiToken = fmt.Sprintf("Bearer %s", ctx.String("auth-token"))
	}

	api.ApikeyAuth = func(token string) (*models.Principal, error) {
		// @step: if no authentication we can pass straight through
		if apiToken == "" {
			return nil, nil
		}
		if apiToken == "" {
			return nil, errors.New(http.StatusUnauthorized, "authentication required")
		}
		if apiToken != token {
			return nil, errors.New(http.StatusForbidden, "invalid authentication")
		}

		return nil, nil
	}

	api.GetAliveHandler = operations.GetAliveHandlerFunc(func(params operations.GetAliveParams) middleware.Responder {
		return operations.NewGetAliveOK()
	})

	api.GetHealthzNamespaceHandler = operations.GetHealthzNamespaceHandlerFunc(func(params operations.GetHealthzNamespaceParams, principal *models.Principal) middleware.Responder {
		if err := svc.Health(params.HTTPRequest.Context(), params.Namespace); err != nil {
			return operations.NewGetHealthzNamespaceDefault(http.StatusInternalServerError).WithPayload(err)
		}
		return operations.NewGetHealthzNamespaceOK()
	})

	api.DeleteRegistryNamespaceNameHandler = operations.DeleteRegistryNamespaceNameHandlerFunc(func(params operations.DeleteRegistryNamespaceNameParams, principal *models.Principal) middleware.Responder {
		if err := svc.Delete(params.HTTPRequest.Context(), params.Namespace, params.Name); err != nil {
			return operations.NewDeleteRegistryNamespaceNameDefault(http.StatusServiceUnavailable).WithPayload(err)
		}
		return operations.NewDeleteRegistryNamespaceNameOK()
	})

	api.DeleteRobotsNamespaceNameHandler = operations.DeleteRobotsNamespaceNameHandlerFunc(func(params operations.DeleteRobotsNamespaceNameParams, principal *models.Principal) middleware.Responder {
		if err := svc.DeleteRobot(params.HTTPRequest.Context(), params.Namespace, params.Name); err != nil {
			return operations.NewDeleteRobotsNamespaceNameDefault(http.StatusServiceUnavailable).WithPayload(err)
		}
		return operations.NewDeleteRobotsNamespaceNameOK()
	})

	api.DeleteTeamsNamespaceNameHandler = operations.DeleteTeamsNamespaceNameHandlerFunc(func(params operations.DeleteTeamsNamespaceNameParams, principal *models.Principal) middleware.Responder {
		if err := svc.DeleteTeam(params.HTTPRequest.Context(), params.Namespace, params.Name); err != nil {
			return operations.NewDeleteTeamsNamespaceNameDefault(http.StatusServiceUnavailable).WithPayload(err)
		}
		return operations.NewDeleteTeamsNamespaceNameOK()
	})

	api.GetRegistryNamespaceHandler = operations.GetRegistryNamespaceHandlerFunc(func(params operations.GetRegistryNamespaceParams, principal *models.Principal) middleware.Responder {
		resp, err := svc.List(params.HTTPRequest.Context(), params.Namespace)
		if err != nil {
			return operations.NewGetRegistryNamespaceNameDefault(http.StatusServiceUnavailable).WithPayload(err)
		}
		return operations.NewGetRegistryNamespaceOK().WithPayload(resp)
	})

	api.GetRegistryNamespaceNameHandler = operations.GetRegistryNamespaceNameHandlerFunc(func(params operations.GetRegistryNamespaceNameParams, principal *models.Principal) middleware.Responder {
		resp, err := svc.Get(params.HTTPRequest.Context(), params.Namespace, params.Name)
		if err != nil {
			return operations.NewGetRegistryNamespaceNameDefault(http.StatusServiceUnavailable).WithPayload(err)
		}
		return operations.NewGetRegistryNamespaceNameOK().WithPayload(resp)
	})

	api.GetRobotsNamespaceHandler = operations.GetRobotsNamespaceHandlerFunc(func(params operations.GetRobotsNamespaceParams, principal *models.Principal) middleware.Responder {
		resp, err := svc.ListRobots(params.HTTPRequest.Context(), params.Namespace)
		if err != nil {
			return operations.NewGetRobotsNamespaceDefault(http.StatusInternalServerError)
		}
		return operations.NewGetRobotsNamespaceOK().WithPayload(resp)
	})

	api.GetRobotsNamespaceNameHandler = operations.GetRobotsNamespaceNameHandlerFunc(func(params operations.GetRobotsNamespaceNameParams, principal *models.Principal) middleware.Responder {
		resp, err := svc.GetRobot(params.HTTPRequest.Context(), params.Namespace, params.Name)
		if err != nil {
			return operations.NewGetRobotsNamespaceDefault(http.StatusServiceUnavailable).WithPayload(err)
		}
		return operations.NewGetRobotsNamespaceNameOK().WithPayload(resp)
	})

	api.GetTeamsNamespaceHandler = operations.GetTeamsNamespaceHandlerFunc(func(params operations.GetTeamsNamespaceParams, principal *models.Principal) middleware.Responder {
		resp, err := svc.ListTeams(params.HTTPRequest.Context(), params.Namespace)
		if err != nil {
			return operations.NewGetTeamsNamespaceDefault(http.StatusInternalServerError)
		}
		return operations.NewGetTeamsNamespaceOK().WithPayload(resp)
	})

	api.GetTeamsNamespaceNameHandler = operations.GetTeamsNamespaceNameHandlerFunc(func(params operations.GetTeamsNamespaceNameParams, principal *models.Principal) middleware.Responder {
		resp, err := svc.GetTeam(params.HTTPRequest.Context(), params.Namespace, params.Name)
		if err != nil {
			return operations.NewGetTeamsNamespaceDefault(http.StatusServiceUnavailable).WithPayload(err)
		}
		return operations.NewGetTeamsNamespaceNameOK().WithPayload(resp)
	})

	api.PutRegistryNamespaceNameHandler = operations.PutRegistryNamespaceNameHandlerFunc(func(params operations.PutRegistryNamespaceNameParams, principal *models.Principal) middleware.Responder {
		resp, err := svc.Create(params.HTTPRequest.Context(), params.Repository)
		if err != nil {
			return operations.NewDeleteRegistryNamespaceNameDefault(http.StatusInternalServerError).WithPayload(err)
		}
		return operations.NewPutRegistryNamespaceNameOK().WithPayload(resp)
	})

	api.PutRobotsNamespaceNameHandler = operations.PutRobotsNamespaceNameHandlerFunc(func(params operations.PutRobotsNamespaceNameParams, principal *models.Principal) middleware.Responder {
		resp, err := svc.CreateRobot(params.HTTPRequest.Context(), params.Robot)
		if err != nil {
			return operations.NewPutRobotsNamespaceNameDefault(http.StatusInternalServerError).WithPayload(err)
		}
		return operations.NewPutRobotsNamespaceNameOK().WithPayload(resp)
	})

	api.PutTeamsNamespaceNameHandler = operations.PutTeamsNamespaceNameHandlerFunc(func(params operations.PutTeamsNamespaceNameParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation .PutTeamsNamespaceName has not yet been implemented")
	})

	handler := alice.New(authinfo.New).Then(api.Serve(nil))

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
