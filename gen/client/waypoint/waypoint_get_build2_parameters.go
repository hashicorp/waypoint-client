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

// NewWaypointGetBuild2Params creates a new WaypointGetBuild2Params object
// with the default values initialized.
func NewWaypointGetBuild2Params() *WaypointGetBuild2Params {
	var ()
	return &WaypointGetBuild2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointGetBuild2ParamsWithTimeout creates a new WaypointGetBuild2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointGetBuild2ParamsWithTimeout(timeout time.Duration) *WaypointGetBuild2Params {
	var ()
	return &WaypointGetBuild2Params{

		timeout: timeout,
	}
}

// NewWaypointGetBuild2ParamsWithContext creates a new WaypointGetBuild2Params object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointGetBuild2ParamsWithContext(ctx context.Context) *WaypointGetBuild2Params {
	var ()
	return &WaypointGetBuild2Params{

		Context: ctx,
	}
}

// NewWaypointGetBuild2ParamsWithHTTPClient creates a new WaypointGetBuild2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointGetBuild2ParamsWithHTTPClient(client *http.Client) *WaypointGetBuild2Params {
	var ()
	return &WaypointGetBuild2Params{
		HTTPClient: client,
	}
}

/*WaypointGetBuild2Params contains all the parameters to send to the API endpoint
for the waypoint get build2 operation typically these are written to a http.Request
*/
type WaypointGetBuild2Params struct {

	/*RefID*/
	RefID *string
	/*RefSequenceApplicationApplication*/
	RefSequenceApplicationApplication string
	/*RefSequenceApplicationProject*/
	RefSequenceApplicationProject string
	/*RefSequenceNumber*/
	RefSequenceNumber string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint get build2 params
func (o *WaypointGetBuild2Params) WithTimeout(timeout time.Duration) *WaypointGetBuild2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint get build2 params
func (o *WaypointGetBuild2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint get build2 params
func (o *WaypointGetBuild2Params) WithContext(ctx context.Context) *WaypointGetBuild2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint get build2 params
func (o *WaypointGetBuild2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint get build2 params
func (o *WaypointGetBuild2Params) WithHTTPClient(client *http.Client) *WaypointGetBuild2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint get build2 params
func (o *WaypointGetBuild2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRefID adds the refID to the waypoint get build2 params
func (o *WaypointGetBuild2Params) WithRefID(refID *string) *WaypointGetBuild2Params {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the waypoint get build2 params
func (o *WaypointGetBuild2Params) SetRefID(refID *string) {
	o.RefID = refID
}

// WithRefSequenceApplicationApplication adds the refSequenceApplicationApplication to the waypoint get build2 params
func (o *WaypointGetBuild2Params) WithRefSequenceApplicationApplication(refSequenceApplicationApplication string) *WaypointGetBuild2Params {
	o.SetRefSequenceApplicationApplication(refSequenceApplicationApplication)
	return o
}

// SetRefSequenceApplicationApplication adds the refSequenceApplicationApplication to the waypoint get build2 params
func (o *WaypointGetBuild2Params) SetRefSequenceApplicationApplication(refSequenceApplicationApplication string) {
	o.RefSequenceApplicationApplication = refSequenceApplicationApplication
}

// WithRefSequenceApplicationProject adds the refSequenceApplicationProject to the waypoint get build2 params
func (o *WaypointGetBuild2Params) WithRefSequenceApplicationProject(refSequenceApplicationProject string) *WaypointGetBuild2Params {
	o.SetRefSequenceApplicationProject(refSequenceApplicationProject)
	return o
}

// SetRefSequenceApplicationProject adds the refSequenceApplicationProject to the waypoint get build2 params
func (o *WaypointGetBuild2Params) SetRefSequenceApplicationProject(refSequenceApplicationProject string) {
	o.RefSequenceApplicationProject = refSequenceApplicationProject
}

// WithRefSequenceNumber adds the refSequenceNumber to the waypoint get build2 params
func (o *WaypointGetBuild2Params) WithRefSequenceNumber(refSequenceNumber string) *WaypointGetBuild2Params {
	o.SetRefSequenceNumber(refSequenceNumber)
	return o
}

// SetRefSequenceNumber adds the refSequenceNumber to the waypoint get build2 params
func (o *WaypointGetBuild2Params) SetRefSequenceNumber(refSequenceNumber string) {
	o.RefSequenceNumber = refSequenceNumber
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointGetBuild2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RefID != nil {

		// query param ref.id
		var qrRefID string
		if o.RefID != nil {
			qrRefID = *o.RefID
		}
		qRefID := qrRefID
		if qRefID != "" {
			if err := r.SetQueryParam("ref.id", qRefID); err != nil {
				return err
			}
		}

	}

	// path param ref.sequence.application.application
	if err := r.SetPathParam("ref.sequence.application.application", o.RefSequenceApplicationApplication); err != nil {
		return err
	}

	// path param ref.sequence.application.project
	if err := r.SetPathParam("ref.sequence.application.project", o.RefSequenceApplicationProject); err != nil {
		return err
	}

	// path param ref.sequence.number
	if err := r.SetPathParam("ref.sequence.number", o.RefSequenceNumber); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
