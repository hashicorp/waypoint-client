// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointUpsertTriggerRequest hashicorp waypoint upsert trigger request
//
// swagger:model hashicorp.waypoint.UpsertTriggerRequest
type HashicorpWaypointUpsertTriggerRequest struct {

	// Trigger URL to upsert. If the id in the message is empty, then this
	// will be an insert. Otherwise, this will be an update. If the ID
	// isn't found, it will be an error.
	Trigger *HashicorpWaypointTrigger `json:"trigger,omitempty"`
}

// Validate validates this hashicorp waypoint upsert trigger request
func (m *HashicorpWaypointUpsertTriggerRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTrigger(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointUpsertTriggerRequest) validateTrigger(formats strfmt.Registry) error {

	if swag.IsZero(m.Trigger) { // not required
		return nil
	}

	if m.Trigger != nil {
		if err := m.Trigger.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trigger")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointUpsertTriggerRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointUpsertTriggerRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointUpsertTriggerRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
