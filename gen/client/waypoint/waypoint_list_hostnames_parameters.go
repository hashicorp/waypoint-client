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

// NewWaypointListHostnamesParams creates a new WaypointListHostnamesParams object
// with the default values initialized.
func NewWaypointListHostnamesParams() *WaypointListHostnamesParams {
	var ()
	return &WaypointListHostnamesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointListHostnamesParamsWithTimeout creates a new WaypointListHostnamesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointListHostnamesParamsWithTimeout(timeout time.Duration) *WaypointListHostnamesParams {
	var ()
	return &WaypointListHostnamesParams{

		timeout: timeout,
	}
}

// NewWaypointListHostnamesParamsWithContext creates a new WaypointListHostnamesParams object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointListHostnamesParamsWithContext(ctx context.Context) *WaypointListHostnamesParams {
	var ()
	return &WaypointListHostnamesParams{

		Context: ctx,
	}
}

// NewWaypointListHostnamesParamsWithHTTPClient creates a new WaypointListHostnamesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointListHostnamesParamsWithHTTPClient(client *http.Client) *WaypointListHostnamesParams {
	var ()
	return &WaypointListHostnamesParams{
		HTTPClient: client,
	}
}

/*WaypointListHostnamesParams contains all the parameters to send to the API endpoint
for the waypoint list hostnames operation typically these are written to a http.Request
*/
type WaypointListHostnamesParams struct {

	/*TargetApplicationApplicationApplication*/
	TargetApplicationApplicationApplication string
	/*TargetApplicationApplicationProject*/
	TargetApplicationApplicationProject string
	/*TargetApplicationWorkspaceWorkspace*/
	TargetApplicationWorkspaceWorkspace *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint list hostnames params
func (o *WaypointListHostnamesParams) WithTimeout(timeout time.Duration) *WaypointListHostnamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint list hostnames params
func (o *WaypointListHostnamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint list hostnames params
func (o *WaypointListHostnamesParams) WithContext(ctx context.Context) *WaypointListHostnamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint list hostnames params
func (o *WaypointListHostnamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint list hostnames params
func (o *WaypointListHostnamesParams) WithHTTPClient(client *http.Client) *WaypointListHostnamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint list hostnames params
func (o *WaypointListHostnamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTargetApplicationApplicationApplication adds the targetApplicationApplicationApplication to the waypoint list hostnames params
func (o *WaypointListHostnamesParams) WithTargetApplicationApplicationApplication(targetApplicationApplicationApplication string) *WaypointListHostnamesParams {
	o.SetTargetApplicationApplicationApplication(targetApplicationApplicationApplication)
	return o
}

// SetTargetApplicationApplicationApplication adds the targetApplicationApplicationApplication to the waypoint list hostnames params
func (o *WaypointListHostnamesParams) SetTargetApplicationApplicationApplication(targetApplicationApplicationApplication string) {
	o.TargetApplicationApplicationApplication = targetApplicationApplicationApplication
}

// WithTargetApplicationApplicationProject adds the targetApplicationApplicationProject to the waypoint list hostnames params
func (o *WaypointListHostnamesParams) WithTargetApplicationApplicationProject(targetApplicationApplicationProject string) *WaypointListHostnamesParams {
	o.SetTargetApplicationApplicationProject(targetApplicationApplicationProject)
	return o
}

// SetTargetApplicationApplicationProject adds the targetApplicationApplicationProject to the waypoint list hostnames params
func (o *WaypointListHostnamesParams) SetTargetApplicationApplicationProject(targetApplicationApplicationProject string) {
	o.TargetApplicationApplicationProject = targetApplicationApplicationProject
}

// WithTargetApplicationWorkspaceWorkspace adds the targetApplicationWorkspaceWorkspace to the waypoint list hostnames params
func (o *WaypointListHostnamesParams) WithTargetApplicationWorkspaceWorkspace(targetApplicationWorkspaceWorkspace *string) *WaypointListHostnamesParams {
	o.SetTargetApplicationWorkspaceWorkspace(targetApplicationWorkspaceWorkspace)
	return o
}

// SetTargetApplicationWorkspaceWorkspace adds the targetApplicationWorkspaceWorkspace to the waypoint list hostnames params
func (o *WaypointListHostnamesParams) SetTargetApplicationWorkspaceWorkspace(targetApplicationWorkspaceWorkspace *string) {
	o.TargetApplicationWorkspaceWorkspace = targetApplicationWorkspaceWorkspace
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointListHostnamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param target.application.application.application
	if err := r.SetPathParam("target.application.application.application", o.TargetApplicationApplicationApplication); err != nil {
		return err
	}

	// path param target.application.application.project
	if err := r.SetPathParam("target.application.application.project", o.TargetApplicationApplicationProject); err != nil {
		return err
	}

	if o.TargetApplicationWorkspaceWorkspace != nil {

		// query param target.application.workspace.workspace
		var qrTargetApplicationWorkspaceWorkspace string
		if o.TargetApplicationWorkspaceWorkspace != nil {
			qrTargetApplicationWorkspaceWorkspace = *o.TargetApplicationWorkspaceWorkspace
		}
		qTargetApplicationWorkspaceWorkspace := qrTargetApplicationWorkspaceWorkspace
		if qTargetApplicationWorkspaceWorkspace != "" {
			if err := r.SetQueryParam("target.application.workspace.workspace", qTargetApplicationWorkspaceWorkspace); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
