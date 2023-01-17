// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// HashicorpWaypointJobConfigSource Source is the location where the configuration was loaded from.
//
//  - UNKNOWN: Unknown should never be set, but represents a zero value.
//  - FILE: File is when the waypoint.hcl was loaded from a file either
// on disk (local actions) or the attached repository (GitOps).
//  - SERVER: Server is when the waypoint.hcl was loaded from the server
// from being written directly in the project settings.
//  - JOB: Job is when the waypoint.hcl was loaded directly from the job by
// being embedded in the "waypoint_hcl" field (tag 12).
//
// swagger:model hashicorp.waypoint.Job.Config.Source
type HashicorpWaypointJobConfigSource string

const (

	// HashicorpWaypointJobConfigSourceUNKNOWN captures enum value "UNKNOWN"
	HashicorpWaypointJobConfigSourceUNKNOWN HashicorpWaypointJobConfigSource = "UNKNOWN"

	// HashicorpWaypointJobConfigSourceFILE captures enum value "FILE"
	HashicorpWaypointJobConfigSourceFILE HashicorpWaypointJobConfigSource = "FILE"

	// HashicorpWaypointJobConfigSourceSERVER captures enum value "SERVER"
	HashicorpWaypointJobConfigSourceSERVER HashicorpWaypointJobConfigSource = "SERVER"

	// HashicorpWaypointJobConfigSourceJOB captures enum value "JOB"
	HashicorpWaypointJobConfigSourceJOB HashicorpWaypointJobConfigSource = "JOB"
)

// for schema
var hashicorpWaypointJobConfigSourceEnum []interface{}

func init() {
	var res []HashicorpWaypointJobConfigSource
	if err := json.Unmarshal([]byte(`["UNKNOWN","FILE","SERVER","JOB"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpWaypointJobConfigSourceEnum = append(hashicorpWaypointJobConfigSourceEnum, v)
	}
}

func (m HashicorpWaypointJobConfigSource) validateHashicorpWaypointJobConfigSourceEnum(path, location string, value HashicorpWaypointJobConfigSource) error {
	if err := validate.EnumCase(path, location, value, hashicorpWaypointJobConfigSourceEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp waypoint job config source
func (m HashicorpWaypointJobConfigSource) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpWaypointJobConfigSourceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
