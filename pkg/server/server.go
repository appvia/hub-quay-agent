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

package server

import (
	"github.com/appvia/hub-quay-agent/pkg/api"
	"github.com/appvia/hub-quay-agent/pkg/client"

	log "github.com/sirupsen/logrus"
)

// serverImpl is the service state wrapper
type serverImpl struct {
	client.Client
}

// New creates and returns a new api handler
func New(o *Options) (api.Handler, error) {
	log.WithFields(log.Fields{
		"hostname": o.HostnameAPI,
	}).Info("creating quay hub agent provider")

	// @step: create the quay client
	qc, err := client.New(o.HostnameAPI, o.AccessToken)
	if err != nil {
		return nil, err
	}

	return &serverImpl{Client: qc}, nil
}
