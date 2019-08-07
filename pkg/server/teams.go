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

// CreateTeam is responsible for creation a team in the org
func (s *serverImpl) CreateTeam(ctx context.Context, team *models.Team) (*models.Team, *models.APIError) {
	log.WithFields(log.Fields{
		"name":      team.Name,
		"namespace": team.Namespace,
	}).Debug("attempting to create or update team")

	fullname := fmt.Sprintf("%s/%s", sv(team.Namespace), sv(team.Name))

	model := &client.Team{
		Name: fullname,
		Role: "member",
	}
	members := &client.Members{}
	for _, x := range team.Spec.Members {
		members.Members = append(members.Members, &client.Member{Name: x})
	}

	if _, err := s.Teams().Create(ctx, model, members); err != nil {
		log.WithFields(log.Fields{
			"error":     err.Error(),
			"name":      team.Name,
			"namespace": team.Namespace,
		}).Error("creating or updating team account")

		return nil, newError("creating or update team account", err).model()
	}

	return team, nil
}

// DeleteTeam is responsible for deleting a team in the org
func (s *serverImpl) DeleteTeam(ctx context.Context, namespace, name string) *models.APIError {
	log.WithFields(log.Fields{
		"name":      name,
		"namespace": namespace,
	}).Debug("attempting to delete the team")

	fullname := fmt.Sprintf("%s/%s", namespace, name)

	if found, err := s.Teams().Has(ctx, fullname); err != nil {
		return newError("checking team exists", err).model()
	} else if !found {
		return newError("resource does not exist", nil).model()
	}

	if err := s.Teams().Delete(ctx, fullname); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
			"name":  fullname,
		}).Error("deleting team")

		return newError("deleting team", err).model()
	}

	return nil
}

// GetTeam is responsible for retrieving the state of a team
func (s *serverImpl) GetTeam(ctx context.Context, namespace, name string) (*models.Team, *models.APIError) {
	fullname := fmt.Sprintf("%s/%s", namespace, name)

	if found, err := s.Teams().Has(ctx, fullname); err != nil {
		return nil, newError("checking team exists", err).model()
	} else if !found {
		return nil, newError("resource does not exist", nil).model()
	}

	// @step: we get the team members
	members, err := s.Teams().ListMembers(ctx, fullname)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
			"name":  fullname,
		}).Error("retrieving team members")

		return nil, newError("retrieving team members", err).model()
	}

	// @step: we fill in the spec and return
	resp := &models.Team{
		Object: models.Object{
			Name:      sp(name),
			Namespace: sp(namespace),
		},
		Spec: &models.TeamSpec{},
	}
	for _, x := range members.Members {
		resp.Spec.Members = append(resp.Spec.Members, x.Name)
	}

	return resp, nil
}

// ListTeams is responsible for listing the teams
func (s *serverImpl) ListTeams(ctx context.Context, namespace string) (*models.TeamList, *models.APIError) {
	teams, err := s.Teams().List(ctx, namespace)
	if err != nil {
		log.WithFields(log.Fields{
			"error":     err.Error(),
			"namespace": namespace,
		}).Error("listing teams")

		return nil, newError("listing teams", err).model()
	}

	list := &models.TeamList{
		Object: models.Object{
			Name:      sp("teams"),
			Namespace: sp(namespace),
		},
	}

	for _, x := range teams.Teams {
		team, err := s.GetTeam(ctx, namespace, x.Name)
		if err != nil {
			return nil, err
		}

		list.Items = append(list.Items, team)
	}

	return list, nil
}
