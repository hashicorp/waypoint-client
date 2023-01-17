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

// NewWaypointGetLatestBuild2Params creates a new WaypointGetLatestBuild2Params object
// with the default values initialized.
func NewWaypointGetLatestBuild2Params() *WaypointGetLatestBuild2Params {
	var ()
	return &WaypointGetLatestBuild2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointGetLatestBuild2ParamsWithTimeout creates a new WaypointGetLatestBuild2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointGetLatestBuild2ParamsWithTimeout(timeout time.Duration) *WaypointGetLatestBuild2Params {
	var ()
	return &WaypointGetLatestBuild2Params{

		timeout: timeout,
	}
}

// NewWaypointGetLatestBuild2ParamsWithContext creates a new WaypointGetLatestBuild2Params object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointGetLatestBuild2ParamsWithContext(ctx context.Context) *WaypointGetLatestBuild2Params {
	var ()
	return &WaypointGetLatestBuild2Params{

		Context: ctx,
	}
}

// NewWaypointGetLatestBuild2ParamsWithHTTPClient creates a new WaypointGetLatestBuild2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointGetLatestBuild2ParamsWithHTTPClient(client *http.Client) *WaypointGetLatestBuild2Params {
	var ()
	return &WaypointGetLatestBuild2Params{
		HTTPClient: client,
	}
}

/*WaypointGetLatestBuild2Params contains all the parameters to send to the API endpoint
for the waypoint get latest build2 operation typically these are written to a http.Request
*/
type WaypointGetLatestBuild2Params struct {

	/*ApplicationApplication*/
	ApplicationApplication string
	/*ApplicationProject*/
	ApplicationProject string
	/*NamespaceID
	  namespace_id is the id of the namespace.

	*/
	NamespaceID string
	/*WorkspaceWorkspace*/
	WorkspaceWorkspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint get latest build2 params
func (o *WaypointGetLatestBuild2Params) WithTimeout(timeout time.Duration) *WaypointGetLatestBuild2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint get latest build2 params
func (o *WaypointGetLatestBuild2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint get latest build2 params
func (o *WaypointGetLatestBuild2Params) WithContext(ctx context.Context) *WaypointGetLatestBuild2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint get latest build2 params
func (o *WaypointGetLatestBuild2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint get latest build2 params
func (o *WaypointGetLatestBuild2Params) WithHTTPClient(client *http.Client) *WaypointGetLatestBuild2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint get latest build2 params
func (o *WaypointGetLatestBuild2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationApplication adds the applicationApplication to the waypoint get latest build2 params
func (o *WaypointGetLatestBuild2Params) WithApplicationApplication(applicationApplication string) *WaypointGetLatestBuild2Params {
	o.SetApplicationApplication(applicationApplication)
	return o
}

// SetApplicationApplication adds the applicationApplication to the waypoint get latest build2 params
func (o *WaypointGetLatestBuild2Params) SetApplicationApplication(applicationApplication string) {
	o.ApplicationApplication = applicationApplication
}

// WithApplicationProject adds the applicationProject to the waypoint get latest build2 params
func (o *WaypointGetLatestBuild2Params) WithApplicationProject(applicationProject string) *WaypointGetLatestBuild2Params {
	o.SetApplicationProject(applicationProject)
	return o
}

// SetApplicationProject adds the applicationProject to the waypoint get latest build2 params
func (o *WaypointGetLatestBuild2Params) SetApplicationProject(applicationProject string) {
	o.ApplicationProject = applicationProject
}

// WithNamespaceID adds the namespaceID to the waypoint get latest build2 params
func (o *WaypointGetLatestBuild2Params) WithNamespaceID(namespaceID string) *WaypointGetLatestBuild2Params {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint get latest build2 params
func (o *WaypointGetLatestBuild2Params) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithWorkspaceWorkspace adds the workspaceWorkspace to the waypoint get latest build2 params
func (o *WaypointGetLatestBuild2Params) WithWorkspaceWorkspace(workspaceWorkspace string) *WaypointGetLatestBuild2Params {
	o.SetWorkspaceWorkspace(workspaceWorkspace)
	return o
}

// SetWorkspaceWorkspace adds the workspaceWorkspace to the waypoint get latest build2 params
func (o *WaypointGetLatestBuild2Params) SetWorkspaceWorkspace(workspaceWorkspace string) {
	o.WorkspaceWorkspace = workspaceWorkspace
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointGetLatestBuild2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param workspace.workspace
	if err := r.SetPathParam("workspace.workspace", o.WorkspaceWorkspace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
