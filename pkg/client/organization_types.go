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

// Organization defines the organization
type Organization struct {
	Name                string           `json:"name"`
	Email               string           `json:"email"`
	InvoiceEmail        bool             `json:"invoice_email"`
	InvoiceEmailAddress interface{}      `json:"invoice_email_address"`
	IsAdmin             bool             `json:"is_admin"`
	IsFreeAccount       bool             `json:"is_free_account"`
	IsMember            bool             `json:"is_member"`
	OrderedTeams        []string         `json:"ordered_teams"`
	TagExpirationS      int              `json:"tag_expiration_s"`
	Teams               map[string]*Team `json:"teams"`
}
