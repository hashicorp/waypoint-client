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

// NewWaypointListUsersParams creates a new WaypointListUsersParams object
// with the default values initialized.
func NewWaypointListUsersParams() *WaypointListUsersParams {

	return &WaypointListUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointListUsersParamsWithTimeout creates a new WaypointListUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointListUsersParamsWithTimeout(timeout time.Duration) *WaypointListUsersParams {

	return &WaypointListUsersParams{

		timeout: timeout,
	}
}

// NewWaypointListUsersParamsWithContext creates a new WaypointListUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointListUsersParamsWithContext(ctx context.Context) *WaypointListUsersParams {

	return &WaypointListUsersParams{

		Context: ctx,
	}
}

// NewWaypointListUsersParamsWithHTTPClient creates a new WaypointListUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointListUsersParamsWithHTTPClient(client *http.Client) *WaypointListUsersParams {

	return &WaypointListUsersParams{
		HTTPClient: client,
	}
}

/*WaypointListUsersParams contains all the parameters to send to the API endpoint
for the waypoint list users operation typically these are written to a http.Request
*/
type WaypointListUsersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint list users params
func (o *WaypointListUsersParams) WithTimeout(timeout time.Duration) *WaypointListUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint list users params
func (o *WaypointListUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint list users params
func (o *WaypointListUsersParams) WithContext(ctx context.Context) *WaypointListUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint list users params
func (o *WaypointListUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint list users params
func (o *WaypointListUsersParams) WithHTTPClient(client *http.Client) *WaypointListUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint list users params
func (o *WaypointListUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointListUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
