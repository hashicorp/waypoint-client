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

// NewWaypointGetStatusReportParams creates a new WaypointGetStatusReportParams object
// with the default values initialized.
func NewWaypointGetStatusReportParams() *WaypointGetStatusReportParams {
	var ()
	return &WaypointGetStatusReportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointGetStatusReportParamsWithTimeout creates a new WaypointGetStatusReportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointGetStatusReportParamsWithTimeout(timeout time.Duration) *WaypointGetStatusReportParams {
	var ()
	return &WaypointGetStatusReportParams{

		timeout: timeout,
	}
}

// NewWaypointGetStatusReportParamsWithContext creates a new WaypointGetStatusReportParams object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointGetStatusReportParamsWithContext(ctx context.Context) *WaypointGetStatusReportParams {
	var ()
	return &WaypointGetStatusReportParams{

		Context: ctx,
	}
}

// NewWaypointGetStatusReportParamsWithHTTPClient creates a new WaypointGetStatusReportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointGetStatusReportParamsWithHTTPClient(client *http.Client) *WaypointGetStatusReportParams {
	var ()
	return &WaypointGetStatusReportParams{
		HTTPClient: client,
	}
}

/*WaypointGetStatusReportParams contains all the parameters to send to the API endpoint
for the waypoint get status report operation typically these are written to a http.Request
*/
type WaypointGetStatusReportParams struct {

	/*RefID*/
	RefID string
	/*RefSequenceApplicationApplication*/
	RefSequenceApplicationApplication *string
	/*RefSequenceApplicationProject*/
	RefSequenceApplicationProject *string
	/*RefSequenceNumber*/
	RefSequenceNumber *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint get status report params
func (o *WaypointGetStatusReportParams) WithTimeout(timeout time.Duration) *WaypointGetStatusReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint get status report params
func (o *WaypointGetStatusReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint get status report params
func (o *WaypointGetStatusReportParams) WithContext(ctx context.Context) *WaypointGetStatusReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint get status report params
func (o *WaypointGetStatusReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint get status report params
func (o *WaypointGetStatusReportParams) WithHTTPClient(client *http.Client) *WaypointGetStatusReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint get status report params
func (o *WaypointGetStatusReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRefID adds the refID to the waypoint get status report params
func (o *WaypointGetStatusReportParams) WithRefID(refID string) *WaypointGetStatusReportParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the waypoint get status report params
func (o *WaypointGetStatusReportParams) SetRefID(refID string) {
	o.RefID = refID
}

// WithRefSequenceApplicationApplication adds the refSequenceApplicationApplication to the waypoint get status report params
func (o *WaypointGetStatusReportParams) WithRefSequenceApplicationApplication(refSequenceApplicationApplication *string) *WaypointGetStatusReportParams {
	o.SetRefSequenceApplicationApplication(refSequenceApplicationApplication)
	return o
}

// SetRefSequenceApplicationApplication adds the refSequenceApplicationApplication to the waypoint get status report params
func (o *WaypointGetStatusReportParams) SetRefSequenceApplicationApplication(refSequenceApplicationApplication *string) {
	o.RefSequenceApplicationApplication = refSequenceApplicationApplication
}

// WithRefSequenceApplicationProject adds the refSequenceApplicationProject to the waypoint get status report params
func (o *WaypointGetStatusReportParams) WithRefSequenceApplicationProject(refSequenceApplicationProject *string) *WaypointGetStatusReportParams {
	o.SetRefSequenceApplicationProject(refSequenceApplicationProject)
	return o
}

// SetRefSequenceApplicationProject adds the refSequenceApplicationProject to the waypoint get status report params
func (o *WaypointGetStatusReportParams) SetRefSequenceApplicationProject(refSequenceApplicationProject *string) {
	o.RefSequenceApplicationProject = refSequenceApplicationProject
}

// WithRefSequenceNumber adds the refSequenceNumber to the waypoint get status report params
func (o *WaypointGetStatusReportParams) WithRefSequenceNumber(refSequenceNumber *string) *WaypointGetStatusReportParams {
	o.SetRefSequenceNumber(refSequenceNumber)
	return o
}

// SetRefSequenceNumber adds the refSequenceNumber to the waypoint get status report params
func (o *WaypointGetStatusReportParams) SetRefSequenceNumber(refSequenceNumber *string) {
	o.RefSequenceNumber = refSequenceNumber
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointGetStatusReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ref.id
	if err := r.SetPathParam("ref.id", o.RefID); err != nil {
		return err
	}

	if o.RefSequenceApplicationApplication != nil {

		// query param ref.sequence.application.application
		var qrRefSequenceApplicationApplication string
		if o.RefSequenceApplicationApplication != nil {
			qrRefSequenceApplicationApplication = *o.RefSequenceApplicationApplication
		}
		qRefSequenceApplicationApplication := qrRefSequenceApplicationApplication
		if qRefSequenceApplicationApplication != "" {
			if err := r.SetQueryParam("ref.sequence.application.application", qRefSequenceApplicationApplication); err != nil {
				return err
			}
		}

	}

	if o.RefSequenceApplicationProject != nil {

		// query param ref.sequence.application.project
		var qrRefSequenceApplicationProject string
		if o.RefSequenceApplicationProject != nil {
			qrRefSequenceApplicationProject = *o.RefSequenceApplicationProject
		}
		qRefSequenceApplicationProject := qrRefSequenceApplicationProject
		if qRefSequenceApplicationProject != "" {
			if err := r.SetQueryParam("ref.sequence.application.project", qRefSequenceApplicationProject); err != nil {
				return err
			}
		}

	}

	if o.RefSequenceNumber != nil {

		// query param ref.sequence.number
		var qrRefSequenceNumber string
		if o.RefSequenceNumber != nil {
			qrRefSequenceNumber = *o.RefSequenceNumber
		}
		qRefSequenceNumber := qrRefSequenceNumber
		if qRefSequenceNumber != "" {
			if err := r.SetQueryParam("ref.sequence.number", qRefSequenceNumber); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
