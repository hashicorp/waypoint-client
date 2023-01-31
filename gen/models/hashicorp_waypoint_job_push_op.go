// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointJobPushOp hashicorp waypoint job push op
//
// swagger:model hashicorp.waypoint.Job.PushOp
type HashicorpWaypointJobPushOp struct {

	// Build to push
	Build *HashicorpWaypointBuild `json:"build,omitempty"`
}

// Validate validates this hashicorp waypoint job push op
func (m *HashicorpWaypointJobPushOp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuild(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointJobPushOp) validateBuild(formats strfmt.Registry) error {

	if swag.IsZero(m.Build) { // not required
		return nil
	}

	if m.Build != nil {
		if err := m.Build.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("build")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointJobPushOp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointJobPushOp) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointJobPushOp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
