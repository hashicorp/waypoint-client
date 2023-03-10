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

// HashicorpWaypointReleaseLoadDetails hashicorp waypoint release load details
//
// swagger:model hashicorp.waypoint.Release.LoadDetails
type HashicorpWaypointReleaseLoadDetails string

const (

	// HashicorpWaypointReleaseLoadDetailsNONE captures enum value "NONE"
	HashicorpWaypointReleaseLoadDetailsNONE HashicorpWaypointReleaseLoadDetails = "NONE"

	// HashicorpWaypointReleaseLoadDetailsDEPLOYMENT captures enum value "DEPLOYMENT"
	HashicorpWaypointReleaseLoadDetailsDEPLOYMENT HashicorpWaypointReleaseLoadDetails = "DEPLOYMENT"

	// HashicorpWaypointReleaseLoadDetailsARTIFACT captures enum value "ARTIFACT"
	HashicorpWaypointReleaseLoadDetailsARTIFACT HashicorpWaypointReleaseLoadDetails = "ARTIFACT"

	// HashicorpWaypointReleaseLoadDetailsBUILD captures enum value "BUILD"
	HashicorpWaypointReleaseLoadDetailsBUILD HashicorpWaypointReleaseLoadDetails = "BUILD"
)

// for schema
var hashicorpWaypointReleaseLoadDetailsEnum []interface{}

func init() {
	var res []HashicorpWaypointReleaseLoadDetails
	if err := json.Unmarshal([]byte(`["NONE","DEPLOYMENT","ARTIFACT","BUILD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpWaypointReleaseLoadDetailsEnum = append(hashicorpWaypointReleaseLoadDetailsEnum, v)
	}
}

func (m HashicorpWaypointReleaseLoadDetails) validateHashicorpWaypointReleaseLoadDetailsEnum(path, location string, value HashicorpWaypointReleaseLoadDetails) error {
	if err := validate.EnumCase(path, location, value, hashicorpWaypointReleaseLoadDetailsEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp waypoint release load details
func (m HashicorpWaypointReleaseLoadDetails) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpWaypointReleaseLoadDetailsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
