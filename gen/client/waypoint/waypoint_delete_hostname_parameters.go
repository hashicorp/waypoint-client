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

// NewWaypointDeleteHostnameParams creates a new WaypointDeleteHostnameParams object
// with the default values initialized.
func NewWaypointDeleteHostnameParams() *WaypointDeleteHostnameParams {
	var ()
	return &WaypointDeleteHostnameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointDeleteHostnameParamsWithTimeout creates a new WaypointDeleteHostnameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointDeleteHostnameParamsWithTimeout(timeout time.Duration) *WaypointDeleteHostnameParams {
	var ()
	return &WaypointDeleteHostnameParams{

		timeout: timeout,
	}
}

// NewWaypointDeleteHostnameParamsWithContext creates a new WaypointDeleteHostnameParams object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointDeleteHostnameParamsWithContext(ctx context.Context) *WaypointDeleteHostnameParams {
	var ()
	return &WaypointDeleteHostnameParams{

		Context: ctx,
	}
}

// NewWaypointDeleteHostnameParamsWithHTTPClient creates a new WaypointDeleteHostnameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointDeleteHostnameParamsWithHTTPClient(client *http.Client) *WaypointDeleteHostnameParams {
	var ()
	return &WaypointDeleteHostnameParams{
		HTTPClient: client,
	}
}

/*WaypointDeleteHostnameParams contains all the parameters to send to the API endpoint
for the waypoint delete hostname operation typically these are written to a http.Request
*/
type WaypointDeleteHostnameParams struct {

	/*Hostname*/
	Hostname string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint delete hostname params
func (o *WaypointDeleteHostnameParams) WithTimeout(timeout time.Duration) *WaypointDeleteHostnameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint delete hostname params
func (o *WaypointDeleteHostnameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint delete hostname params
func (o *WaypointDeleteHostnameParams) WithContext(ctx context.Context) *WaypointDeleteHostnameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint delete hostname params
func (o *WaypointDeleteHostnameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint delete hostname params
func (o *WaypointDeleteHostnameParams) WithHTTPClient(client *http.Client) *WaypointDeleteHostnameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint delete hostname params
func (o *WaypointDeleteHostnameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHostname adds the hostname to the waypoint delete hostname params
func (o *WaypointDeleteHostnameParams) WithHostname(hostname string) *WaypointDeleteHostnameParams {
	o.SetHostname(hostname)
	return o
}

// SetHostname adds the hostname to the waypoint delete hostname params
func (o *WaypointDeleteHostnameParams) SetHostname(hostname string) {
	o.Hostname = hostname
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointDeleteHostnameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param hostname
	if err := r.SetPathParam("hostname", o.Hostname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
