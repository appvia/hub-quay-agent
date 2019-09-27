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

// ImageAnalysisSpec The resource specification for a image analysis
//
// swagger:model ImageAnalysisSpec
type ImageAnalysisSpec struct {

	// features
	Features []*ImageFeature `json:"features"`

	// The namespace catagory for the layer
	Namespace string `json:"namespace,omitempty"`

	// The status of the image analysis, which can be queued or scanned
	//
	// Required: true
	// Enum: [scanned queued]
	Status interface{} `json:"status"`

	// tag
	// Required: true
	Tag *RepositoryTag `json:"tag"`
}

// Validate validates this image analysis spec
func (m *ImageAnalysisSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageAnalysisSpec) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	for i := 0; i < len(m.Features); i++ {
		if swag.IsZero(m.Features[i]) { // not required
			continue
		}

		if m.Features[i] != nil {
			if err := m.Features[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("features" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var imageAnalysisSpecTypeStatusPropEnum []interface{}

func init() {
	var res []interface{}
	if err := json.Unmarshal([]byte(`["scanned","queued"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		imageAnalysisSpecTypeStatusPropEnum = append(imageAnalysisSpecTypeStatusPropEnum, v)
	}
}

// prop value enum
func (m *ImageAnalysisSpec) validateStatusEnum(path, location string, value interface{}) error {
	if err := validate.Enum(path, location, value, imageAnalysisSpecTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ImageAnalysisSpec) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *ImageAnalysisSpec) validateTag(formats strfmt.Registry) error {

	if err := validate.Required("tag", "body", m.Tag); err != nil {
		return err
	}

	if m.Tag != nil {
		if err := m.Tag.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tag")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageAnalysisSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageAnalysisSpec) UnmarshalBinary(b []byte) error {
	var res ImageAnalysisSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}