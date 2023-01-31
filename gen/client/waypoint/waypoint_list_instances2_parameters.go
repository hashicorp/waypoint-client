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

// NewWaypointListInstances2Params creates a new WaypointListInstances2Params object
// with the default values initialized.
func NewWaypointListInstances2Params() *WaypointListInstances2Params {
	var ()
	return &WaypointListInstances2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointListInstances2ParamsWithTimeout creates a new WaypointListInstances2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointListInstances2ParamsWithTimeout(timeout time.Duration) *WaypointListInstances2Params {
	var ()
	return &WaypointListInstances2Params{

		timeout: timeout,
	}
}

// NewWaypointListInstances2ParamsWithContext creates a new WaypointListInstances2Params object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointListInstances2ParamsWithContext(ctx context.Context) *WaypointListInstances2Params {
	var ()
	return &WaypointListInstances2Params{

		Context: ctx,
	}
}

// NewWaypointListInstances2ParamsWithHTTPClient creates a new WaypointListInstances2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointListInstances2ParamsWithHTTPClient(client *http.Client) *WaypointListInstances2Params {
	var ()
	return &WaypointListInstances2Params{
		HTTPClient: client,
	}
}

/*WaypointListInstances2Params contains all the parameters to send to the API endpoint
for the waypoint list instances2 operation typically these are written to a http.Request
*/
type WaypointListInstances2Params struct {

	/*ApplicationApplicationApplication*/
	ApplicationApplicationApplication string
	/*ApplicationApplicationProject*/
	ApplicationApplicationProject string
	/*ApplicationWorkspaceWorkspace*/
	ApplicationWorkspaceWorkspace *string
	/*DeploymentID
	  List instances for a specific deployment.

	*/
	DeploymentID *string
	/*WaitTimeout
	  Time to wait before retrying a request to connect to requested instance.

	*/
	WaitTimeout *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint list instances2 params
func (o *WaypointListInstances2Params) WithTimeout(timeout time.Duration) *WaypointListInstances2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint list instances2 params
func (o *WaypointListInstances2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint list instances2 params
func (o *WaypointListInstances2Params) WithContext(ctx context.Context) *WaypointListInstances2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint list instances2 params
func (o *WaypointListInstances2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint list instances2 params
func (o *WaypointListInstances2Params) WithHTTPClient(client *http.Client) *WaypointListInstances2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint list instances2 params
func (o *WaypointListInstances2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationApplicationApplication adds the applicationApplicationApplication to the waypoint list instances2 params
func (o *WaypointListInstances2Params) WithApplicationApplicationApplication(applicationApplicationApplication string) *WaypointListInstances2Params {
	o.SetApplicationApplicationApplication(applicationApplicationApplication)
	return o
}

// SetApplicationApplicationApplication adds the applicationApplicationApplication to the waypoint list instances2 params
func (o *WaypointListInstances2Params) SetApplicationApplicationApplication(applicationApplicationApplication string) {
	o.ApplicationApplicationApplication = applicationApplicationApplication
}

// WithApplicationApplicationProject adds the applicationApplicationProject to the waypoint list instances2 params
func (o *WaypointListInstances2Params) WithApplicationApplicationProject(applicationApplicationProject string) *WaypointListInstances2Params {
	o.SetApplicationApplicationProject(applicationApplicationProject)
	return o
}

// SetApplicationApplicationProject adds the applicationApplicationProject to the waypoint list instances2 params
func (o *WaypointListInstances2Params) SetApplicationApplicationProject(applicationApplicationProject string) {
	o.ApplicationApplicationProject = applicationApplicationProject
}

// WithApplicationWorkspaceWorkspace adds the applicationWorkspaceWorkspace to the waypoint list instances2 params
func (o *WaypointListInstances2Params) WithApplicationWorkspaceWorkspace(applicationWorkspaceWorkspace *string) *WaypointListInstances2Params {
	o.SetApplicationWorkspaceWorkspace(applicationWorkspaceWorkspace)
	return o
}

// SetApplicationWorkspaceWorkspace adds the applicationWorkspaceWorkspace to the waypoint list instances2 params
func (o *WaypointListInstances2Params) SetApplicationWorkspaceWorkspace(applicationWorkspaceWorkspace *string) {
	o.ApplicationWorkspaceWorkspace = applicationWorkspaceWorkspace
}

// WithDeploymentID adds the deploymentID to the waypoint list instances2 params
func (o *WaypointListInstances2Params) WithDeploymentID(deploymentID *string) *WaypointListInstances2Params {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the waypoint list instances2 params
func (o *WaypointListInstances2Params) SetDeploymentID(deploymentID *string) {
	o.DeploymentID = deploymentID
}

// WithWaitTimeout adds the waitTimeout to the waypoint list instances2 params
func (o *WaypointListInstances2Params) WithWaitTimeout(waitTimeout *string) *WaypointListInstances2Params {
	o.SetWaitTimeout(waitTimeout)
	return o
}

// SetWaitTimeout adds the waitTimeout to the waypoint list instances2 params
func (o *WaypointListInstances2Params) SetWaitTimeout(waitTimeout *string) {
	o.WaitTimeout = waitTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointListInstances2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param application.application.application
	if err := r.SetPathParam("application.application.application", o.ApplicationApplicationApplication); err != nil {
		return err
	}

	// path param application.application.project
	if err := r.SetPathParam("application.application.project", o.ApplicationApplicationProject); err != nil {
		return err
	}

	if o.ApplicationWorkspaceWorkspace != nil {

		// query param application.workspace.workspace
		var qrApplicationWorkspaceWorkspace string
		if o.ApplicationWorkspaceWorkspace != nil {
			qrApplicationWorkspaceWorkspace = *o.ApplicationWorkspaceWorkspace
		}
		qApplicationWorkspaceWorkspace := qrApplicationWorkspaceWorkspace
		if qApplicationWorkspaceWorkspace != "" {
			if err := r.SetQueryParam("application.workspace.workspace", qApplicationWorkspaceWorkspace); err != nil {
				return err
			}
		}

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

	if o.WaitTimeout != nil {

		// query param wait_timeout
		var qrWaitTimeout string
		if o.WaitTimeout != nil {
			qrWaitTimeout = *o.WaitTimeout
		}
		qWaitTimeout := qrWaitTimeout
		if qWaitTimeout != "" {
			if err := r.SetQueryParam("wait_timeout", qWaitTimeout); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}