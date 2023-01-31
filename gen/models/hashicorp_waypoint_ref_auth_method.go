// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointRefAuthMethod AuthMethod references an auth method.
//
// swagger:model hashicorp.waypoint.Ref.AuthMethod
type HashicorpWaypointRefAuthMethod struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this hashicorp waypoint ref auth method
func (m *HashicorpWaypointRefAuthMethod) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointRefAuthMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointRefAuthMethod) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointRefAuthMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
