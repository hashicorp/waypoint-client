// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointReleasePreload hashicorp waypoint release preload
//
// swagger:model hashicorp.waypoint.Release.Preload
type HashicorpWaypointReleasePreload struct {

	// Populated when LoadDetails is set.
	Artifact *HashicorpWaypointPushedArtifact `json:"artifact,omitempty"`

	// Populated when LoadDetails is set.
	Build *HashicorpWaypointBuild `json:"build,omitempty"`

	// Populated when LoadDetails is set.
	Deployment *HashicorpWaypointDeployment `json:"deployment,omitempty"`

	// The ref that was used in the job to run this operation. This is
	// also accessible by querying the job via the job_id and should always
	// match.
	//
	// This may be null under multiple circumstances: (1) the job was
	// manually triggered with local data (no datasource) or (2) the job
	// was run in earlier versions of Waypoint before we tracked this or
	// (3) the job hasn't yet loaded the data.
	//
	// This is always pre-populated if it is exists.
	JobDataSourceRef *HashicorpWaypointJobDataSourceRef `json:"job_data_source_ref,omitempty"`
}

// Validate validates this hashicorp waypoint release preload
func (m *HashicorpWaypointReleasePreload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArtifact(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuild(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeployment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobDataSourceRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointReleasePreload) validateArtifact(formats strfmt.Registry) error {

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

func (m *HashicorpWaypointReleasePreload) validateBuild(formats strfmt.Registry) error {

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

func (m *HashicorpWaypointReleasePreload) validateDeployment(formats strfmt.Registry) error {

	if swag.IsZero(m.Deployment) { // not required
		return nil
	}

	if m.Deployment != nil {
		if err := m.Deployment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deployment")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointReleasePreload) validateJobDataSourceRef(formats strfmt.Registry) error {

	if swag.IsZero(m.JobDataSourceRef) { // not required
		return nil
	}

	if m.JobDataSourceRef != nil {
		if err := m.JobDataSourceRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("job_data_source_ref")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointReleasePreload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointReleasePreload) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointReleasePreload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
