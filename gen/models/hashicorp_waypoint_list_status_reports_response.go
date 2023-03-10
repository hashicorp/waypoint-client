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

// HashicorpWaypointListStatusReportsResponse hashicorp waypoint list status reports response
//
// swagger:model hashicorp.waypoint.ListStatusReportsResponse
type HashicorpWaypointListStatusReportsResponse struct {

	// status reports
	StatusReports []*HashicorpWaypointStatusReport `json:"status_reports"`
}

// Validate validates this hashicorp waypoint list status reports response
func (m *HashicorpWaypointListStatusReportsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatusReports(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointListStatusReportsResponse) validateStatusReports(formats strfmt.Registry) error {

	if swag.IsZero(m.StatusReports) { // not required
		return nil
	}

	for i := 0; i < len(m.StatusReports); i++ {
		if swag.IsZero(m.StatusReports[i]) { // not required
			continue
		}

		if m.StatusReports[i] != nil {
			if err := m.StatusReports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status_reports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointListStatusReportsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointListStatusReportsResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointListStatusReportsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
