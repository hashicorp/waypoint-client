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

// HashicorpCloudWaypointNamespaceBillingDeprovisionReason hashicorp cloud waypoint namespace billing deprovision reason
//
// swagger:model hashicorp.cloud.waypoint.NamespaceBillingDeprovision.Reason
type HashicorpCloudWaypointNamespaceBillingDeprovisionReason string

const (

	// HashicorpCloudWaypointNamespaceBillingDeprovisionReasonDELINQUENTBILLINGACCOUNT captures enum value "DELINQUENT_BILLING_ACCOUNT"
	HashicorpCloudWaypointNamespaceBillingDeprovisionReasonDELINQUENTBILLINGACCOUNT HashicorpCloudWaypointNamespaceBillingDeprovisionReason = "DELINQUENT_BILLING_ACCOUNT"

	// HashicorpCloudWaypointNamespaceBillingDeprovisionReasonUSERREQUEST captures enum value "USER_REQUEST"
	HashicorpCloudWaypointNamespaceBillingDeprovisionReasonUSERREQUEST HashicorpCloudWaypointNamespaceBillingDeprovisionReason = "USER_REQUEST"
)

// for schema
var hashicorpCloudWaypointNamespaceBillingDeprovisionReasonEnum []interface{}

func init() {
	var res []HashicorpCloudWaypointNamespaceBillingDeprovisionReason
	if err := json.Unmarshal([]byte(`["DELINQUENT_BILLING_ACCOUNT","USER_REQUEST"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudWaypointNamespaceBillingDeprovisionReasonEnum = append(hashicorpCloudWaypointNamespaceBillingDeprovisionReasonEnum, v)
	}
}

func (m HashicorpCloudWaypointNamespaceBillingDeprovisionReason) validateHashicorpCloudWaypointNamespaceBillingDeprovisionReasonEnum(path, location string, value HashicorpCloudWaypointNamespaceBillingDeprovisionReason) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudWaypointNamespaceBillingDeprovisionReasonEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud waypoint namespace billing deprovision reason
func (m HashicorpCloudWaypointNamespaceBillingDeprovisionReason) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudWaypointNamespaceBillingDeprovisionReasonEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
