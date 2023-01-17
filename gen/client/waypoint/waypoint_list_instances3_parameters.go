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

// NewWaypointListInstances3Params creates a new WaypointListInstances3Params object
// with the default values initialized.
func NewWaypointListInstances3Params() *WaypointListInstances3Params {
	var ()
	return &WaypointListInstances3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointListInstances3ParamsWithTimeout creates a new WaypointListInstances3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointListInstances3ParamsWithTimeout(timeout time.Duration) *WaypointListInstances3Params {
	var ()
	return &WaypointListInstances3Params{

		timeout: timeout,
	}
}

// NewWaypointListInstances3ParamsWithContext creates a new WaypointListInstances3Params object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointListInstances3ParamsWithContext(ctx context.Context) *WaypointListInstances3Params {
	var ()
	return &WaypointListInstances3Params{

		Context: ctx,
	}
}

// NewWaypointListInstances3ParamsWithHTTPClient creates a new WaypointListInstances3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointListInstances3ParamsWithHTTPClient(client *http.Client) *WaypointListInstances3Params {
	var ()
	return &WaypointListInstances3Params{
		HTTPClient: client,
	}
}

/*WaypointListInstances3Params contains all the parameters to send to the API endpoint
for the waypoint list instances3 operation typically these are written to a http.Request
*/
type WaypointListInstances3Params struct {

	/*ApplicationApplicationApplication*/
	ApplicationApplicationApplication string
	/*ApplicationApplicationProject*/
	ApplicationApplicationProject string
	/*ApplicationWorkspaceWorkspace*/
	ApplicationWorkspaceWorkspace string
	/*DeploymentID
	  List instances for a specific deployment.

	*/
	DeploymentID *string
	/*NamespaceID
	  namespace_id is the id of the namespace.

	*/
	NamespaceID string
	/*WaitTimeout
	  Time to wait before retrying a request to connect to requested instance.

	*/
	WaitTimeout *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint list instances3 params
func (o *WaypointListInstances3Params) WithTimeout(timeout time.Duration) *WaypointListInstances3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint list instances3 params
func (o *WaypointListInstances3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint list instances3 params
func (o *WaypointListInstances3Params) WithContext(ctx context.Context) *WaypointListInstances3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint list instances3 params
func (o *WaypointListInstances3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint list instances3 params
func (o *WaypointListInstances3Params) WithHTTPClient(client *http.Client) *WaypointListInstances3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint list instances3 params
func (o *WaypointListInstances3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationApplicationApplication adds the applicationApplicationApplication to the waypoint list instances3 params
func (o *WaypointListInstances3Params) WithApplicationApplicationApplication(applicationApplicationApplication string) *WaypointListInstances3Params {
	o.SetApplicationApplicationApplication(applicationApplicationApplication)
	return o
}

// SetApplicationApplicationApplication adds the applicationApplicationApplication to the waypoint list instances3 params
func (o *WaypointListInstances3Params) SetApplicationApplicationApplication(applicationApplicationApplication string) {
	o.ApplicationApplicationApplication = applicationApplicationApplication
}

// WithApplicationApplicationProject adds the applicationApplicationProject to the waypoint list instances3 params
func (o *WaypointListInstances3Params) WithApplicationApplicationProject(applicationApplicationProject string) *WaypointListInstances3Params {
	o.SetApplicationApplicationProject(applicationApplicationProject)
	return o
}

// SetApplicationApplicationProject adds the applicationApplicationProject to the waypoint list instances3 params
func (o *WaypointListInstances3Params) SetApplicationApplicationProject(applicationApplicationProject string) {
	o.ApplicationApplicationProject = applicationApplicationProject
}

// WithApplicationWorkspaceWorkspace adds the applicationWorkspaceWorkspace to the waypoint list instances3 params
func (o *WaypointListInstances3Params) WithApplicationWorkspaceWorkspace(applicationWorkspaceWorkspace string) *WaypointListInstances3Params {
	o.SetApplicationWorkspaceWorkspace(applicationWorkspaceWorkspace)
	return o
}

// SetApplicationWorkspaceWorkspace adds the applicationWorkspaceWorkspace to the waypoint list instances3 params
func (o *WaypointListInstances3Params) SetApplicationWorkspaceWorkspace(applicationWorkspaceWorkspace string) {
	o.ApplicationWorkspaceWorkspace = applicationWorkspaceWorkspace
}

// WithDeploymentID adds the deploymentID to the waypoint list instances3 params
func (o *WaypointListInstances3Params) WithDeploymentID(deploymentID *string) *WaypointListInstances3Params {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the waypoint list instances3 params
func (o *WaypointListInstances3Params) SetDeploymentID(deploymentID *string) {
	o.DeploymentID = deploymentID
}

// WithNamespaceID adds the namespaceID to the waypoint list instances3 params
func (o *WaypointListInstances3Params) WithNamespaceID(namespaceID string) *WaypointListInstances3Params {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint list instances3 params
func (o *WaypointListInstances3Params) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithWaitTimeout adds the waitTimeout to the waypoint list instances3 params
func (o *WaypointListInstances3Params) WithWaitTimeout(waitTimeout *string) *WaypointListInstances3Params {
	o.SetWaitTimeout(waitTimeout)
	return o
}

// SetWaitTimeout adds the waitTimeout to the waypoint list instances3 params
func (o *WaypointListInstances3Params) SetWaitTimeout(waitTimeout *string) {
	o.WaitTimeout = waitTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointListInstances3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param application.workspace.workspace
	if err := r.SetPathParam("application.workspace.workspace", o.ApplicationWorkspaceWorkspace); err != nil {
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
