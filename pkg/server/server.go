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

	"github.com/appvia/hub-quay-agent/pkg/api"
	"github.com/appvia/hub-quay-agent/pkg/client"
	"github.com/appvia/hub-quay-agent/pkg/transport/models"

	log "github.com/sirupsen/logrus"
)

// serverImpl is the service state wrapper
type serverImpl struct {
	client.Client
}

// New creates and returns a new api handler
func New(o *Options) (api.Handler, error) {
	log.WithFields(log.Fields{
		"hostname": o.HostnameAPI,
	}).Info("creating quay hub agent provider")

	// @step: create the quay client
	qc, err := client.New(o.HostnameAPI, o.AccessToken)
	if err != nil {
		return nil, err
	}

	return &serverImpl{Client: qc}, nil
}

// Create is responsible for creating a repo
func (s *serverImpl) Create(ctx context.Context, r *models.Repository) (*models.Repository, *models.APIError) {
	namespace, name := ParseName(sv(r.Name))

	// @note: the repository creation is idempotent
	err := s.Repositories().Create(ctx, &client.NewRepo{
		Description: r.Spec.Description,
		Namespace:   namespace,
		Repository:  name,
		Visibility:  r.Spec.Visibility,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"error":     err.Error(),
			"name":      name,
			"namespace": namespace,
		}).Error("failed to create the repository")

		return nil, newError("failed to create the repository", err).model()
	}

	return r, nil
}

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
		Name:        name,
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

// Get is responsible for getting a repository
func (s *serverImpl) Get(ctx context.Context, fullname string) (*models.Repository, *models.APIError) {
	namespace, name := ParseName(fullname)

	// @step: check if the resource exists
	if found, err := s.Repositories().Has(ctx, fullname); err != nil {
		return nil, newError("checking repository exists", err).model()
	} else if !found {
		return nil, newError("resource does not exist", nil).model()
	}

	repo := &models.Repository{Object: models.Object{Name: sp(fullname)}}

	reason, err := func() (string, error) {
		r, err := s.Repositories().Get(ctx, fullname)
		if err != nil {
			return "retrieving the repository", err
		}

		// @step: get a list of members and robots
		members, err := s.Repositories().ListUsers(ctx, fullname)
		if err != nil {
			return "retrieving repository members", err
		}
		robots, err := s.Repositories().ListRobots(ctx, fullname)
		if err != nil {
			return "retrieving repository robots", err
		}

		visibility := "public"
		if !r.IsPublic {
			visibility = "private"
		}

		repo.Spec = &models.RepositorySpec{Visibility: visibility}
		for _, x := range members {
			repo.Spec.Members = append(repo.Spec.Members, &models.Permission{Name: &x.Name, Permission: &x.Role})
		}
		for _, x := range robots {
			repo.Spec.Robots = append(repo.Spec.Robots, &models.Permission{Name: &x.Name, Permission: &x.Role})
		}

		return "", nil
	}()
	if err != nil {
		log.WithFields(log.Fields{
			"error":     err.Error(),
			"name":      name,
			"namespace": namespace,
			"reason":    reason,
		}).Error("failed to retrieve repository")

		return nil, newError(reason, err).model()
	}

	return repo, nil
}

// GetRobot retrieves the robot account if any
func (s *serverImpl) GetRobot(ctx context.Context, fullname string) (*models.Robot, *models.APIError) {
	namespace, name := ParseName(fullname)

	if found, err := s.Robots().Has(ctx, name); err != nil {
		return nil, newError("checking robot exists", err).model()
	} else if !found {
		return nil, newError("resource does not exist", nil).model()
	}
	robot := &models.Robot{Object: models.Object{Name: sp(fullname)}}

	r, err := s.Robots().Get(ctx, name)
	if err != nil {
		log.WithFields(log.Fields{
			"error":     err.Error(),
			"name":      name,
			"namespace": namespace,
		}).Error("retrieving robot account")

		return nil, newError("retrieving robot account", err).model()
	}
	robot.Spec = &models.RobotSpec{
		Description: sp(r.Description),
		Token:       r.Token,
	}

	return robot, nil
}

// Delete is responsible for deleting a repo
func (s *serverImpl) Delete(ctx context.Context, repo *models.Repository) *models.APIError {
	namespace, name := ParseName(sv(repo.Name))

	log.WithFields(log.Fields{
		"name":      name,
		"namespace": namespace,
	}).Debug("attempting to delete the repository")

	// @note: the delete the idempotent, so it's fine to call without checking
	if err := s.Repositories().Delete(ctx, name); err != nil {
		log.WithFields(log.Fields{
			"error":     err.Error(),
			"name":      name,
			"namespace": namespace,
		}).Error("deleting the repository")

		return newError("deleting the repository", err).model()
	}

	return nil
}

// List is responsible for listing all the repostories
func (s *serverImpl) List(ctx context.Context, namespace string) (*models.RepositoryList, *models.APIError) {
	repos, err := s.Repositories().List(ctx, namespace)
	if err != nil {
		log.WithFields(log.Fields{
			"error":     err.Error(),
			"namespace": namespace,
		}).Error("listing repositories")

		return nil, newError("listing repositories", err).model()
	}

	list := &models.RepositoryList{
		Object: models.Object{
			Name: sp(namespace),
		},
	}
	for _, x := range repos.Repositories {
		repo, err := s.Get(ctx, fmt.Sprintf("%s/%s", x.Namespace, x.Name))
		if err != nil {
			return nil, err
		}
		list.Items = append(list.Items, repo)
	}

	return list, nil
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
