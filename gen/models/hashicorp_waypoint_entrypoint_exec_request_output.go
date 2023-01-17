// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointEntrypointExecRequestOutput hashicorp waypoint entrypoint exec request output
//
// swagger:model hashicorp.waypoint.EntrypointExecRequest.Output
type HashicorpWaypointEntrypointExecRequestOutput struct {

	// channel
	Channel HashicorpWaypointEntrypointExecRequestOutputChannel `json:"channel,omitempty"`

	// data
	// Format: byte
	Data strfmt.Base64 `json:"data,omitempty"`
}

// Validate validates this hashicorp waypoint entrypoint exec request output
func (m *HashicorpWaypointEntrypointExecRequestOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointEntrypointExecRequestOutput) validateChannel(formats strfmt.Registry) error {

	if swag.IsZero(m.Channel) { // not required
		return nil
	}

	if err := m.Channel.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("channel")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointEntrypointExecRequestOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointEntrypointExecRequestOutput) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointEntrypointExecRequestOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
