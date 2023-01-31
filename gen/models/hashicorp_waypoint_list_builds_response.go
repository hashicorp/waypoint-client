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

// HashicorpWaypointListBuildsResponse hashicorp waypoint list builds response
//
// swagger:model hashicorp.waypoint.ListBuildsResponse
type HashicorpWaypointListBuildsResponse struct {

	// builds is the list of builds.
	Builds []*HashicorpWaypointBuild `json:"builds"`
}

// Validate validates this hashicorp waypoint list builds response
func (m *HashicorpWaypointListBuildsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuilds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointListBuildsResponse) validateBuilds(formats strfmt.Registry) error {

	if swag.IsZero(m.Builds) { // not required
		return nil
	}

	for i := 0; i < len(m.Builds); i++ {
		if swag.IsZero(m.Builds[i]) { // not required
			continue
		}

		if m.Builds[i] != nil {
			if err := m.Builds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("builds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointListBuildsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointListBuildsResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointListBuildsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
