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

// NewWaypointGetAuthMethodParams creates a new WaypointGetAuthMethodParams object
// with the default values initialized.
func NewWaypointGetAuthMethodParams() *WaypointGetAuthMethodParams {
	var ()
	return &WaypointGetAuthMethodParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointGetAuthMethodParamsWithTimeout creates a new WaypointGetAuthMethodParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointGetAuthMethodParamsWithTimeout(timeout time.Duration) *WaypointGetAuthMethodParams {
	var ()
	return &WaypointGetAuthMethodParams{

		timeout: timeout,
	}
}

// NewWaypointGetAuthMethodParamsWithContext creates a new WaypointGetAuthMethodParams object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointGetAuthMethodParamsWithContext(ctx context.Context) *WaypointGetAuthMethodParams {
	var ()
	return &WaypointGetAuthMethodParams{

		Context: ctx,
	}
}

// NewWaypointGetAuthMethodParamsWithHTTPClient creates a new WaypointGetAuthMethodParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointGetAuthMethodParamsWithHTTPClient(client *http.Client) *WaypointGetAuthMethodParams {
	var ()
	return &WaypointGetAuthMethodParams{
		HTTPClient: client,
	}
}

/*WaypointGetAuthMethodParams contains all the parameters to send to the API endpoint
for the waypoint get auth method operation typically these are written to a http.Request
*/
type WaypointGetAuthMethodParams struct {

	/*AuthMethodName*/
	AuthMethodName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint get auth method params
func (o *WaypointGetAuthMethodParams) WithTimeout(timeout time.Duration) *WaypointGetAuthMethodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint get auth method params
func (o *WaypointGetAuthMethodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint get auth method params
func (o *WaypointGetAuthMethodParams) WithContext(ctx context.Context) *WaypointGetAuthMethodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint get auth method params
func (o *WaypointGetAuthMethodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint get auth method params
func (o *WaypointGetAuthMethodParams) WithHTTPClient(client *http.Client) *WaypointGetAuthMethodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint get auth method params
func (o *WaypointGetAuthMethodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthMethodName adds the authMethodName to the waypoint get auth method params
func (o *WaypointGetAuthMethodParams) WithAuthMethodName(authMethodName string) *WaypointGetAuthMethodParams {
	o.SetAuthMethodName(authMethodName)
	return o
}

// SetAuthMethodName adds the authMethodName to the waypoint get auth method params
func (o *WaypointGetAuthMethodParams) SetAuthMethodName(authMethodName string) {
	o.AuthMethodName = authMethodName
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointGetAuthMethodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param auth_method.name
	if err := r.SetPathParam("auth_method.name", o.AuthMethodName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}