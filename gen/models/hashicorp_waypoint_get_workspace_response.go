// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointGetWorkspaceResponse hashicorp waypoint get workspace response
//
// swagger:model hashicorp.waypoint.GetWorkspaceResponse
type HashicorpWaypointGetWorkspaceResponse struct {

	// workspace
	Workspace *HashicorpWaypointWorkspace `json:"workspace,omitempty"`
}

// Validate validates this hashicorp waypoint get workspace response
func (m *HashicorpWaypointGetWorkspaceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWorkspace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointGetWorkspaceResponse) validateWorkspace(formats strfmt.Registry) error {

	if swag.IsZero(m.Workspace) { // not required
		return nil
	}

	if m.Workspace != nil {
		if err := m.Workspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointGetWorkspaceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointGetWorkspaceResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointGetWorkspaceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
