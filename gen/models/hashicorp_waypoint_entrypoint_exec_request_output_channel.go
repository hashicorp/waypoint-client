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

// HashicorpWaypointEntrypointExecRequestOutputChannel hashicorp waypoint entrypoint exec request output channel
//
// swagger:model hashicorp.waypoint.EntrypointExecRequest.Output.Channel
type HashicorpWaypointEntrypointExecRequestOutputChannel string

const (

	// HashicorpWaypointEntrypointExecRequestOutputChannelUNKNOWN captures enum value "UNKNOWN"
	HashicorpWaypointEntrypointExecRequestOutputChannelUNKNOWN HashicorpWaypointEntrypointExecRequestOutputChannel = "UNKNOWN"

	// HashicorpWaypointEntrypointExecRequestOutputChannelSTDOUT captures enum value "STDOUT"
	HashicorpWaypointEntrypointExecRequestOutputChannelSTDOUT HashicorpWaypointEntrypointExecRequestOutputChannel = "STDOUT"

	// HashicorpWaypointEntrypointExecRequestOutputChannelSTDERR captures enum value "STDERR"
	HashicorpWaypointEntrypointExecRequestOutputChannelSTDERR HashicorpWaypointEntrypointExecRequestOutputChannel = "STDERR"
)

// for schema
var hashicorpWaypointEntrypointExecRequestOutputChannelEnum []interface{}

func init() {
	var res []HashicorpWaypointEntrypointExecRequestOutputChannel
	if err := json.Unmarshal([]byte(`["UNKNOWN","STDOUT","STDERR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpWaypointEntrypointExecRequestOutputChannelEnum = append(hashicorpWaypointEntrypointExecRequestOutputChannelEnum, v)
	}
}

func (m HashicorpWaypointEntrypointExecRequestOutputChannel) validateHashicorpWaypointEntrypointExecRequestOutputChannelEnum(path, location string, value HashicorpWaypointEntrypointExecRequestOutputChannel) error {
	if err := validate.EnumCase(path, location, value, hashicorpWaypointEntrypointExecRequestOutputChannelEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp waypoint entrypoint exec request output channel
func (m HashicorpWaypointEntrypointExecRequestOutputChannel) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpWaypointEntrypointExecRequestOutputChannelEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
