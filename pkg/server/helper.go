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
	"strings"
)

// filterRobotName removes the org from the robot name
func filterRobotName(name string) string {
	e := strings.Split(name, "+")
	if len(e) == 2 {
		return e[1]
	}

	return name
}

// ParseName extracts the namespace and image
func ParseName(name string) (string, string) {
	items := strings.Split(name, "/")
	if len(items) == 2 {
		return items[0], items[1]
	}
	items = strings.Split(name, "+")
	if len(items) == 2 {
		return items[0], items[1]
	}

	return "", name
}

func sp(v string) *string {
	return &v
}

// sv returns the dereferenced value
func sv(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}
