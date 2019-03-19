// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright (C) 2019  Rohith Jayawardene <gambol99@gmail.com>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RepositorySpec The definitions for a repository
// swagger:model RepositorySpec
type RepositorySpec struct {

	// A description about what the image is used for
	Description string `json:"description,omitempty"`

	// A list of members whom has access to the repository
	Members []*Permission `json:"members"`

	// A list of robot accounts who access to the repository
	Robots []*Permission `json:"robots"`

	// A collection of tags associated to the image
	Tags map[string]string `json:"tags,omitempty"`

	// The docker pull url for this image
	URL string `json:"url,omitempty"`

	// The visibility of the repository in the registry
	// Enum: [internal private public]
	Visibility string `json:"visibility,omitempty"`
}

// Validate validates this repository spec
func (m *RepositorySpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMembers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRobots(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVisibility(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RepositorySpec) validateMembers(formats strfmt.Registry) error {

	if swag.IsZero(m.Members) { // not required
		return nil
	}

	for i := 0; i < len(m.Members); i++ {
		if swag.IsZero(m.Members[i]) { // not required
			continue
		}

		if m.Members[i] != nil {
			if err := m.Members[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RepositorySpec) validateRobots(formats strfmt.Registry) error {

	if swag.IsZero(m.Robots) { // not required
		return nil
	}

	for i := 0; i < len(m.Robots); i++ {
		if swag.IsZero(m.Robots[i]) { // not required
			continue
		}

		if m.Robots[i] != nil {
			if err := m.Robots[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("robots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var repositorySpecTypeVisibilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["internal","private","public"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		repositorySpecTypeVisibilityPropEnum = append(repositorySpecTypeVisibilityPropEnum, v)
	}
}

const (

	// RepositorySpecVisibilityInternal captures enum value "internal"
	RepositorySpecVisibilityInternal string = "internal"

	// RepositorySpecVisibilityPrivate captures enum value "private"
	RepositorySpecVisibilityPrivate string = "private"

	// RepositorySpecVisibilityPublic captures enum value "public"
	RepositorySpecVisibilityPublic string = "public"
)

// prop value enum
func (m *RepositorySpec) validateVisibilityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, repositorySpecTypeVisibilityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RepositorySpec) validateVisibility(formats strfmt.Registry) error {

	if swag.IsZero(m.Visibility) { // not required
		return nil
	}

	// value enum
	if err := m.validateVisibilityEnum("visibility", "body", m.Visibility); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RepositorySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RepositorySpec) UnmarshalBinary(b []byte) error {
	var res RepositorySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
