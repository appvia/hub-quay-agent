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

package authinfo

import (
	"context"
	"net/http"

	"github.com/appvia/hub-quay-agent/pkg/client"
)

const (
	// AuthInfoHeader is the passed authentication header
	AuthInfoHeader = "X-AuthInfo"
)

type authinfoImpl struct{}

// New creates and returns an middlware for auth
func New(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if v, found := r.Header[AuthInfoHeader]; found {
			if len(v) >= 1 {
				ctx := context.WithValue(r.Context(), client.AuthKey, v[0])

				next.ServeHTTP(w, r.WithContext(ctx))
			}
		} else {
			next.ServeHTTP(w, r)
		}
	}

	return http.HandlerFunc(fn)
}
