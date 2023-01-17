// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointPipelineStepBuild hashicorp waypoint pipeline step build
//
// swagger:model hashicorp.waypoint.Pipeline.Step.Build
type HashicorpWaypointPipelineStepBuild struct {

	// Whether or not to push the built artifact to a remote container registry
	// TODO(briancain): ensure default to false because this will be inside
	// an ODR container
	DisablePush bool `json:"disable_push,omitempty"`
}

// Validate validates this hashicorp waypoint pipeline step build
func (m *HashicorpWaypointPipelineStepBuild) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointPipelineStepBuild) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointPipelineStepBuild) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointPipelineStepBuild
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
