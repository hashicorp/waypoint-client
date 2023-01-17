// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointJobDestroyProjectOp DestroyProjectOp triggers the deletion of a project from the database
// as well as (optionally) the destruction of all resources created within
// a project
//
// swagger:model hashicorp.waypoint.Job.DestroyProjectOp
type HashicorpWaypointJobDestroyProjectOp struct {

	// project
	Project *HashicorpWaypointRefProject `json:"project,omitempty"`

	// skip destroy resources
	SkipDestroyResources bool `json:"skip_destroy_resources,omitempty"`
}

// Validate validates this hashicorp waypoint job destroy project op
func (m *HashicorpWaypointJobDestroyProjectOp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointJobDestroyProjectOp) validateProject(formats strfmt.Registry) error {

	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointJobDestroyProjectOp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointJobDestroyProjectOp) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointJobDestroyProjectOp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
