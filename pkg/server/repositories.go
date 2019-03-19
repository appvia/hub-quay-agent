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

// Create is responsible for creating a repo
func (s *serverImpl) Create(ctx context.Context, r *models.Repository) (*models.Repository, *models.APIError) {

	err := func() *models.APIError {
		// @note: the repository creation is idempotent
		if err := s.Repositories().Create(ctx, &client.NewRepo{
			Description: r.Spec.Description,
			Namespace:   sv(r.Namespace),
			Repository:  sv(r.Name),
			Visibility:  r.Spec.Visibility,
		}); err != nil {
			return newError("creating repository", err).model()
		}

		fullname := fmt.Sprintf("%s/%s", sv(r.Namespace), sv(r.Name))

		// @step: ensure the users are there
		members, err := s.Repositories().ListUsers(ctx, fullname)
		if err != nil {
			return newError("retrieving members", err).model()
		}
		adding, removing := createPermissions(members, r.Spec.Members)

		log.WithFields(log.Fields{
			"adding":   adding,
			"name":     fullname,
			"removing": removing,
		}).Debug("adjusted user permissions on repository")

		if err := s.Repositories().AddUsers(ctx, fullname, adding); err != nil {
			return newError("adding members", err).model()
		}
		if err := s.Repositories().DeleteUsers(ctx, fullname, removing); err != nil {
			return newError("removing members", err).model()
		}

		// @step: synchronize the robot accounts
		robots, err := s.Repositories().ListRobots(ctx, fullname)
		if err != nil {
			return newError("retrieving robots", err).model()
		}
		adding, removing = createPermissions(robots, r.Spec.Robots)

		log.WithFields(log.Fields{
			"adding":   adding,
			"name":     fullname,
			"removing": removing,
		}).Debug("adjusted robot permissions on repository")

		if err := s.Repositories().AddRobots(ctx, fullname, adding); err != nil {
			return newError("adding robots", err).model()
		}
		if err := s.Repositories().DeleteRobots(ctx, fullname, removing); err != nil {
			return newError("removing robots", err).model()
		}

		return nil
	}()
	if err != nil {
		log.WithFields(log.Fields{
			"error":     err.Reason,
			"name":      sv(r.Name),
			"namespace": sv(r.Namespace),
		}).Error("creating or updating repository")

		return nil, err
	}

	// @step: fill in the model for them
	if r.Spec.URL == "" {
		r.Spec.URL = sv(r.Name)
	}

	return r, nil
}

// Get is responsible for getting a repository
func (s *serverImpl) Get(ctx context.Context, namespace, name string) (*models.Repository, *models.APIError) {
	fullname := fmt.Sprintf("%s/%s", namespace, name)

	// @step: check if the resource exists
	if found, err := s.Repositories().Has(ctx, fullname); err != nil {
		return nil, newError("checking repository exists", err).model()
	} else if !found {
		return nil, newError("resource does not exist", nil).model()
	}

	repo := &models.Repository{
		Object: models.Object{
			Name:      sp(name),
			Namespace: sp(namespace),
		},
	}

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

		repo.Spec = &models.RepositorySpec{
			Description: r.Description,
			Tags:        make(map[string]string, 0),
			Robots:      make([]*models.Permission, 0),
			Members:     make([]*models.Permission, 0),
			Visibility:  visibility,
		}

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

// Delete is responsible for deleting a repo
func (s *serverImpl) Delete(ctx context.Context, namespace, name string) *models.APIError {
	log.WithFields(log.Fields{
		"name":      name,
		"namespace": namespace,
	}).Debug("attempting to delete the repository")

	fullname := fmt.Sprintf("%s/%s", namespace, name)

	// @note: the delete the idempotent, so it's fine to call without checking
	if err := s.Repositories().Delete(ctx, fullname); err != nil {
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
			Namespace: sp(namespace),
		},
		Items: make([]*models.Repository, 0),
	}
	for _, x := range repos.Repositories {
		repo, err := s.Get(ctx, x.Namespace, x.Name)
		if err != nil {
			return nil, err
		}
		list.Items = append(list.Items, repo)
	}

	return list, nil
}
