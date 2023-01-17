// Code generated by go-swagger; DO NOT EDIT.

package waypoint_control_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewWaypointControlServiceGetNamespaceParams creates a new WaypointControlServiceGetNamespaceParams object
// with the default values initialized.
func NewWaypointControlServiceGetNamespaceParams() *WaypointControlServiceGetNamespaceParams {
	var ()
	return &WaypointControlServiceGetNamespaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointControlServiceGetNamespaceParamsWithTimeout creates a new WaypointControlServiceGetNamespaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointControlServiceGetNamespaceParamsWithTimeout(timeout time.Duration) *WaypointControlServiceGetNamespaceParams {
	var ()
	return &WaypointControlServiceGetNamespaceParams{

		timeout: timeout,
	}
}

// NewWaypointControlServiceGetNamespaceParamsWithContext creates a new WaypointControlServiceGetNamespaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointControlServiceGetNamespaceParamsWithContext(ctx context.Context) *WaypointControlServiceGetNamespaceParams {
	var ()
	return &WaypointControlServiceGetNamespaceParams{

		Context: ctx,
	}
}

// NewWaypointControlServiceGetNamespaceParamsWithHTTPClient creates a new WaypointControlServiceGetNamespaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointControlServiceGetNamespaceParamsWithHTTPClient(client *http.Client) *WaypointControlServiceGetNamespaceParams {
	var ()
	return &WaypointControlServiceGetNamespaceParams{
		HTTPClient: client,
	}
}

/*WaypointControlServiceGetNamespaceParams contains all the parameters to send to the API endpoint
for the waypoint control service get namespace operation typically these are written to a http.Request
*/
type WaypointControlServiceGetNamespaceParams struct {

	/*CheckOnly*/
	CheckOnly *bool
	/*LocationOrganizationID
	  organization_id is the id of the organization.

	*/
	LocationOrganizationID string
	/*LocationProjectID
	  project_id is the projects id.

	*/
	LocationProjectID string
	/*LocationRegionProvider
	  provider is the named cloud provider ("aws", "gcp", "azure").

	*/
	LocationRegionProvider *string
	/*LocationRegionRegion
	  region is the cloud region ("us-west1", "us-east1").

	*/
	LocationRegionRegion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint control service get namespace params
func (o *WaypointControlServiceGetNamespaceParams) WithTimeout(timeout time.Duration) *WaypointControlServiceGetNamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint control service get namespace params
func (o *WaypointControlServiceGetNamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint control service get namespace params
func (o *WaypointControlServiceGetNamespaceParams) WithContext(ctx context.Context) *WaypointControlServiceGetNamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint control service get namespace params
func (o *WaypointControlServiceGetNamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint control service get namespace params
func (o *WaypointControlServiceGetNamespaceParams) WithHTTPClient(client *http.Client) *WaypointControlServiceGetNamespaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint control service get namespace params
func (o *WaypointControlServiceGetNamespaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCheckOnly adds the checkOnly to the waypoint control service get namespace params
func (o *WaypointControlServiceGetNamespaceParams) WithCheckOnly(checkOnly *bool) *WaypointControlServiceGetNamespaceParams {
	o.SetCheckOnly(checkOnly)
	return o
}

// SetCheckOnly adds the checkOnly to the waypoint control service get namespace params
func (o *WaypointControlServiceGetNamespaceParams) SetCheckOnly(checkOnly *bool) {
	o.CheckOnly = checkOnly
}

// WithLocationOrganizationID adds the locationOrganizationID to the waypoint control service get namespace params
func (o *WaypointControlServiceGetNamespaceParams) WithLocationOrganizationID(locationOrganizationID string) *WaypointControlServiceGetNamespaceParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the waypoint control service get namespace params
func (o *WaypointControlServiceGetNamespaceParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the waypoint control service get namespace params
func (o *WaypointControlServiceGetNamespaceParams) WithLocationProjectID(locationProjectID string) *WaypointControlServiceGetNamespaceParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the waypoint control service get namespace params
func (o *WaypointControlServiceGetNamespaceParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the waypoint control service get namespace params
func (o *WaypointControlServiceGetNamespaceParams) WithLocationRegionProvider(locationRegionProvider *string) *WaypointControlServiceGetNamespaceParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the waypoint control service get namespace params
func (o *WaypointControlServiceGetNamespaceParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the waypoint control service get namespace params
func (o *WaypointControlServiceGetNamespaceParams) WithLocationRegionRegion(locationRegionRegion *string) *WaypointControlServiceGetNamespaceParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the waypoint control service get namespace params
func (o *WaypointControlServiceGetNamespaceParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointControlServiceGetNamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CheckOnly != nil {

		// query param check_only
		var qrCheckOnly bool
		if o.CheckOnly != nil {
			qrCheckOnly = *o.CheckOnly
		}
		qCheckOnly := swag.FormatBool(qrCheckOnly)
		if qCheckOnly != "" {
			if err := r.SetQueryParam("check_only", qCheckOnly); err != nil {
				return err
			}
		}

	}

	// path param location.organization_id
	if err := r.SetPathParam("location.organization_id", o.LocationOrganizationID); err != nil {
		return err
	}

	// path param location.project_id
	if err := r.SetPathParam("location.project_id", o.LocationProjectID); err != nil {
		return err
	}

	if o.LocationRegionProvider != nil {

		// query param location.region.provider
		var qrLocationRegionProvider string
		if o.LocationRegionProvider != nil {
			qrLocationRegionProvider = *o.LocationRegionProvider
		}
		qLocationRegionProvider := qrLocationRegionProvider
		if qLocationRegionProvider != "" {
			if err := r.SetQueryParam("location.region.provider", qLocationRegionProvider); err != nil {
				return err
			}
		}

	}

	if o.LocationRegionRegion != nil {

		// query param location.region.region
		var qrLocationRegionRegion string
		if o.LocationRegionRegion != nil {
			qrLocationRegionRegion = *o.LocationRegionRegion
		}
		qLocationRegionRegion := qrLocationRegionRegion
		if qLocationRegionRegion != "" {
			if err := r.SetQueryParam("location.region.region", qLocationRegionRegion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
