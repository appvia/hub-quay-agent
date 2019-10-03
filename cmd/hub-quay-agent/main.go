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
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	release = "v0.0.7"
)

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	app := &cli.App{
		Name:    "hub-quay-agent",
		Author:  "Rohith Jayawardene",
		Email:   "gambol99@gmail.com",
		Usage:   "A backend agent used to provision resources within quay.io",
		Version: release,

		OnUsageError: func(context *cli.Context, err error, _ bool) error {
			fmt.Fprintf(os.Stderr, "[error] invalid options %s\n", err)
			return err
		},

		Action: func(ctx *cli.Context) error {
			return invokeServerAction(ctx)
		},

		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "listen",
				Usage:  "the interface to bind the service to `INTERFACE`",
				Value:  "127.0.0.1",
				EnvVar: "LISTEN",
			},
			cli.IntFlag{
				Name:   "http-port",
				Usage:  "network interface the service should listen on `PORT`",
				Value:  10080,
				EnvVar: "HTTP_PORT",
			},
			cli.IntFlag{
				Name:   "https-port",
				Usage:  "network interface the service should listen on `PORT`",
				Value:  10443,
				EnvVar: "HTTPS_PORT",
			},
			cli.StringFlag{
				Name:   "tls-cert",
				Usage:  "the path to the file containing the certificate pem `PATH`",
				EnvVar: "TLS_CERT",
			},
			cli.StringFlag{
				Name:   "tls-key",
				Usage:  "the path to the file containing the private key pem `PATH`",
				EnvVar: "TLS_KEY",
			},
			cli.StringFlag{
				Name:   "auth-token",
				Usage:  "authentication token used to verifier the caller `TOKEN`",
				EnvVar: "AUTH_TOKEN",
			},
			cli.StringFlag{
				Name:   "quay-endpoint-url",
				Usage:  "the url for the quay.io api `URL`",
				Value:  "https://quay.io",
				EnvVar: "QUAY_ENDPOINT_URL",
			},
			cli.StringFlag{
				Name:   "quay-api-token",
				Usage:  "an authentication token used to permit api access `TOKEN`",
				EnvVar: "QUAY_API_TOKEN",
			},
			cli.BoolFlag{
				Name:   "verbose",
				Usage:  "indicates if we should enable verbose logging `BOOL`",
				EnvVar: "VERBOSE",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "[error] %s\n", err)
		os.Exit(1)
	}

}
