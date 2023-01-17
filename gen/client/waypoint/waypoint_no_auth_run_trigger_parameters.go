// Code generated by go-swagger; DO NOT EDIT.

package waypoint

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

	"github.com/hashicorp/waypoint-client/gen/models"
)

// NewWaypointNoAuthRunTriggerParams creates a new WaypointNoAuthRunTriggerParams object
// with the default values initialized.
func NewWaypointNoAuthRunTriggerParams() *WaypointNoAuthRunTriggerParams {
	var ()
	return &WaypointNoAuthRunTriggerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointNoAuthRunTriggerParamsWithTimeout creates a new WaypointNoAuthRunTriggerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointNoAuthRunTriggerParamsWithTimeout(timeout time.Duration) *WaypointNoAuthRunTriggerParams {
	var ()
	return &WaypointNoAuthRunTriggerParams{

		timeout: timeout,
	}
}

// NewWaypointNoAuthRunTriggerParamsWithContext creates a new WaypointNoAuthRunTriggerParams object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointNoAuthRunTriggerParamsWithContext(ctx context.Context) *WaypointNoAuthRunTriggerParams {
	var ()
	return &WaypointNoAuthRunTriggerParams{

		Context: ctx,
	}
}

// NewWaypointNoAuthRunTriggerParamsWithHTTPClient creates a new WaypointNoAuthRunTriggerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointNoAuthRunTriggerParamsWithHTTPClient(client *http.Client) *WaypointNoAuthRunTriggerParams {
	var ()
	return &WaypointNoAuthRunTriggerParams{
		HTTPClient: client,
	}
}

/*WaypointNoAuthRunTriggerParams contains all the parameters to send to the API endpoint
for the waypoint no auth run trigger operation typically these are written to a http.Request
*/
type WaypointNoAuthRunTriggerParams struct {

	/*Body*/
	Body *models.HashicorpWaypointRunTriggerRequest
	/*NamespaceID
	  namespace_id is the id of the namespace.

	*/
	NamespaceID string
	/*RefID*/
	RefID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint no auth run trigger params
func (o *WaypointNoAuthRunTriggerParams) WithTimeout(timeout time.Duration) *WaypointNoAuthRunTriggerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint no auth run trigger params
func (o *WaypointNoAuthRunTriggerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint no auth run trigger params
func (o *WaypointNoAuthRunTriggerParams) WithContext(ctx context.Context) *WaypointNoAuthRunTriggerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint no auth run trigger params
func (o *WaypointNoAuthRunTriggerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint no auth run trigger params
func (o *WaypointNoAuthRunTriggerParams) WithHTTPClient(client *http.Client) *WaypointNoAuthRunTriggerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint no auth run trigger params
func (o *WaypointNoAuthRunTriggerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint no auth run trigger params
func (o *WaypointNoAuthRunTriggerParams) WithBody(body *models.HashicorpWaypointRunTriggerRequest) *WaypointNoAuthRunTriggerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint no auth run trigger params
func (o *WaypointNoAuthRunTriggerParams) SetBody(body *models.HashicorpWaypointRunTriggerRequest) {
	o.Body = body
}

// WithNamespaceID adds the namespaceID to the waypoint no auth run trigger params
func (o *WaypointNoAuthRunTriggerParams) WithNamespaceID(namespaceID string) *WaypointNoAuthRunTriggerParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint no auth run trigger params
func (o *WaypointNoAuthRunTriggerParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithRefID adds the refID to the waypoint no auth run trigger params
func (o *WaypointNoAuthRunTriggerParams) WithRefID(refID string) *WaypointNoAuthRunTriggerParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the waypoint no auth run trigger params
func (o *WaypointNoAuthRunTriggerParams) SetRefID(refID string) {
	o.RefID = refID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointNoAuthRunTriggerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

	// path param ref.id
	if err := r.SetPathParam("ref.id", o.RefID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
