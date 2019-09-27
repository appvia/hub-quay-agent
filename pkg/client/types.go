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
	"errors"
)

// AuthInfo provides the authentication info
type AuthInfo string

var (
	// AuthKey is the contextual key for auth
	AuthKey AuthInfo = "authinfo"
)

var (
	// ErrUnauthorized indicates the credentials are incorrect
	ErrUnauthorized = errors.New("invalid credentials or permissions")
)

// Client is the client contract
type Client interface {
	// Repositories returns the repositories api
	Repositories() Repositories
	// Robots returns the robots api
	Robots() Robots
	// Teams returns the teams client
	Teams() Teams
	// Handle is a generic handler for the http requests
	Handle(context.Context, string, string, interface{}, interface{}) error
}

// Repositories is the contract to the repositories
type Repositories interface {
	// ImageAnalysis returns the scan analysis for a specific tag
	ImageAnalysis(context.Context, string, string) (*ImageAnalysis, error)
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
	// ListUsers is responsible for getting the user permissions
	ListUsers(context.Context, string) ([]*Permission, error)
	// ListRobots is responsible for getting the robot permissions
	ListRobots(context.Context, string) ([]*Permission, error)
	// ListTeams is responsibl for getting the team permissions
	ListTeams(context.Context, string) ([]*Permission, error)
	// AddUsers is responsible for adding a user
	AddUsers(context.Context, string, []*Permission) error
	// DeleteUser removes a user permission
	DeleteUsers(context.Context, string, []*Permission) error
	// AddRobots is responsible for adding a robot
	AddRobots(context.Context, string, []*Permission) error
	// DeleteRobots removes a user permission
	DeleteRobots(context.Context, string, []*Permission) error
	// AddTeams is responsible for adding a robot
	AddTeams(context.Context, string, []*Permission) error
	// DeleteTeams removes a user permission
	DeleteTeams(context.Context, string, []*Permission) error
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

// Teams is the contract to the teams API
type Teams interface {
	// Create is responsible creating a team
	Create(context.Context, *Team, *Members) (*Team, error)
	// Delete is responsible deleting the team
	Delete(context.Context, string) error
	// Get is responsible getting a team
	Get(context.Context, string) (*Team, error)
	// Has is responsible checking if the team exists
	Has(context.Context, string) (bool, error)
	// List is responsible getting a list o teams
	List(context.Context, string) (*TeamList, error)
	// ListMembers is responsible for listing the members in a team
	ListMembers(context.Context, string) (*Members, error)
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
