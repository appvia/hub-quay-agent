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
	"fmt"
	"net/http"
)

type robotImpl struct {
	Client
}

// Create is responsible creating a robot
func (r *robotImpl) Create(ctx context.Context, robot *Robot) (*Robot, error) {
	namespace, name := parseRobotName(robot.Name)
	if namespace == "" {
		return nil, errors.New("invalid robot name")
	}
	uri := fmt.Sprintf("organization/%s/robots/%s", namespace, name)

	// @step: chec if the robot token exist already
	if found, err := r.Has(ctx, robot.Name); err != nil {
		return nil, err
	} else if found {
		current, err := r.Get(ctx, robot.Name)
		if err != nil {
			return nil, err
		}
		if current.Description != robot.Description {
			if err := r.Handle(ctx, http.MethodPut, uri, robot, &robot); err != nil {
				return robot, err
			}
		}

		return robot, nil
	}
	resp := &Robot{Description: robot.Description}

	return resp, r.Handle(ctx, http.MethodPut, uri, resp, resp)
}

// Delete is responsible deleting the robot
func (r *robotImpl) Delete(ctx context.Context, fullname string) error {
	if found, err := r.Has(ctx, fullname); err != nil {
		return err
	} else if !found {
		return nil
	}
	namespace, name := parseRobotName(fullname)

	uri := fmt.Sprintf("organization/%s/robots/%s", namespace, name)

	return r.Handle(ctx, http.MethodDelete, uri, nil, nil)
}

// Has is responsible checking if a robot token exists
func (r *robotImpl) Has(ctx context.Context, name string) (bool, error) {
	if _, err := r.Get(ctx, name); err != nil {
		if aerr, ok := err.(*Error); ok {
			if aerr.Status == http.StatusNotFound {
				return false, nil
			}
		}

		return false, err
	}

	return true, nil
}

// Get is responsible getting a robot
func (r *robotImpl) Get(ctx context.Context, fullname string) (*Robot, error) {
	namespace, name := parseRobotName(fullname)
	if namespace == "" {
		return nil, errors.New("invalid robot name")
	}
	robot := &Robot{}

	uri := fmt.Sprintf("organization/%s/robots/%s", namespace, name)

	return robot, r.Handle(ctx, http.MethodGet, uri, nil, robot)
}

// List is responsible getting a list o robots
func (r *robotImpl) List(ctx context.Context, namespace string) (*RobotList, error) {
	list := &RobotList{}

	uri := fmt.Sprintf("organization/%s/robots", namespace)

	return list, r.Handle(ctx, http.MethodGet, uri, nil, list)
}
