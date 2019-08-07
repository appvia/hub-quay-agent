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
	"context"

	"github.com/appvia/hub-quay-agent/pkg/transport/models"
)

// Health responds to the health of the agent
func (s *serverImpl) Health(ctx context.Context, namespace string) *models.APIError {
	// @step: attempt to list all the repositories in the org
	if _, err := s.Repositories().List(ctx, namespace); err != nil {
		return newError("listing the repositories", err).model()
	}

	return nil
}