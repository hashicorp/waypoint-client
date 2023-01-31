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

// NewWaypointUIListReleasesParams creates a new WaypointUIListReleasesParams object
// with the default values initialized.
func NewWaypointUIListReleasesParams() *WaypointUIListReleasesParams {
	var (
		orderOrderDefault    = string("UNSET")
		physicalStateDefault = string("UNKNOWN")
	)
	return &WaypointUIListReleasesParams{
		OrderOrder:    &orderOrderDefault,
		PhysicalState: &physicalStateDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointUIListReleasesParamsWithTimeout creates a new WaypointUIListReleasesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointUIListReleasesParamsWithTimeout(timeout time.Duration) *WaypointUIListReleasesParams {
	var (
		orderOrderDefault    = string("UNSET")
		physicalStateDefault = string("UNKNOWN")
	)
	return &WaypointUIListReleasesParams{
		OrderOrder:    &orderOrderDefault,
		PhysicalState: &physicalStateDefault,

		timeout: timeout,
	}
}

// NewWaypointUIListReleasesParamsWithContext creates a new WaypointUIListReleasesParams object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointUIListReleasesParamsWithContext(ctx context.Context) *WaypointUIListReleasesParams {
	var (
		orderOrderDefault    = string("UNSET")
		physicalStateDefault = string("UNKNOWN")
	)
	return &WaypointUIListReleasesParams{
		OrderOrder:    &orderOrderDefault,
		PhysicalState: &physicalStateDefault,

		Context: ctx,
	}
}

// NewWaypointUIListReleasesParamsWithHTTPClient creates a new WaypointUIListReleasesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointUIListReleasesParamsWithHTTPClient(client *http.Client) *WaypointUIListReleasesParams {
	var (
		orderOrderDefault    = string("UNSET")
		physicalStateDefault = string("UNKNOWN")
	)
	return &WaypointUIListReleasesParams{
		OrderOrder:    &orderOrderDefault,
		PhysicalState: &physicalStateDefault,
		HTTPClient:    client,
	}
}

/*WaypointUIListReleasesParams contains all the parameters to send to the API endpoint
for the waypoint UI list releases operation typically these are written to a http.Request
*/
type WaypointUIListReleasesParams struct {

	/*ApplicationApplication*/
	ApplicationApplication *string
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
	WorkspaceWorkspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint UI list releases params
func (o *WaypointUIListReleasesParams) WithTimeout(timeout time.Duration) *WaypointUIListReleasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint UI list releases params
func (o *WaypointUIListReleasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint UI list releases params
func (o *WaypointUIListReleasesParams) WithContext(ctx context.Context) *WaypointUIListReleasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint UI list releases params
func (o *WaypointUIListReleasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint UI list releases params
func (o *WaypointUIListReleasesParams) WithHTTPClient(client *http.Client) *WaypointUIListReleasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint UI list releases params
func (o *WaypointUIListReleasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationApplication adds the applicationApplication to the waypoint UI list releases params
func (o *WaypointUIListReleasesParams) WithApplicationApplication(applicationApplication *string) *WaypointUIListReleasesParams {
	o.SetApplicationApplication(applicationApplication)
	return o
}

// SetApplicationApplication adds the applicationApplication to the waypoint UI list releases params
func (o *WaypointUIListReleasesParams) SetApplicationApplication(applicationApplication *string) {
	o.ApplicationApplication = applicationApplication
}

// WithApplicationProject adds the applicationProject to the waypoint UI list releases params
func (o *WaypointUIListReleasesParams) WithApplicationProject(applicationProject *string) *WaypointUIListReleasesParams {
	o.SetApplicationProject(applicationProject)
	return o
}

// SetApplicationProject adds the applicationProject to the waypoint UI list releases params
func (o *WaypointUIListReleasesParams) SetApplicationProject(applicationProject *string) {
	o.ApplicationProject = applicationProject
}

// WithOrderDesc adds the orderDesc to the waypoint UI list releases params
func (o *WaypointUIListReleasesParams) WithOrderDesc(orderDesc *bool) *WaypointUIListReleasesParams {
	o.SetOrderDesc(orderDesc)
	return o
}

// SetOrderDesc adds the orderDesc to the waypoint UI list releases params
func (o *WaypointUIListReleasesParams) SetOrderDesc(orderDesc *bool) {
	o.OrderDesc = orderDesc
}

// WithOrderLimit adds the orderLimit to the waypoint UI list releases params
func (o *WaypointUIListReleasesParams) WithOrderLimit(orderLimit *int64) *WaypointUIListReleasesParams {
	o.SetOrderLimit(orderLimit)
	return o
}

// SetOrderLimit adds the orderLimit to the waypoint UI list releases params
func (o *WaypointUIListReleasesParams) SetOrderLimit(orderLimit *int64) {
	o.OrderLimit = orderLimit
}

// WithOrderOrder adds the orderOrder to the waypoint UI list releases params
func (o *WaypointUIListReleasesParams) WithOrderOrder(orderOrder *string) *WaypointUIListReleasesParams {
	o.SetOrderOrder(orderOrder)
	return o
}

// SetOrderOrder adds the orderOrder to the waypoint UI list releases params
func (o *WaypointUIListReleasesParams) SetOrderOrder(orderOrder *string) {
	o.OrderOrder = orderOrder
}

// WithPhysicalState adds the physicalState to the waypoint UI list releases params
func (o *WaypointUIListReleasesParams) WithPhysicalState(physicalState *string) *WaypointUIListReleasesParams {
	o.SetPhysicalState(physicalState)
	return o
}

// SetPhysicalState adds the physicalState to the waypoint UI list releases params
func (o *WaypointUIListReleasesParams) SetPhysicalState(physicalState *string) {
	o.PhysicalState = physicalState
}

// WithWorkspaceWorkspace adds the workspaceWorkspace to the waypoint UI list releases params
func (o *WaypointUIListReleasesParams) WithWorkspaceWorkspace(workspaceWorkspace string) *WaypointUIListReleasesParams {
	o.SetWorkspaceWorkspace(workspaceWorkspace)
	return o
}

// SetWorkspaceWorkspace adds the workspaceWorkspace to the waypoint UI list releases params
func (o *WaypointUIListReleasesParams) SetWorkspaceWorkspace(workspaceWorkspace string) {
	o.WorkspaceWorkspace = workspaceWorkspace
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointUIListReleasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApplicationApplication != nil {

		// query param application.application
		var qrApplicationApplication string
		if o.ApplicationApplication != nil {
			qrApplicationApplication = *o.ApplicationApplication
		}
		qApplicationApplication := qrApplicationApplication
		if qApplicationApplication != "" {
			if err := r.SetQueryParam("application.application", qApplicationApplication); err != nil {
				return err
			}
		}

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

	// path param workspace.workspace
	if err := r.SetPathParam("workspace.workspace", o.WorkspaceWorkspace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}