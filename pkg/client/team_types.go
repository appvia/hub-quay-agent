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

// TeamList is a list of robots
type TeamList struct {
	Teams []*Team `json:"teams,omitempty"`
}

// Team defines a team configuration
type Team struct {
	IsSynced    bool   `json:"is_synced"`
	CanView     bool   `json:"can_view"`
	Role        string `json:"role"`
	Name        string `json:"name"`
	MemberCount int    `json:"member_count"`
	RepoCount   int    `json:"repo_count"`
	Description string `json:"description"`
}

// Members is a collection of members
type Members struct {
	CanEdit bool      `json:"can_edit"`
	Members []*Member `json:"members"`
	Name    string    `json:"name"`
}

// Member is a member of a team
type Member struct {
	Invited bool   `json:"invited"`
	Kind    string `json:"kind"`
	Name    string `json:"name"`
	IsRobot bool   `json:"is_robot"`
}
