// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointValidateJobRequest hashicorp waypoint validate job request
//
// swagger:model hashicorp.waypoint.ValidateJobRequest
type HashicorpWaypointValidateJobRequest struct {

	// If true, will NOT validate that the job is assignable.
	DisableAssign bool `json:"disable_assign,omitempty"`

	// The job to validate.
	Job *HashicorpWaypointJob `json:"job,omitempty"`
}

// Validate validates this hashicorp waypoint validate job request
func (m *HashicorpWaypointValidateJobRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJob(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointValidateJobRequest) validateJob(formats strfmt.Registry) error {

	if swag.IsZero(m.Job) { // not required
		return nil
	}

	if m.Job != nil {
		if err := m.Job.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("job")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointValidateJobRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointValidateJobRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointValidateJobRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
