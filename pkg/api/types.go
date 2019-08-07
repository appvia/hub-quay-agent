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

package api

import (
	"context"

	"github.com/appvia/hub-quay-agent/pkg/transport/models"
)

// Handler defines the handlers for the service api
type Handler interface {
	// Create is responsible for creating a repo
	Create(context.Context, *models.Repository) (*models.Repository, *models.APIError)
	// CreateRobot is responsible for creating a robot account
	CreateRobot(context.Context, *models.Robot) (*models.Robot, *models.APIError)
	// CreateTeam is responsible for creation a team in the org
	CreateTeam(context.Context, *models.Team) (*models.Team, *models.APIError)
	// Delete is responsible for deleting a repo
	Delete(context.Context, string, string) *models.APIError
	// DeleteRobot is responsible for deleting robot account
	DeleteRobot(context.Context, string, string) *models.APIError
	// DeleteTeam is responsible for deleting a team in the org
	DeleteTeam(context.Context, string, string) *models.APIError
	// Get is responsible for retrieving a repository
	Get(context.Context, string, string) (*models.Repository, *models.APIError)
	// GetRobot is responsible for retrieving a robot
	GetRobot(context.Context, string, string) (*models.Robot, *models.APIError)
	// GetTeam is responsible for retrieving the state of a team
	GetTeam(context.Context, string, string) (*models.Team, *models.APIError)
	// Health is responsible for checking the health of the agent
	Health(context.Context, string) *models.APIError
	// List is responsible for listing all the repostories
	List(context.Context, string) (*models.RepositoryList, *models.APIError)
	// ListRobots is responsible for listing all the robots
	ListRobots(context.Context, string) (*models.RobotList, *models.APIError)
	// ListTeams is responsible for listing the teams
	ListTeams(context.Context, string) (*models.TeamList, *models.APIError)
}
