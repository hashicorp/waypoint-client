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

// NewWaypointGetTaskParams creates a new WaypointGetTaskParams object
// with the default values initialized.
func NewWaypointGetTaskParams() *WaypointGetTaskParams {
	var ()
	return &WaypointGetTaskParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointGetTaskParamsWithTimeout creates a new WaypointGetTaskParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointGetTaskParamsWithTimeout(timeout time.Duration) *WaypointGetTaskParams {
	var ()
	return &WaypointGetTaskParams{

		timeout: timeout,
	}
}

// NewWaypointGetTaskParamsWithContext creates a new WaypointGetTaskParams object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointGetTaskParamsWithContext(ctx context.Context) *WaypointGetTaskParams {
	var ()
	return &WaypointGetTaskParams{

		Context: ctx,
	}
}

// NewWaypointGetTaskParamsWithHTTPClient creates a new WaypointGetTaskParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointGetTaskParamsWithHTTPClient(client *http.Client) *WaypointGetTaskParams {
	var ()
	return &WaypointGetTaskParams{
		HTTPClient: client,
	}
}

/*WaypointGetTaskParams contains all the parameters to send to the API endpoint
for the waypoint get task operation typically these are written to a http.Request
*/
type WaypointGetTaskParams struct {

	/*RefID
	  the id of the tracktask record

	*/
	RefID string
	/*RefJobID
	  The main "run" job ID that the task initiated.

	*/
	RefJobID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint get task params
func (o *WaypointGetTaskParams) WithTimeout(timeout time.Duration) *WaypointGetTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint get task params
func (o *WaypointGetTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint get task params
func (o *WaypointGetTaskParams) WithContext(ctx context.Context) *WaypointGetTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint get task params
func (o *WaypointGetTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint get task params
func (o *WaypointGetTaskParams) WithHTTPClient(client *http.Client) *WaypointGetTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint get task params
func (o *WaypointGetTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRefID adds the refID to the waypoint get task params
func (o *WaypointGetTaskParams) WithRefID(refID string) *WaypointGetTaskParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the waypoint get task params
func (o *WaypointGetTaskParams) SetRefID(refID string) {
	o.RefID = refID
}

// WithRefJobID adds the refJobID to the waypoint get task params
func (o *WaypointGetTaskParams) WithRefJobID(refJobID *string) *WaypointGetTaskParams {
	o.SetRefJobID(refJobID)
	return o
}

// SetRefJobID adds the refJobId to the waypoint get task params
func (o *WaypointGetTaskParams) SetRefJobID(refJobID *string) {
	o.RefJobID = refJobID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointGetTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ref.id
	if err := r.SetPathParam("ref.id", o.RefID); err != nil {
		return err
	}

	if o.RefJobID != nil {

		// query param ref.job_id
		var qrRefJobID string
		if o.RefJobID != nil {
			qrRefJobID = *o.RefJobID
		}
		qRefJobID := qrRefJobID
		if qRefJobID != "" {
			if err := r.SetQueryParam("ref.job_id", qRefJobID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}