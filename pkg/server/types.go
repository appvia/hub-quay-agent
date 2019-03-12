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
	"errors"
)

var (
	// ErrResourceNotFound indicates the resource does not exist
	ErrResourceNotFound = errors.New("resource does not exist")
	// ErrRetrivingResource indicates an error retirving the resource
	ErrRetrivingResource = errors.New("failed to retrieve resource")
)

// Options is the configuration for the service
type Options struct {
	// AccessToken to the token use to call the api
	AccessToken string
	// HostnameAPI is the hostname of the api
	HostnameAPI string
}
