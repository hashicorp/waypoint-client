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

// NewWaypointGetLatestBuildParams creates a new WaypointGetLatestBuildParams object
// with the default values initialized.
func NewWaypointGetLatestBuildParams() *WaypointGetLatestBuildParams {
	var ()
	return &WaypointGetLatestBuildParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointGetLatestBuildParamsWithTimeout creates a new WaypointGetLatestBuildParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointGetLatestBuildParamsWithTimeout(timeout time.Duration) *WaypointGetLatestBuildParams {
	var ()
	return &WaypointGetLatestBuildParams{

		timeout: timeout,
	}
}

// NewWaypointGetLatestBuildParamsWithContext creates a new WaypointGetLatestBuildParams object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointGetLatestBuildParamsWithContext(ctx context.Context) *WaypointGetLatestBuildParams {
	var ()
	return &WaypointGetLatestBuildParams{

		Context: ctx,
	}
}

// NewWaypointGetLatestBuildParamsWithHTTPClient creates a new WaypointGetLatestBuildParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointGetLatestBuildParamsWithHTTPClient(client *http.Client) *WaypointGetLatestBuildParams {
	var ()
	return &WaypointGetLatestBuildParams{
		HTTPClient: client,
	}
}

/*WaypointGetLatestBuildParams contains all the parameters to send to the API endpoint
for the waypoint get latest build operation typically these are written to a http.Request
*/
type WaypointGetLatestBuildParams struct {

	/*ApplicationApplication*/
	ApplicationApplication string
	/*ApplicationProject*/
	ApplicationProject string
	/*NamespaceID
	  namespace_id is the id of the namespace.

	*/
	NamespaceID string
	/*WorkspaceWorkspace*/
	WorkspaceWorkspace *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint get latest build params
func (o *WaypointGetLatestBuildParams) WithTimeout(timeout time.Duration) *WaypointGetLatestBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint get latest build params
func (o *WaypointGetLatestBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint get latest build params
func (o *WaypointGetLatestBuildParams) WithContext(ctx context.Context) *WaypointGetLatestBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint get latest build params
func (o *WaypointGetLatestBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint get latest build params
func (o *WaypointGetLatestBuildParams) WithHTTPClient(client *http.Client) *WaypointGetLatestBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint get latest build params
func (o *WaypointGetLatestBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationApplication adds the applicationApplication to the waypoint get latest build params
func (o *WaypointGetLatestBuildParams) WithApplicationApplication(applicationApplication string) *WaypointGetLatestBuildParams {
	o.SetApplicationApplication(applicationApplication)
	return o
}

// SetApplicationApplication adds the applicationApplication to the waypoint get latest build params
func (o *WaypointGetLatestBuildParams) SetApplicationApplication(applicationApplication string) {
	o.ApplicationApplication = applicationApplication
}

// WithApplicationProject adds the applicationProject to the waypoint get latest build params
func (o *WaypointGetLatestBuildParams) WithApplicationProject(applicationProject string) *WaypointGetLatestBuildParams {
	o.SetApplicationProject(applicationProject)
	return o
}

// SetApplicationProject adds the applicationProject to the waypoint get latest build params
func (o *WaypointGetLatestBuildParams) SetApplicationProject(applicationProject string) {
	o.ApplicationProject = applicationProject
}

// WithNamespaceID adds the namespaceID to the waypoint get latest build params
func (o *WaypointGetLatestBuildParams) WithNamespaceID(namespaceID string) *WaypointGetLatestBuildParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint get latest build params
func (o *WaypointGetLatestBuildParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithWorkspaceWorkspace adds the workspaceWorkspace to the waypoint get latest build params
func (o *WaypointGetLatestBuildParams) WithWorkspaceWorkspace(workspaceWorkspace *string) *WaypointGetLatestBuildParams {
	o.SetWorkspaceWorkspace(workspaceWorkspace)
	return o
}

// SetWorkspaceWorkspace adds the workspaceWorkspace to the waypoint get latest build params
func (o *WaypointGetLatestBuildParams) SetWorkspaceWorkspace(workspaceWorkspace *string) {
	o.WorkspaceWorkspace = workspaceWorkspace
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointGetLatestBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param application.application
	if err := r.SetPathParam("application.application", o.ApplicationApplication); err != nil {
		return err
	}

	// path param application.project
	if err := r.SetPathParam("application.project", o.ApplicationProject); err != nil {
		return err
	}

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

	if o.WorkspaceWorkspace != nil {

		// query param workspace.workspace
		var qrWorkspaceWorkspace string
		if o.WorkspaceWorkspace != nil {
			qrWorkspaceWorkspace = *o.WorkspaceWorkspace
		}
		qWorkspaceWorkspace := qrWorkspaceWorkspace
		if qWorkspaceWorkspace != "" {
			if err := r.SetQueryParam("workspace.workspace", qWorkspaceWorkspace); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
