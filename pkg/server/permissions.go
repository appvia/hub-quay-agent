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
	"github.com/appvia/hub-quay-agent/pkg/client"
	"github.com/appvia/hub-quay-agent/pkg/transport/models"
)

// createPermissions performs a diff on the users and hands back the add and removes
func createPermissions(members []*client.Permission, expected []*models.Permission) ([]*client.Permission, []*client.Permission) {
	var add, remove []*client.Permission

	// @step: find any permissions that need removing
	for _, x := range members {
		if !hasModelPermission(x.Name, expected) {
			remove = append(remove, x)
		}
	}

	// @step: find any permissions which need adding or changing
	for _, x := range expected {
		if !hasClientPermission(sv(x.Name), sv(x.Permission), members) {
			add = append(add, &client.Permission{Name: sv(x.Name), Role: sv(x.Permission)})
		}
	}

	return add, remove
}

func hasModelPermission(name string, permissions []*models.Permission) bool {
	for _, x := range permissions {
		if name == sv(x.Name) {
			return true
		}
	}

	return false
}

func hasClientPermission(name, access string, permissions []*client.Permission) bool {
	for _, x := range permissions {
		if name == x.Name && access == x.Role {
			return true
		}
	}

	return false
}
