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

package keyauth

import (
	"fmt"
	"net/http"
)

// New creates and returns middleware
func New(token string) func(next http.Handler) http.Handler {
	expected := fmt.Sprintf("Bearer %s", token)

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if token != "" {
				if h, found := r.Header["Authorization"]; found {
					if h[0] != expected {
						w.WriteHeader(http.StatusForbidden)
						return
					}
				} else {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
			}
			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
