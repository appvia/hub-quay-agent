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
	"strings"
)

type teamsImpl struct {
	Client
}

// Get is responsible for retrieving the team
func (t *teamsImpl) Get(ctx context.Context, fullname string) (*Team, error) {
	if found, err := t.Has(ctx, fullname); err != nil {
		return nil, err
	} else if !found {
		return nil, errors.New("team does not exist in organization")
	}
	namespace, name := t.parseFullName(fullname)

	teams, err := t.List(ctx, namespace)
	if err != nil {
		return nil, err
	}

	for _, x := range teams.Teams {
		if name == x.Name {
			return x, nil
		}
	}

	// @potentially the situation has changed between now and the has()
	return nil, errors.New("team does not exist in organization")
}

// Create is used to create a team
func (t *teamsImpl) Create(ctx context.Context, team *Team, members *Members) (*Team, error) {
	namespace, name := t.parseFullName(team.Name)

	found, err := t.Has(ctx, team.Name)
	if err != nil {
		return nil, err
	}
	if !found {
		// @step: we are creating a new team
		// PUT /api/v1/organization/{orgname}/team/{teamname}
		uri := fmt.Sprintf("organization/%s/team/%s", namespace, name)
		if err := t.Handle(ctx, http.MethodPut, uri, nil, nil); err != nil {
			return nil, err
		}
		// @step: iterate the members and add them
		for _, x := range members.Members {
			uri = fmt.Sprintf("organization/%s/team/%s/members/%s", namespace, name, x.Name)
			if err := t.Handle(ctx, http.MethodPut, uri, nil, nil); err != nil {
				return nil, err
			}
		}

		return team, nil
	}

	// @step: else we are updating an existing team in the org
	current := &Members{}
	uri := fmt.Sprintf("organization/%s/team/%s/members", namespace, name)

	if err := t.Handle(ctx, http.MethodGet, uri, nil, current); err != nil {
		return nil, err
	}

	exists := func(member string, list []*Member) bool {
		for _, x := range list {
			if x.Name == member {
				return true
			}
		}
		return false
	}

	// @step: iterate the list and delete or add those whom are
	// supposed to be there
	for _, x := range members.Members {
		if !exists(x.Name, current.Members) {
			// PUT /api/v1/organization/{orgname}/team/{teamname}/members/{membername}
			uri := fmt.Sprintf("organization/%s/team/%s/members/%s", namespace, name, x.Name)
			if err := t.Handle(ctx, http.MethodPut, uri, nil, nil); err != nil {
				return nil, err
			}
		}
	}
	for _, x := range current.Members {
		if !exists(x.Name, members.Members) {
			uri := fmt.Sprintf("organization/%s/team/%s/members/%s", namespace, name, x.Name)
			if err := t.Handle(ctx, http.MethodDelete, uri, nil, nil); err != nil {
				return nil, err
			}
		}
	}

	return team, nil
}

// Delete is responsible for deleting a team from the org
func (t *teamsImpl) Delete(ctx context.Context, fullname string) error {
	if found, err := t.Has(ctx, fullname); err != nil {
		return err
	} else if !found {
		return nil
	}

	namespace, name := t.parseFullName(fullname)
	uri := fmt.Sprintf("organization/%s/team/%s", namespace, name)

	return t.Handle(ctx, http.MethodDelete, uri, nil, nil)
}

// Has checks if the team exists
func (t *teamsImpl) Has(ctx context.Context, fullname string) (bool, error) {
	namespace, name := t.parseFullName(fullname)

	teams, err := t.List(ctx, namespace)
	if err != nil {
		return false, err
	}

	for _, x := range teams.Teams {
		if x.Name == name {
			return true, nil
		}
	}

	return false, nil
}

// List is responsible for listing all the teams in the organization
func (t *teamsImpl) List(ctx context.Context, namespace string) (*TeamList, error) {
	// @step we need to query the organization whih has a list of teams
	org := &Organization{}
	uri := fmt.Sprintf("organization/%s", namespace)

	if err := t.Handle(ctx, http.MethodGet, uri, nil, org); err != nil {
		return nil, err
	}

	list := &TeamList{}

	for _, x := range org.Teams {
		list.Teams = append(list.Teams, x)
	}

	return list, nil
}

func (t teamsImpl) parseFullName(fullname string) (string, string) {
	items := strings.Split(fullname, "/")
	if len(items) == 2 {
		return items[0], items[1]
	}

	return "", fullname
}
