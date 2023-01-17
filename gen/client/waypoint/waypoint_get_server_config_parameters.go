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
)

// NewWaypointGetServerConfigParams creates a new WaypointGetServerConfigParams object
// with the default values initialized.
func NewWaypointGetServerConfigParams() *WaypointGetServerConfigParams {
	var ()
	return &WaypointGetServerConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointGetServerConfigParamsWithTimeout creates a new WaypointGetServerConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointGetServerConfigParamsWithTimeout(timeout time.Duration) *WaypointGetServerConfigParams {
	var ()
	return &WaypointGetServerConfigParams{

		timeout: timeout,
	}
}

// NewWaypointGetServerConfigParamsWithContext creates a new WaypointGetServerConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointGetServerConfigParamsWithContext(ctx context.Context) *WaypointGetServerConfigParams {
	var ()
	return &WaypointGetServerConfigParams{

		Context: ctx,
	}
}

// NewWaypointGetServerConfigParamsWithHTTPClient creates a new WaypointGetServerConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointGetServerConfigParamsWithHTTPClient(client *http.Client) *WaypointGetServerConfigParams {
	var ()
	return &WaypointGetServerConfigParams{
		HTTPClient: client,
	}
}

/*WaypointGetServerConfigParams contains all the parameters to send to the API endpoint
for the waypoint get server config operation typically these are written to a http.Request
*/
type WaypointGetServerConfigParams struct {

	/*NamespaceID
	  namespace_id is the id of the namespace.

	*/
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint get server config params
func (o *WaypointGetServerConfigParams) WithTimeout(timeout time.Duration) *WaypointGetServerConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint get server config params
func (o *WaypointGetServerConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint get server config params
func (o *WaypointGetServerConfigParams) WithContext(ctx context.Context) *WaypointGetServerConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint get server config params
func (o *WaypointGetServerConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint get server config params
func (o *WaypointGetServerConfigParams) WithHTTPClient(client *http.Client) *WaypointGetServerConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint get server config params
func (o *WaypointGetServerConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespaceID adds the namespaceID to the waypoint get server config params
func (o *WaypointGetServerConfigParams) WithNamespaceID(namespaceID string) *WaypointGetServerConfigParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint get server config params
func (o *WaypointGetServerConfigParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointGetServerConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
