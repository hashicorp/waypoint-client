// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointJobDocsResultResult hashicorp waypoint job docs result result
//
// swagger:model hashicorp.waypoint.Job.DocsResult.Result
type HashicorpWaypointJobDocsResultResult struct {

	// component that the docs are for
	Component *HashicorpWaypointComponent `json:"component,omitempty"`

	// docs
	Docs *HashicorpWaypointDocumentation `json:"docs,omitempty"`
}

// Validate validates this hashicorp waypoint job docs result result
func (m *HashicorpWaypointJobDocsResultResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDocs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointJobDocsResultResult) validateComponent(formats strfmt.Registry) error {

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

func (m *HashicorpWaypointJobDocsResultResult) validateDocs(formats strfmt.Registry) error {

	if swag.IsZero(m.Docs) { // not required
		return nil
	}

	if m.Docs != nil {
		if err := m.Docs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("docs")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointJobDocsResultResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointJobDocsResultResult) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointJobDocsResultResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
