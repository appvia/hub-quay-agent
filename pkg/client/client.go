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

package client

import (
	"net/http"
	"net/url"
	"strings"
	"time"
)

type clientImpl struct {
	// hc is the http client to operate with
	hc *http.Client
	// endpoint is the quay endpoint
	endpoint string
	// token is the access token for a api
	token string
	// rc is the repos client
	rc Repositories
	// roc is the robots client
	roc Robots
}

// New creates a new client
func New(endpoint, token string) (Client, error) {
	if _, err := url.Parse(endpoint); err != nil {
		return nil, err
	}
	if !strings.HasPrefix(endpoint, "https://") && !strings.HasPrefix(endpoint, "http://") {
		endpoint = "https://" + endpoint
	}
	hc := &http.Client{Timeout: 10 * time.Second}

	qc := &clientImpl{endpoint: endpoint, hc: hc, token: token}
	qc.rc = &repositoryImpl{Client: qc}
	qc.roc = &robotImpl{Client: qc}

	return qc, nil
}

// Repositories returns the repostories client
func (c *clientImpl) Repositories() Repositories {
	return c.rc
}

// Robots returns the repostories client
func (c *clientImpl) Robots() Robots {
	return c.roc
}
