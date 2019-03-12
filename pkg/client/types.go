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
	"context"
)

// AuthInfo provides the authentication info
type AuthInfo string

var (
	// AuthKey is the contextual key for auth
	AuthKey AuthInfo = "authinfo"
)

// Client is the client contract
type Client interface {
	// Repositories returns the repositories api
	Repositories() Repositories
	// Robots returns the robots api
	Robots() Robots
	// Handle is a generic handler for the http requests
	Handle(context.Context, string, string, interface{}, interface{}) error
}

// Repositories is the contract to the repositories
type Repositories interface {
	// Create is responsible for creating a repository
	Create(context.Context, *NewRepo) error
	// Delete is responsible for deleting a repository
	Delete(context.Context, string) error
	// Get is responsible for getting a repository
	Get(context.Context, string) (*Repository, error)
	// Has checks if the repository exists
	Has(context.Context, string) (bool, error)
	// List is responsible for listing all the repositories
	List(context.Context, string) (*RepositoryList, error)
	// ListPermissons is responsible for listing all the permissions
	ListPermissions(context.Context, string) ([]*Permission, error)
	// ListUsers is responsible for getting the user permissions
	ListUsers(context.Context, string) ([]*Permission, error)
	// ListRobots is responsible for getting the robot permissions
	ListRobots(context.Context, string) ([]*Permission, error)
	// AddUser is responsible for adding a user
	AddUser(context.Context, string, string, string) error
	// AddRobot is responsible for adding a robot
	AddRobot(context.Context, string, string, string) error
}

// Robots is the contract to the robots
type Robots interface {
	// Create is responsible creating a robot
	Create(context.Context, *Robot) (*Robot, error)
	// Delete is responsible deleting the robot
	Delete(context.Context, string) error
	// Get is responsible getting a robot
	Get(context.Context, string) (*Robot, error)
	// Has is responsible checking if the robot exists
	Has(context.Context, string) (bool, error)
	// List is responsible getting a list o robots
	List(context.Context, string) (*RobotList, error)
}

// Error is a generic error handed back by the API
type Error struct {
	// Status is a status of an error
	Status int `json:"status"`
	// ErrorMessage is a description to the message
	ErrorMessage string `json:"error_message"`
	// Title is a title to the error
	Title string `json:"title"`
	// ErrorType defines a error type
	ErrorType string `json:"error_type"`
	// Details provides additional detail
	Detail string `json:"detail"`
	// Message is a human readible error message
	Message string `json:"message"`
	// Type is a error message type
	Type string `json:"type"`
}
