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

// HashicorpWaypointJobState hashicorp waypoint job state
//
// swagger:model hashicorp.waypoint.Job.State
type HashicorpWaypointJobState string

const (

	// HashicorpWaypointJobStateUNKNOWN captures enum value "UNKNOWN"
	HashicorpWaypointJobStateUNKNOWN HashicorpWaypointJobState = "UNKNOWN"

	// HashicorpWaypointJobStateQUEUED captures enum value "QUEUED"
	HashicorpWaypointJobStateQUEUED HashicorpWaypointJobState = "QUEUED"

	// HashicorpWaypointJobStateWAITING captures enum value "WAITING"
	HashicorpWaypointJobStateWAITING HashicorpWaypointJobState = "WAITING"

	// HashicorpWaypointJobStateRUNNING captures enum value "RUNNING"
	HashicorpWaypointJobStateRUNNING HashicorpWaypointJobState = "RUNNING"

	// HashicorpWaypointJobStateERROR captures enum value "ERROR"
	HashicorpWaypointJobStateERROR HashicorpWaypointJobState = "ERROR"

	// HashicorpWaypointJobStateSUCCESS captures enum value "SUCCESS"
	HashicorpWaypointJobStateSUCCESS HashicorpWaypointJobState = "SUCCESS"
)

// for schema
var hashicorpWaypointJobStateEnum []interface{}

func init() {
	var res []HashicorpWaypointJobState
	if err := json.Unmarshal([]byte(`["UNKNOWN","QUEUED","WAITING","RUNNING","ERROR","SUCCESS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpWaypointJobStateEnum = append(hashicorpWaypointJobStateEnum, v)
	}
}

func (m HashicorpWaypointJobState) validateHashicorpWaypointJobStateEnum(path, location string, value HashicorpWaypointJobState) error {
	if err := validate.EnumCase(path, location, value, hashicorpWaypointJobStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp waypoint job state
func (m HashicorpWaypointJobState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpWaypointJobStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
