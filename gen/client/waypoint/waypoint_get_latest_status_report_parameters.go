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

// NewWaypointGetLatestStatusReportParams creates a new WaypointGetLatestStatusReportParams object
// with the default values initialized.
func NewWaypointGetLatestStatusReportParams() *WaypointGetLatestStatusReportParams {
	var ()
	return &WaypointGetLatestStatusReportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointGetLatestStatusReportParamsWithTimeout creates a new WaypointGetLatestStatusReportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointGetLatestStatusReportParamsWithTimeout(timeout time.Duration) *WaypointGetLatestStatusReportParams {
	var ()
	return &WaypointGetLatestStatusReportParams{

		timeout: timeout,
	}
}

// NewWaypointGetLatestStatusReportParamsWithContext creates a new WaypointGetLatestStatusReportParams object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointGetLatestStatusReportParamsWithContext(ctx context.Context) *WaypointGetLatestStatusReportParams {
	var ()
	return &WaypointGetLatestStatusReportParams{

		Context: ctx,
	}
}

// NewWaypointGetLatestStatusReportParamsWithHTTPClient creates a new WaypointGetLatestStatusReportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointGetLatestStatusReportParamsWithHTTPClient(client *http.Client) *WaypointGetLatestStatusReportParams {
	var ()
	return &WaypointGetLatestStatusReportParams{
		HTTPClient: client,
	}
}

/*WaypointGetLatestStatusReportParams contains all the parameters to send to the API endpoint
for the waypoint get latest status report operation typically these are written to a http.Request
*/
type WaypointGetLatestStatusReportParams struct {

	/*ApplicationApplication*/
	ApplicationApplication string
	/*ApplicationProject*/
	ApplicationProject string
	/*DeploymentID*/
	DeploymentID *string
	/*NamespaceID
	  namespace_id is the id of the namespace.

	*/
	NamespaceID string
	/*ReleaseID*/
	ReleaseID *string
	/*WorkspaceWorkspace*/
	WorkspaceWorkspace *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint get latest status report params
func (o *WaypointGetLatestStatusReportParams) WithTimeout(timeout time.Duration) *WaypointGetLatestStatusReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint get latest status report params
func (o *WaypointGetLatestStatusReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint get latest status report params
func (o *WaypointGetLatestStatusReportParams) WithContext(ctx context.Context) *WaypointGetLatestStatusReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint get latest status report params
func (o *WaypointGetLatestStatusReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint get latest status report params
func (o *WaypointGetLatestStatusReportParams) WithHTTPClient(client *http.Client) *WaypointGetLatestStatusReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint get latest status report params
func (o *WaypointGetLatestStatusReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationApplication adds the applicationApplication to the waypoint get latest status report params
func (o *WaypointGetLatestStatusReportParams) WithApplicationApplication(applicationApplication string) *WaypointGetLatestStatusReportParams {
	o.SetApplicationApplication(applicationApplication)
	return o
}

// SetApplicationApplication adds the applicationApplication to the waypoint get latest status report params
func (o *WaypointGetLatestStatusReportParams) SetApplicationApplication(applicationApplication string) {
	o.ApplicationApplication = applicationApplication
}

// WithApplicationProject adds the applicationProject to the waypoint get latest status report params
func (o *WaypointGetLatestStatusReportParams) WithApplicationProject(applicationProject string) *WaypointGetLatestStatusReportParams {
	o.SetApplicationProject(applicationProject)
	return o
}

// SetApplicationProject adds the applicationProject to the waypoint get latest status report params
func (o *WaypointGetLatestStatusReportParams) SetApplicationProject(applicationProject string) {
	o.ApplicationProject = applicationProject
}

// WithDeploymentID adds the deploymentID to the waypoint get latest status report params
func (o *WaypointGetLatestStatusReportParams) WithDeploymentID(deploymentID *string) *WaypointGetLatestStatusReportParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the waypoint get latest status report params
func (o *WaypointGetLatestStatusReportParams) SetDeploymentID(deploymentID *string) {
	o.DeploymentID = deploymentID
}

// WithNamespaceID adds the namespaceID to the waypoint get latest status report params
func (o *WaypointGetLatestStatusReportParams) WithNamespaceID(namespaceID string) *WaypointGetLatestStatusReportParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint get latest status report params
func (o *WaypointGetLatestStatusReportParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithReleaseID adds the releaseID to the waypoint get latest status report params
func (o *WaypointGetLatestStatusReportParams) WithReleaseID(releaseID *string) *WaypointGetLatestStatusReportParams {
	o.SetReleaseID(releaseID)
	return o
}

// SetReleaseID adds the releaseId to the waypoint get latest status report params
func (o *WaypointGetLatestStatusReportParams) SetReleaseID(releaseID *string) {
	o.ReleaseID = releaseID
}

// WithWorkspaceWorkspace adds the workspaceWorkspace to the waypoint get latest status report params
func (o *WaypointGetLatestStatusReportParams) WithWorkspaceWorkspace(workspaceWorkspace *string) *WaypointGetLatestStatusReportParams {
	o.SetWorkspaceWorkspace(workspaceWorkspace)
	return o
}

// SetWorkspaceWorkspace adds the workspaceWorkspace to the waypoint get latest status report params
func (o *WaypointGetLatestStatusReportParams) SetWorkspaceWorkspace(workspaceWorkspace *string) {
	o.WorkspaceWorkspace = workspaceWorkspace
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointGetLatestStatusReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DeploymentID != nil {

		// query param deployment_id
		var qrDeploymentID string
		if o.DeploymentID != nil {
			qrDeploymentID = *o.DeploymentID
		}
		qDeploymentID := qrDeploymentID
		if qDeploymentID != "" {
			if err := r.SetQueryParam("deployment_id", qDeploymentID); err != nil {
				return err
			}
		}

	}

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

	if o.ReleaseID != nil {

		// query param release_id
		var qrReleaseID string
		if o.ReleaseID != nil {
			qrReleaseID = *o.ReleaseID
		}
		qReleaseID := qrReleaseID
		if qReleaseID != "" {
			if err := r.SetQueryParam("release_id", qReleaseID); err != nil {
				return err
			}
		}

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
