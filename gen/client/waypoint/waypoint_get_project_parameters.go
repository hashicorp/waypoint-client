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

// NewWaypointGetProjectParams creates a new WaypointGetProjectParams object
// with the default values initialized.
func NewWaypointGetProjectParams() *WaypointGetProjectParams {
	var ()
	return &WaypointGetProjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointGetProjectParamsWithTimeout creates a new WaypointGetProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointGetProjectParamsWithTimeout(timeout time.Duration) *WaypointGetProjectParams {
	var ()
	return &WaypointGetProjectParams{

		timeout: timeout,
	}
}

// NewWaypointGetProjectParamsWithContext creates a new WaypointGetProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointGetProjectParamsWithContext(ctx context.Context) *WaypointGetProjectParams {
	var ()
	return &WaypointGetProjectParams{

		Context: ctx,
	}
}

// NewWaypointGetProjectParamsWithHTTPClient creates a new WaypointGetProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointGetProjectParamsWithHTTPClient(client *http.Client) *WaypointGetProjectParams {
	var ()
	return &WaypointGetProjectParams{
		HTTPClient: client,
	}
}

/*WaypointGetProjectParams contains all the parameters to send to the API endpoint
for the waypoint get project operation typically these are written to a http.Request
*/
type WaypointGetProjectParams struct {

	/*ProjectProject*/
	ProjectProject string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint get project params
func (o *WaypointGetProjectParams) WithTimeout(timeout time.Duration) *WaypointGetProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint get project params
func (o *WaypointGetProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint get project params
func (o *WaypointGetProjectParams) WithContext(ctx context.Context) *WaypointGetProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint get project params
func (o *WaypointGetProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint get project params
func (o *WaypointGetProjectParams) WithHTTPClient(client *http.Client) *WaypointGetProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint get project params
func (o *WaypointGetProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectProject adds the projectProject to the waypoint get project params
func (o *WaypointGetProjectParams) WithProjectProject(projectProject string) *WaypointGetProjectParams {
	o.SetProjectProject(projectProject)
	return o
}

// SetProjectProject adds the projectProject to the waypoint get project params
func (o *WaypointGetProjectParams) SetProjectProject(projectProject string) {
	o.ProjectProject = projectProject
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointGetProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project.project
	if err := r.SetPathParam("project.project", o.ProjectProject); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
