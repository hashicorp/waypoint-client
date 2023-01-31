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

// HashicorpWaypointVariableFinalValueSource  - UNKNOWN: This is the final used value's source
// that is supplied to the user in outputs.
//
// swagger:model hashicorp.waypoint.Variable.FinalValue.Source
type HashicorpWaypointVariableFinalValueSource string

const (

	// HashicorpWaypointVariableFinalValueSourceUNKNOWN captures enum value "UNKNOWN"
	HashicorpWaypointVariableFinalValueSourceUNKNOWN HashicorpWaypointVariableFinalValueSource = "UNKNOWN"

	// HashicorpWaypointVariableFinalValueSourceDEFAULT captures enum value "DEFAULT"
	HashicorpWaypointVariableFinalValueSourceDEFAULT HashicorpWaypointVariableFinalValueSource = "DEFAULT"

	// HashicorpWaypointVariableFinalValueSourceFILE captures enum value "FILE"
	HashicorpWaypointVariableFinalValueSourceFILE HashicorpWaypointVariableFinalValueSource = "FILE"

	// HashicorpWaypointVariableFinalValueSourceCLI captures enum value "CLI"
	HashicorpWaypointVariableFinalValueSourceCLI HashicorpWaypointVariableFinalValueSource = "CLI"

	// HashicorpWaypointVariableFinalValueSourceENV captures enum value "ENV"
	HashicorpWaypointVariableFinalValueSourceENV HashicorpWaypointVariableFinalValueSource = "ENV"

	// HashicorpWaypointVariableFinalValueSourceVCS captures enum value "VCS"
	HashicorpWaypointVariableFinalValueSourceVCS HashicorpWaypointVariableFinalValueSource = "VCS"

	// HashicorpWaypointVariableFinalValueSourceSERVER captures enum value "SERVER"
	HashicorpWaypointVariableFinalValueSourceSERVER HashicorpWaypointVariableFinalValueSource = "SERVER"

	// HashicorpWaypointVariableFinalValueSourceDYNAMIC captures enum value "DYNAMIC"
	HashicorpWaypointVariableFinalValueSourceDYNAMIC HashicorpWaypointVariableFinalValueSource = "DYNAMIC"
)

// for schema
var hashicorpWaypointVariableFinalValueSourceEnum []interface{}

func init() {
	var res []HashicorpWaypointVariableFinalValueSource
	if err := json.Unmarshal([]byte(`["UNKNOWN","DEFAULT","FILE","CLI","ENV","VCS","SERVER","DYNAMIC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpWaypointVariableFinalValueSourceEnum = append(hashicorpWaypointVariableFinalValueSourceEnum, v)
	}
}

func (m HashicorpWaypointVariableFinalValueSource) validateHashicorpWaypointVariableFinalValueSourceEnum(path, location string, value HashicorpWaypointVariableFinalValueSource) error {
	if err := validate.EnumCase(path, location, value, hashicorpWaypointVariableFinalValueSourceEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp waypoint variable final value source
func (m HashicorpWaypointVariableFinalValueSource) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpWaypointVariableFinalValueSourceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
