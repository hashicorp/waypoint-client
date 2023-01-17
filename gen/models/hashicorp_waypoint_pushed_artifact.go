// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointPushedArtifact hashicorp waypoint pushed artifact
//
// swagger:model hashicorp.waypoint.PushedArtifact
type HashicorpWaypointPushedArtifact struct {

	// application that this belongs to
	Application *HashicorpWaypointRefApplication `json:"application,omitempty"`

	// artifact is the artifact that was a result from the push.
	Artifact *HashicorpWaypointArtifact `json:"artifact,omitempty"`

	// If include_build was set on the list request, this will include
	// the Build value associated with the given build_id.
	Build *HashicorpWaypointBuild `json:"build,omitempty"`

	// the id of the build that this pushed artifact was sourced from.
	BuildID string `json:"build_id,omitempty"`

	// component that pushed this artifact
	Component *HashicorpWaypointComponent `json:"component,omitempty"`

	// id is a unique ID for this push
	ID string `json:"id,omitempty"`

	// ID of the job that created this. This may be empty.
	JobID string `json:"job_id,omitempty"`

	// labels are the set of labels that are present on this build.
	Labels map[string]string `json:"labels,omitempty"`

	// Preload is information that is available via further queries but it
	// sometimes pre-populated in the initial load (see the field docs for more
	// info).
	Preload *HashicorpWaypointPushedArtifactPreload `json:"preload,omitempty"`

	// The sequence number for this build.
	Sequence string `json:"sequence,omitempty"`

	// status of the push operation
	Status *HashicorpWaypointStatus `json:"status,omitempty"`

	// template data for HCL variables and template functions, json-encoded
	// Format: byte
	TemplateData strfmt.Base64 `json:"template_data,omitempty"`

	// The workspace that this exists in
	Workspace *HashicorpWaypointRefWorkspace `json:"workspace,omitempty"`
}

// Validate validates this hashicorp waypoint pushed artifact
func (m *HashicorpWaypointPushedArtifact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArtifact(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuild(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreload(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointPushedArtifact) validateApplication(formats strfmt.Registry) error {

	if swag.IsZero(m.Application) { // not required
		return nil
	}

	if m.Application != nil {
		if err := m.Application.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointPushedArtifact) validateArtifact(formats strfmt.Registry) error {

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

func (m *HashicorpWaypointPushedArtifact) validateBuild(formats strfmt.Registry) error {

	if swag.IsZero(m.Build) { // not required
		return nil
	}

	if m.Build != nil {
		if err := m.Build.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("build")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointPushedArtifact) validateComponent(formats strfmt.Registry) error {

	if swag.IsZero(m.Component) { // not required
		return nil
	}

	if m.Component != nil {
		if err := m.Component.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointPushedArtifact) validatePreload(formats strfmt.Registry) error {

	if swag.IsZero(m.Preload) { // not required
		return nil
	}

	if m.Preload != nil {
		if err := m.Preload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preload")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointPushedArtifact) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointPushedArtifact) validateWorkspace(formats strfmt.Registry) error {

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
func (m *HashicorpWaypointPushedArtifact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointPushedArtifact) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointPushedArtifact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
