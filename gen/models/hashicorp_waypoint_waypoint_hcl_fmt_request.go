// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointWaypointHclFmtRequest hashicorp waypoint waypoint hcl fmt request
//
// swagger:model hashicorp.waypoint.WaypointHclFmtRequest
type HashicorpWaypointWaypointHclFmtRequest struct {

	// waypoint hcl
	// Format: byte
	WaypointHcl strfmt.Base64 `json:"waypoint_hcl,omitempty"`
}

// Validate validates this hashicorp waypoint waypoint hcl fmt request
func (m *HashicorpWaypointWaypointHclFmtRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointWaypointHclFmtRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointWaypointHclFmtRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointWaypointHclFmtRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
