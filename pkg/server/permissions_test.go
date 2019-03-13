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
	"testing"

	"github.com/appvia/hub-quay-agent/pkg/client"
	"github.com/appvia/hub-quay-agent/pkg/transport/models"

	"github.com/stretchr/testify/assert"
)

func TestCreatePermissions(t *testing.T) {
	cs := []struct {
		Members  []*client.Permission
		Expected []*models.Permission
		Add      []*client.Permission
		Remove   []*client.Permission
	}{
		{
			Members: []*client.Permission{
				{Name: "test", Role: "read"},
			},
			Expected: []*models.Permission{
				{Name: sp("test"), Permission: sp("read")},
			},
		},
		{
			Members: []*client.Permission{},
			Expected: []*models.Permission{
				{Name: sp("test"), Permission: sp("read")},
			},
			Add: []*client.Permission{
				{Name: "test", Role: "read"},
			},
		},
		{
			Members: []*client.Permission{
				{Name: "test1", Role: "read"},
			},
			Expected: []*models.Permission{
				{Name: sp("test"), Permission: sp("read")},
			},
			Add: []*client.Permission{
				{Name: "test", Role: "read"},
			},
			Remove: []*client.Permission{
				{Name: "test1", Role: "read"},
			},
		},
		{
			Members: []*client.Permission{
				{Name: "test", Role: "read"},
			},
			Expected: []*models.Permission{
				{Name: sp("test"), Permission: sp("write")},
			},
			Add: []*client.Permission{
				{Name: "test", Role: "write"},
			},
		},
	}
	for _, x := range cs {
		add, remove := createPermissions(x.Members, x.Expected)
		assert.Equal(t, x.Add, add)
		assert.Equal(t, x.Remove, remove)
	}
}
