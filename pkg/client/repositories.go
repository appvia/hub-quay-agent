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

type repositoryImpl struct {
	Client
}

// Create is responsible for creating a repository
func (r *repositoryImpl) Create(ctx context.Context, repo *NewRepo) error {
	repo.RepoKind = "image"

	if found, err := r.Has(ctx, repo.Name()); err != nil {
		return err
	} else if found {
		// @check is anything has been changes
		current, err := r.Get(ctx, repo.Name())
		if err != nil {
			return err
		}

		if current.Description != repo.Description {
			uri := fmt.Sprintf("/repository/%s/%s", repo.Namespace, repo.Name)

			if err := r.Handle(ctx, http.MethodPut, uri, repo, nil); err != nil {
				return fmt.Errorf("unable to update desription: %s", err)
			}
		}
		if current.IsPublic && repo.Visibility == "private" || !current.IsPublic && repo.Visibility == "public" {
			uri := fmt.Sprintf("/repository/%s/%s/changevisibility", repo.Namespace, repo.Name)

			if err := r.Handle(ctx, http.MethodPost, uri, repo, nil); err != nil {
				return fmt.Errorf("unable to update visibility: %s", err)
			}
		}
	}

	return r.Handle(ctx, http.MethodPost, "/repository", &repo, nil)
}

// Delete is responsible for deleting a repository
func (r *repositoryImpl) Delete(ctx context.Context, name string) error {
	// @step: check if the repository exists
	if found, err := r.Has(ctx, name); err != nil {
		return err
	} else if !found {
		return nil
	}
	uri := fmt.Sprintf("/repository/%s", name)

	return r.Handle(ctx, http.MethodDelete, uri, nil, nil)
}

// Has checks if a repository exists
func (r *repositoryImpl) Has(ctx context.Context, name string) (bool, error) {
	if _, err := r.Get(ctx, name); err != nil {
		aerr, ok := err.(*Error)
		if ok {
			if aerr.Status == http.StatusNotFound {
				return false, nil
			}
		}

		return false, err
	}

	return true, nil
}

// Get retrieves a repository from the registry
func (r *repositoryImpl) Get(ctx context.Context, name string) (*Repository, error) {
	repo := &Repository{}
	uri := fmt.Sprintf("/repository/%s?includeTags=true", name)

	return repo, r.Handle(ctx, http.MethodGet, uri, nil, repo)
}

// AddUser is responsible for adding a user
func (r *repositoryImpl) AddUser(ctx context.Context, repo, member, role string) error {
	var userRole struct {
		Role string `json:"role"`
	}
	uri := fmt.Sprintf("/repository/%s/permissions/user/%s", repo, member)

	return r.Handle(ctx, http.MethodPut, uri, &userRole, nil)
}

// AddRobot is responsible for adding a robot
func (r *repositoryImpl) AddRobot(ctx context.Context, repoName, robotName, roleName string) error {
	// @step: ensure the robot user exists
	if found, err := r.Robots().Has(ctx, robotName); err != nil {
		return err
	} else if !found {
		return errors.New("robot user not found")
	}

	return r.AddUser(ctx, repoName, robotName, roleName)
}

func (r *repositoryImpl) List(ctx context.Context, namespace string) (*RepositoryList, error) {
	list := &RepositoryList{}
	token := ""

	for max := 0; max < 200; max++ {
		resp := &RepositoryList{}

		uri := fmt.Sprintf("/repository?namespace=%s", namespace)
		if token != "" {
			uri = fmt.Sprintf("%s&next_page=%s", uri, token)
		}

		if err := r.Handle(ctx, http.MethodGet, uri, nil, &resp); err != nil {
			return nil, err
		}
		for _, x := range resp.Repositories {
			list.Repositories = append(list.Repositories, x)
		}

		if resp.NextPage == "" {
			return list, nil
		}
		token = resp.NextPage
	}

	return nil, fmt.Errorf("reached the max number of pages when listing repositories")
}

// ListPermisions is a list of users permissions for robots and users
func (r *repositoryImpl) ListPermissions(ctx context.Context, name string) ([]*Permission, error) {
	return r.permissions(ctx, name, true, true)
}

// ListUsers returns the users associated to the repository
func (r *repositoryImpl) ListUsers(ctx context.Context, name string) ([]*Permission, error) {
	return r.permissions(ctx, name, false, true)
}

// ListRobots returns the robots associated to the repository
func (r *repositoryImpl) ListRobots(ctx context.Context, name string) ([]*Permission, error) {
	return r.permissions(ctx, name, true, false)
}

func (r *repositoryImpl) permissions(ctx context.Context, name string, robots, users bool) ([]*Permission, error) {
	var perms struct {
		// Permissions is a collection of permission for user and robots
		Permissions map[string]*Permission `json:"permissions,omitempty"`
	}
	uri := fmt.Sprintf("/repository/%s/permissions/user", name)

	if err := r.Handle(ctx, http.MethodGet, uri, nil, &perms); err != nil {
		return nil, err
	}

	var list []*Permission
	for _, x := range perms.Permissions {
		if x.IsRobot && robots {
			list = append(list, x)
		}
		if !x.IsRobot && users {
			list = append(list, x)
		}
	}

	return list, nil
}
