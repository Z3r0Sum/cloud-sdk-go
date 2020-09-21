// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeploymentCreateSettings Additional configuration for the new deployment object.
//
// swagger:model DeploymentCreateSettings
type DeploymentCreateSettings struct {

	// DEPRECATED (Scheduled to be removed in the next major version): The set of rulesets applies to this deployment.
	IPFilteringSettings *IPFilteringSettings `json:"ip_filtering_settings,omitempty"`

	// Observability settings for this deployment
	Observability *DeploymentObservabilitySettings `json:"observability,omitempty"`

	// The traffic filter rulesets to apply to this deployment.
	TrafficFilterSettings *TrafficFilterSettings `json:"traffic_filter_settings,omitempty"`
}

// Validate validates this deployment create settings
func (m *DeploymentCreateSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPFilteringSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObservability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrafficFilterSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentCreateSettings) validateIPFilteringSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.IPFilteringSettings) { // not required
		return nil
	}

	if m.IPFilteringSettings != nil {
		if err := m.IPFilteringSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip_filtering_settings")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentCreateSettings) validateObservability(formats strfmt.Registry) error {

	if swag.IsZero(m.Observability) { // not required
		return nil
	}

	if m.Observability != nil {
		if err := m.Observability.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("observability")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentCreateSettings) validateTrafficFilterSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.TrafficFilterSettings) { // not required
		return nil
	}

	if m.TrafficFilterSettings != nil {
		if err := m.TrafficFilterSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("traffic_filter_settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentCreateSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentCreateSettings) UnmarshalBinary(b []byte) error {
	var res DeploymentCreateSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
