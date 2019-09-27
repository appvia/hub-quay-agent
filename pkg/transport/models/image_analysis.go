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
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ImageAnalysis The resource definition for a list of vulnerabilities on a image tag
//
// swagger:model ImageAnalysis
type ImageAnalysis struct {
	Object

	// spec
	// Required: true
	Spec *ImageAnalysisSpec `json:"spec"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ImageAnalysis) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Object
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Object = aO0

	// now for regular properties
	var propsImageAnalysis struct {
		Spec *ImageAnalysisSpec `json:"spec"`
	}
	if err := swag.ReadJSON(raw, &propsImageAnalysis); err != nil {
		return err
	}
	m.Spec = propsImageAnalysis.Spec

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ImageAnalysis) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.Object)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsImageAnalysis struct {
		Spec *ImageAnalysisSpec `json:"spec"`
	}
	propsImageAnalysis.Spec = m.Spec

	jsonDataPropsImageAnalysis, errImageAnalysis := swag.WriteJSON(propsImageAnalysis)
	if errImageAnalysis != nil {
		return nil, errImageAnalysis
	}
	_parts = append(_parts, jsonDataPropsImageAnalysis)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this image analysis
func (m *ImageAnalysis) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Object
	if err := m.Object.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageAnalysis) validateSpec(formats strfmt.Registry) error {

	if err := validate.Required("spec", "body", m.Spec); err != nil {
		return err
	}

	if m.Spec != nil {
		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageAnalysis) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageAnalysis) UnmarshalBinary(b []byte) error {
	var res ImageAnalysis
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}