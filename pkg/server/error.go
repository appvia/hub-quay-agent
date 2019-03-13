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
	"fmt"

	"github.com/appvia/hub-quay-agent/pkg/transport/models"
)

// apierror is a api error
type apierror struct {
	// detail provides additional detail to the error
	detail string
	// reason is the error message
	reason string
}

func (a *apierror) Error() string {
	return a.reason
}

func newError(reason string, err error) *apierror {
	a := &apierror{reason: fmt.Sprintf("failed: %s", reason)}
	if err != nil {
		a.detail = err.Error()
	}

	return a
}

func (a *apierror) model() *models.APIError {
	m := &models.APIError{Reason: sp(a.reason), Detail: a.detail}

	return m
}
