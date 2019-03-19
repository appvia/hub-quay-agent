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
	"fmt"

	"github.com/appvia/hub-quay-agent/pkg/client"
	"github.com/appvia/hub-quay-agent/pkg/transport/models"

	log "github.com/sirupsen/logrus"
)

// CreateRobot creates and returns a robot account for us - note it does not do permissions, that is done
// at the repository level
func (s *serverImpl) CreateRobot(ctx context.Context, robot *models.Robot) (*models.Robot, *models.APIError) {
	log.WithFields(log.Fields{
		"name":      robot.Name,
		"namespace": robot.Namespace,
	}).Debug("attempting to create or update robot token")

	fullname := fmt.Sprintf("%s+%s", sv(robot.Namespace), sv(robot.Name))

	r, err := s.Robots().Create(ctx, &client.Robot{
		Description: sv(robot.Spec.Description),
		Name:        fullname,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"error":     err.Error(),
			"name":      robot.Name,
			"namespace": robot.Namespace,
		}).Error("creating or updating robot account")

		return nil, newError("creating or update robot account", err).model()
	}
	robot.Spec.Token = r.Token

	return robot, nil
}

// DeleteRobot is responsible for deleting the robot accounts
func (s *serverImpl) DeleteRobot(ctx context.Context, namespace, name string) *models.APIError {
	fullname := fmt.Sprintf("%s+%s", namespace, name)

	if found, err := s.Robots().Has(ctx, fullname); err != nil {
		return newError("checking robot exists", err).model()
	} else if !found {
		return newError("resource does not exist", nil).model()
	}

	if err := s.Robots().Delete(ctx, fullname); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
			"name":  fullname,
		}).Error("deleting robot account")

		return newError("deleting robot account", err).model()
	}

	return nil
}

// GetRobot retrieves the robot account if any
func (s *serverImpl) GetRobot(ctx context.Context, namespace, name string) (*models.Robot, *models.APIError) {
	fullname := fmt.Sprintf("%s+%s", namespace, name)

	if found, err := s.Robots().Has(ctx, fullname); err != nil {
		return nil, newError("checking robot exists", err).model()
	} else if !found {
		return nil, newError("resource does not exist", nil).model()
	}

	r, err := s.Robots().Get(ctx, fullname)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
			"name":  fullname,
		}).Error("retrieving robot account")

		return nil, newError("retrieving robot account", err).model()
	}

	return &models.Robot{
		Object: models.Object{
			Name:      sp(name),
			Namespace: sp(namespace),
		},
		Spec: &models.RobotSpec{
			Description: sp(r.Description),
			Token:       r.Token,
		},
	}, nil
}

// ListRobots is responsible for listing all the robots
func (s *serverImpl) ListRobots(ctx context.Context, namespace string) (*models.RobotList, *models.APIError) {
	robots, err := s.Robots().List(ctx, namespace)
	if err != nil {
		log.WithFields(log.Fields{
			"error":     err.Error(),
			"namespace": namespace,
		}).Error("listing robot accounts")

		return nil, newError("listing robot accounts", err).model()
	}

	list := &models.RobotList{
		Object: models.Object{Namespace: sp(namespace)},
		Items:  make([]*models.Robot, 0),
	}

	for _, x := range robots.Robots {
		ns, n := parseName(x.Name)

		list.Items = append(list.Items, &models.Robot{
			Object: models.Object{
				Name:      sp(n),
				Namespace: sp(ns),
			},
			Spec: &models.RobotSpec{
				Description: sp(x.Description),
				Token:       x.Token,
			},
		})
	}

	return list, nil
}
