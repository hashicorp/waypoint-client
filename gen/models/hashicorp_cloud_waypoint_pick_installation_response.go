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

// HashicorpCloudWaypointPickInstallationResponse hashicorp cloud waypoint pick installation response
//
// swagger:model hashicorp.cloud.waypoint.PickInstallationResponse
type HashicorpCloudWaypointPickInstallationResponse struct {

	// installations
	Installations []*HashicorpCloudWaypointInstallation `json:"installations"`
}

// Validate validates this hashicorp cloud waypoint pick installation response
func (m *HashicorpCloudWaypointPickInstallationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstallations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointPickInstallationResponse) validateInstallations(formats strfmt.Registry) error {

	if swag.IsZero(m.Installations) { // not required
		return nil
	}

	for i := 0; i < len(m.Installations); i++ {
		if swag.IsZero(m.Installations[i]) { // not required
			continue
		}

		if m.Installations[i] != nil {
			if err := m.Installations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("installations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointPickInstallationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointPickInstallationResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointPickInstallationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
