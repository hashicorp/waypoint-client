// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointArtifact Artifact is the result of a build or registry. This is the metadata only.
// The binary contents of an artifact are expected to be stored in a registry.
//
// swagger:model hashicorp.waypoint.Artifact
type HashicorpWaypointArtifact struct {

	// artifact is the full artifact encoded directly from the component plugin.
	// The receiving end must have access to the component proto files to
	// know how to decode this.
	Artifact *OpaqueanyAny `json:"artifact,omitempty"`

	// This is the JSON-encoded protobuf structure of the Any field above.
	// This is generated by the plugin and Waypoint core does not modify this
	// value or have any enforced structure. This will be different per-plugin.
	ArtifactJSON string `json:"artifact_json,omitempty"`
}

// Validate validates this hashicorp waypoint artifact
func (m *HashicorpWaypointArtifact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArtifact(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointArtifact) validateArtifact(formats strfmt.Registry) error {

	if swag.IsZero(m.Artifact) { // not required
		return nil
	}

	if m.Artifact != nil {
		if err := m.Artifact.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifact")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointArtifact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointArtifact) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointArtifact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
