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
	"github.com/go-openapi/swag"
)

// NewWaypointUIListDeployments2Params creates a new WaypointUIListDeployments2Params object
// with the default values initialized.
func NewWaypointUIListDeployments2Params() *WaypointUIListDeployments2Params {
	var (
		orderOrderDefault    = string("UNSET")
		physicalStateDefault = string("UNKNOWN")
	)
	return &WaypointUIListDeployments2Params{
		OrderOrder:    &orderOrderDefault,
		PhysicalState: &physicalStateDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointUIListDeployments2ParamsWithTimeout creates a new WaypointUIListDeployments2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointUIListDeployments2ParamsWithTimeout(timeout time.Duration) *WaypointUIListDeployments2Params {
	var (
		orderOrderDefault    = string("UNSET")
		physicalStateDefault = string("UNKNOWN")
	)
	return &WaypointUIListDeployments2Params{
		OrderOrder:    &orderOrderDefault,
		PhysicalState: &physicalStateDefault,

		timeout: timeout,
	}
}

// NewWaypointUIListDeployments2ParamsWithContext creates a new WaypointUIListDeployments2Params object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointUIListDeployments2ParamsWithContext(ctx context.Context) *WaypointUIListDeployments2Params {
	var (
		orderOrderDefault    = string("UNSET")
		physicalStateDefault = string("UNKNOWN")
	)
	return &WaypointUIListDeployments2Params{
		OrderOrder:    &orderOrderDefault,
		PhysicalState: &physicalStateDefault,

		Context: ctx,
	}
}

// NewWaypointUIListDeployments2ParamsWithHTTPClient creates a new WaypointUIListDeployments2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointUIListDeployments2ParamsWithHTTPClient(client *http.Client) *WaypointUIListDeployments2Params {
	var (
		orderOrderDefault    = string("UNSET")
		physicalStateDefault = string("UNKNOWN")
	)
	return &WaypointUIListDeployments2Params{
		OrderOrder:    &orderOrderDefault,
		PhysicalState: &physicalStateDefault,
		HTTPClient:    client,
	}
}

/*WaypointUIListDeployments2Params contains all the parameters to send to the API endpoint
for the waypoint UI list deployments2 operation typically these are written to a http.Request
*/
type WaypointUIListDeployments2Params struct {

	/*ApplicationApplication*/
	ApplicationApplication string
	/*ApplicationProject*/
	ApplicationProject *string
	/*OrderDesc*/
	OrderDesc *bool
	/*OrderLimit
	  Limit the number of results.

	*/
	OrderLimit *int64
	/*OrderOrder
	  Order for the results.

	*/
	OrderOrder *string
	/*PhysicalState
	  The physical state to filter for. If this is zero or unset then no
	filtering on physical state will be done.

	*/
	PhysicalState *string
	/*WorkspaceWorkspace*/
	WorkspaceWorkspace *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint UI list deployments2 params
func (o *WaypointUIListDeployments2Params) WithTimeout(timeout time.Duration) *WaypointUIListDeployments2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint UI list deployments2 params
func (o *WaypointUIListDeployments2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint UI list deployments2 params
func (o *WaypointUIListDeployments2Params) WithContext(ctx context.Context) *WaypointUIListDeployments2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint UI list deployments2 params
func (o *WaypointUIListDeployments2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint UI list deployments2 params
func (o *WaypointUIListDeployments2Params) WithHTTPClient(client *http.Client) *WaypointUIListDeployments2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint UI list deployments2 params
func (o *WaypointUIListDeployments2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationApplication adds the applicationApplication to the waypoint UI list deployments2 params
func (o *WaypointUIListDeployments2Params) WithApplicationApplication(applicationApplication string) *WaypointUIListDeployments2Params {
	o.SetApplicationApplication(applicationApplication)
	return o
}

// SetApplicationApplication adds the applicationApplication to the waypoint UI list deployments2 params
func (o *WaypointUIListDeployments2Params) SetApplicationApplication(applicationApplication string) {
	o.ApplicationApplication = applicationApplication
}

// WithApplicationProject adds the applicationProject to the waypoint UI list deployments2 params
func (o *WaypointUIListDeployments2Params) WithApplicationProject(applicationProject *string) *WaypointUIListDeployments2Params {
	o.SetApplicationProject(applicationProject)
	return o
}

// SetApplicationProject adds the applicationProject to the waypoint UI list deployments2 params
func (o *WaypointUIListDeployments2Params) SetApplicationProject(applicationProject *string) {
	o.ApplicationProject = applicationProject
}

// WithOrderDesc adds the orderDesc to the waypoint UI list deployments2 params
func (o *WaypointUIListDeployments2Params) WithOrderDesc(orderDesc *bool) *WaypointUIListDeployments2Params {
	o.SetOrderDesc(orderDesc)
	return o
}

// SetOrderDesc adds the orderDesc to the waypoint UI list deployments2 params
func (o *WaypointUIListDeployments2Params) SetOrderDesc(orderDesc *bool) {
	o.OrderDesc = orderDesc
}

// WithOrderLimit adds the orderLimit to the waypoint UI list deployments2 params
func (o *WaypointUIListDeployments2Params) WithOrderLimit(orderLimit *int64) *WaypointUIListDeployments2Params {
	o.SetOrderLimit(orderLimit)
	return o
}

// SetOrderLimit adds the orderLimit to the waypoint UI list deployments2 params
func (o *WaypointUIListDeployments2Params) SetOrderLimit(orderLimit *int64) {
	o.OrderLimit = orderLimit
}

// WithOrderOrder adds the orderOrder to the waypoint UI list deployments2 params
func (o *WaypointUIListDeployments2Params) WithOrderOrder(orderOrder *string) *WaypointUIListDeployments2Params {
	o.SetOrderOrder(orderOrder)
	return o
}

// SetOrderOrder adds the orderOrder to the waypoint UI list deployments2 params
func (o *WaypointUIListDeployments2Params) SetOrderOrder(orderOrder *string) {
	o.OrderOrder = orderOrder
}

// WithPhysicalState adds the physicalState to the waypoint UI list deployments2 params
func (o *WaypointUIListDeployments2Params) WithPhysicalState(physicalState *string) *WaypointUIListDeployments2Params {
	o.SetPhysicalState(physicalState)
	return o
}

// SetPhysicalState adds the physicalState to the waypoint UI list deployments2 params
func (o *WaypointUIListDeployments2Params) SetPhysicalState(physicalState *string) {
	o.PhysicalState = physicalState
}

// WithWorkspaceWorkspace adds the workspaceWorkspace to the waypoint UI list deployments2 params
func (o *WaypointUIListDeployments2Params) WithWorkspaceWorkspace(workspaceWorkspace *string) *WaypointUIListDeployments2Params {
	o.SetWorkspaceWorkspace(workspaceWorkspace)
	return o
}

// SetWorkspaceWorkspace adds the workspaceWorkspace to the waypoint UI list deployments2 params
func (o *WaypointUIListDeployments2Params) SetWorkspaceWorkspace(workspaceWorkspace *string) {
	o.WorkspaceWorkspace = workspaceWorkspace
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointUIListDeployments2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param application.application
	if err := r.SetPathParam("application.application", o.ApplicationApplication); err != nil {
		return err
	}

	if o.ApplicationProject != nil {

		// query param application.project
		var qrApplicationProject string
		if o.ApplicationProject != nil {
			qrApplicationProject = *o.ApplicationProject
		}
		qApplicationProject := qrApplicationProject
		if qApplicationProject != "" {
			if err := r.SetQueryParam("application.project", qApplicationProject); err != nil {
				return err
			}
		}

	}

	if o.OrderDesc != nil {

		// query param order.desc
		var qrOrderDesc bool
		if o.OrderDesc != nil {
			qrOrderDesc = *o.OrderDesc
		}
		qOrderDesc := swag.FormatBool(qrOrderDesc)
		if qOrderDesc != "" {
			if err := r.SetQueryParam("order.desc", qOrderDesc); err != nil {
				return err
			}
		}

	}

	if o.OrderLimit != nil {

		// query param order.limit
		var qrOrderLimit int64
		if o.OrderLimit != nil {
			qrOrderLimit = *o.OrderLimit
		}
		qOrderLimit := swag.FormatInt64(qrOrderLimit)
		if qOrderLimit != "" {
			if err := r.SetQueryParam("order.limit", qOrderLimit); err != nil {
				return err
			}
		}

	}

	if o.OrderOrder != nil {

		// query param order.order
		var qrOrderOrder string
		if o.OrderOrder != nil {
			qrOrderOrder = *o.OrderOrder
		}
		qOrderOrder := qrOrderOrder
		if qOrderOrder != "" {
			if err := r.SetQueryParam("order.order", qOrderOrder); err != nil {
				return err
			}
		}

	}

	if o.PhysicalState != nil {

		// query param physical_state
		var qrPhysicalState string
		if o.PhysicalState != nil {
			qrPhysicalState = *o.PhysicalState
		}
		qPhysicalState := qrPhysicalState
		if qPhysicalState != "" {
			if err := r.SetQueryParam("physical_state", qPhysicalState); err != nil {
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
