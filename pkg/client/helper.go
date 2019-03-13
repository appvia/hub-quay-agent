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

import "strings"

func parseRobotName(name string) (string, string) {
	items := strings.Split(name, "+")
	if len(items) == 2 {
		return items[0], items[1]
	}

	return "", name
}

func robotShortName(name string) string {
	_, shortname := parseRobotName(name)

	return shortname
}

func hasPermission(list []*Permission, perm *Permission) bool {
	for _, x := range list {
		if x.Name == perm.Name && x.Role == perm.Role {
			return true
		}
	}

	return false
}
