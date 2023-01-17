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

// NewWaypointDestroyProjectParams creates a new WaypointDestroyProjectParams object
// with the default values initialized.
func NewWaypointDestroyProjectParams() *WaypointDestroyProjectParams {
	var ()
	return &WaypointDestroyProjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointDestroyProjectParamsWithTimeout creates a new WaypointDestroyProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointDestroyProjectParamsWithTimeout(timeout time.Duration) *WaypointDestroyProjectParams {
	var ()
	return &WaypointDestroyProjectParams{

		timeout: timeout,
	}
}

// NewWaypointDestroyProjectParamsWithContext creates a new WaypointDestroyProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointDestroyProjectParamsWithContext(ctx context.Context) *WaypointDestroyProjectParams {
	var ()
	return &WaypointDestroyProjectParams{

		Context: ctx,
	}
}

// NewWaypointDestroyProjectParamsWithHTTPClient creates a new WaypointDestroyProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointDestroyProjectParamsWithHTTPClient(client *http.Client) *WaypointDestroyProjectParams {
	var ()
	return &WaypointDestroyProjectParams{
		HTTPClient: client,
	}
}

/*WaypointDestroyProjectParams contains all the parameters to send to the API endpoint
for the waypoint destroy project operation typically these are written to a http.Request
*/
type WaypointDestroyProjectParams struct {

	/*NamespaceID
	  namespace_id is the id of the namespace.

	*/
	NamespaceID string
	/*ProjectProject*/
	ProjectProject string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint destroy project params
func (o *WaypointDestroyProjectParams) WithTimeout(timeout time.Duration) *WaypointDestroyProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint destroy project params
func (o *WaypointDestroyProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint destroy project params
func (o *WaypointDestroyProjectParams) WithContext(ctx context.Context) *WaypointDestroyProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint destroy project params
func (o *WaypointDestroyProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint destroy project params
func (o *WaypointDestroyProjectParams) WithHTTPClient(client *http.Client) *WaypointDestroyProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint destroy project params
func (o *WaypointDestroyProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespaceID adds the namespaceID to the waypoint destroy project params
func (o *WaypointDestroyProjectParams) WithNamespaceID(namespaceID string) *WaypointDestroyProjectParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint destroy project params
func (o *WaypointDestroyProjectParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithProjectProject adds the projectProject to the waypoint destroy project params
func (o *WaypointDestroyProjectParams) WithProjectProject(projectProject string) *WaypointDestroyProjectParams {
	o.SetProjectProject(projectProject)
	return o
}

// SetProjectProject adds the projectProject to the waypoint destroy project params
func (o *WaypointDestroyProjectParams) SetProjectProject(projectProject string) {
	o.ProjectProject = projectProject
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointDestroyProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

	// path param project.project
	if err := r.SetPathParam("project.project", o.ProjectProject); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
