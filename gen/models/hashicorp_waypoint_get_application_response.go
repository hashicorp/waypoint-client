// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointGetApplicationResponse hashicorp waypoint get application response
//
// swagger:model hashicorp.waypoint.GetApplicationResponse
type HashicorpWaypointGetApplicationResponse struct {

	// application
	Application *HashicorpWaypointApplication `json:"application,omitempty"`
}

// Validate validates this hashicorp waypoint get application response
func (m *HashicorpWaypointGetApplicationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointGetApplicationResponse) validateApplication(formats strfmt.Registry) error {

	if swag.IsZero(m.Application) { // not required
		return nil
	}

	if m.Application != nil {
		if err := m.Application.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointGetApplicationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointGetApplicationResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointGetApplicationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
