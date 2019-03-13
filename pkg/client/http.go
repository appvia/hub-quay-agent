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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

// authinfo provides the authentication for requests
func (c *clientImpl) authinfo(ctx context.Context, request *http.Request) {
	token := c.token

	// @step: check for auth in the context
	value := ctx.Value(AuthKey)
	if value != nil {
		if v, ok := value.(string); ok {
			token = v
		}
	}

	// @step: inject the authentication if required
	if token != "" {
		request.Header.Set("Authorization", "Bearer "+token)
	}
}

// handle is a generic handler for http requests to the api
func (c *clientImpl) Handle(ctx context.Context, method, uri string, payload, data interface{}) error {
	location := fmt.Sprintf("%s/api/v1/%s", c.endpoint, strings.TrimPrefix(uri, "/"))

	fields := log.Fields{
		"endpoint": c.endpoint,
		"method":   strings.ToLower(method),
		"url":      location,
	}
	defer func() {
		//log.WithFields(fields).Debug("making request to quay.io api")
	}()

	err := func() error {
		// @step: check is we have a data and if so encode it
		in, err := encode(payload)
		if err != nil {
			return err
		}

		// @step: create the http request
		request, err := http.NewRequest(method, location, in)
		if err != nil {
			return err
		}
		request.WithContext(ctx)
		request.Header.Set("Agent", "hub-quay-agent")
		request.Header.Set("Content-Type", "application/json")

		c.authinfo(ctx, request)

		fields["payload"] = spew.Sdump(payload)

		b, err := httputil.DumpRequest(request, true)
		if err == nil {
			fmt.Printf("REQUEST: %s\n", string(b))
		}

		// @step: perform the request
		resp, err := c.hc.Do(request)
		if err != nil {
			return err
		}
		fields["code"] = resp.StatusCode

		b, err = httputil.DumpResponse(resp, true)
		if err == nil {
			fmt.Printf("RESPONSE: %s\n", string(b))
			fields["response"] = string(b)
		}

		// @step: decode the response if required and or apierror
		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			if data == nil {
				return nil
			}
			return decode(resp.Body, data)
		}

		if resp.Body == nil {
			return &Error{ErrorMessage: "invalid api response", Status: resp.StatusCode}
		}

		apierror := &Error{}

		if err := decode(resp.Body, apierror); err != nil {
			return &Error{ErrorMessage: err.Error(), Status: resp.StatusCode}
		}
		apierror.Status = resp.StatusCode

		return apierror
	}()
	if err != nil {
		aerr, ok := err.(*Error)
		if !ok {
			return &Error{Status: http.StatusInternalServerError, ErrorMessage: err.Error()}
		}

		return aerr
	}

	return nil
}

func decode(in io.Reader, out interface{}) error {
	return json.NewDecoder(in).Decode(out)
}

func encode(in interface{}) (io.Reader, error) {
	if in == nil {
		return nil, nil
	}
	buf := &bytes.Buffer{}

	return buf, json.NewEncoder(buf).Encode(in)
}
