// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointRefDeclaredResource DeclaredResource references a declared resource.
//
// swagger:model hashicorp.waypoint.Ref.DeclaredResource
type HashicorpWaypointRefDeclaredResource struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this hashicorp waypoint ref declared resource
func (m *HashicorpWaypointRefDeclaredResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointRefDeclaredResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointRefDeclaredResource) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointRefDeclaredResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
