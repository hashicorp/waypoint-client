// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointGetJobStreamResponseState hashicorp waypoint get job stream response state
//
// swagger:model hashicorp.waypoint.GetJobStreamResponse.State
type HashicorpWaypointGetJobStreamResponseState struct {

	// canceling is true if the job was requested to be canceled.
	Canceling bool `json:"canceling,omitempty"`

	// current
	Current HashicorpWaypointJobState `json:"current,omitempty"`

	// The full updated job is also sent because additional fields may be
	// set depending on the state (such as the assigned runner, assignment
	// times, etc.)
	Job *HashicorpWaypointJob `json:"job,omitempty"`

	// previous and current are the previous and current states, respectively.
	Previous HashicorpWaypointJobState `json:"previous,omitempty"`
}

// Validate validates this hashicorp waypoint get job stream response state
func (m *HashicorpWaypointGetJobStreamResponseState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJob(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointGetJobStreamResponseState) validateCurrent(formats strfmt.Registry) error {

	if swag.IsZero(m.Current) { // not required
		return nil
	}

	if err := m.Current.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("current")
		}
		return err
	}

	return nil
}

func (m *HashicorpWaypointGetJobStreamResponseState) validateJob(formats strfmt.Registry) error {

	if swag.IsZero(m.Job) { // not required
		return nil
	}

	if m.Job != nil {
		if err := m.Job.Validate3(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("job")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointGetJobStreamResponseState) validatePrevious(formats strfmt.Registry) error {

	if swag.IsZero(m.Previous) { // not required
		return nil
	}

	if err := m.Previous.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("previous")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointGetJobStreamResponseState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointGetJobStreamResponseState) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointGetJobStreamResponseState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
