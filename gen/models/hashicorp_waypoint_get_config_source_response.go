// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointGetConfigSourceResponse hashicorp waypoint get config source response
//
// swagger:model hashicorp.waypoint.GetConfigSourceResponse
type HashicorpWaypointGetConfigSourceResponse struct {

	// config sources
	ConfigSources []*HashicorpWaypointConfigSource `json:"config_sources"`
}

// Validate validates this hashicorp waypoint get config source response
func (m *HashicorpWaypointGetConfigSourceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigSources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointGetConfigSourceResponse) validateConfigSources(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfigSources) { // not required
		return nil
	}

	for i := 0; i < len(m.ConfigSources); i++ {
		if swag.IsZero(m.ConfigSources[i]) { // not required
			continue
		}

		if m.ConfigSources[i] != nil {
			if err := m.ConfigSources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("config_sources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointGetConfigSourceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointGetConfigSourceResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointGetConfigSourceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
