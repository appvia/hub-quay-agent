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

	"github.com/appvia/hub-quay-agent/pkg/client"
	"github.com/appvia/hub-quay-agent/pkg/transport/models"

	log "github.com/sirupsen/logrus"
)

// CreateRobot creates and returns a robot account for us - note it does not do permissions, that is done
// at the repository level
func (s *serverImpl) CreateRobot(ctx context.Context, robot *models.Robot) (*models.Robot, *models.APIError) {
	namespace, name := ParseName(sv(robot.Name))

	log.WithFields(log.Fields{
		"name":      name,
		"namespace": namespace,
	}).Debug("attempting to create or update robot token")

	r, err := s.Robots().Create(ctx, &client.Robot{
		Description: sv(robot.Spec.Description),
		Name:        sv(robot.Name),
	})
	if err != nil {
		log.WithFields(log.Fields{
			"error":     err.Error(),
			"name":      name,
			"namespace": namespace,
		}).Error("creating or updating robot account")

		return nil, newError("creating or update robot account", err).model()
	}
	robot.Spec.Token = r.Token

	return robot, nil
}

// DeleteRobot is responsible for deleting the robot accounts
func (s *serverImpl) DeleteRobot(ctx context.Context, fullname string) *models.APIError {
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
func (s *serverImpl) GetRobot(ctx context.Context, fullname string) (*models.Robot, *models.APIError) {
	if found, err := s.Robots().Has(ctx, fullname); err != nil {
		return nil, newError("checking robot exists", err).model()
	} else if !found {
		return nil, newError("resource does not exist", nil).model()
	}
	robot := &models.Robot{Object: models.Object{Name: sp(fullname)}}

	r, err := s.Robots().Get(ctx, fullname)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
			"name":  fullname,
		}).Error("retrieving robot account")

		return nil, newError("retrieving robot account", err).model()
	}
	robot.Spec = &models.RobotSpec{
		Description: sp(r.Description),
		Token:       r.Token,
	}

	return robot, nil
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
		Object: models.Object{Name: sp(namespace)},
	}

	for _, x := range robots.Robots {
		list.Items = append(list.Items, &models.Robot{
			Object: models.Object{
				Name: sp(x.Name),
			},
			Spec: &models.RobotSpec{
				Description: sp(x.Description),
				Token:       x.Token,
			},
		})
	}

	return list, nil
}
