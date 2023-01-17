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

// NewWaypointGetBuildParams creates a new WaypointGetBuildParams object
// with the default values initialized.
func NewWaypointGetBuildParams() *WaypointGetBuildParams {
	var ()
	return &WaypointGetBuildParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointGetBuildParamsWithTimeout creates a new WaypointGetBuildParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointGetBuildParamsWithTimeout(timeout time.Duration) *WaypointGetBuildParams {
	var ()
	return &WaypointGetBuildParams{

		timeout: timeout,
	}
}

// NewWaypointGetBuildParamsWithContext creates a new WaypointGetBuildParams object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointGetBuildParamsWithContext(ctx context.Context) *WaypointGetBuildParams {
	var ()
	return &WaypointGetBuildParams{

		Context: ctx,
	}
}

// NewWaypointGetBuildParamsWithHTTPClient creates a new WaypointGetBuildParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointGetBuildParamsWithHTTPClient(client *http.Client) *WaypointGetBuildParams {
	var ()
	return &WaypointGetBuildParams{
		HTTPClient: client,
	}
}

/*WaypointGetBuildParams contains all the parameters to send to the API endpoint
for the waypoint get build operation typically these are written to a http.Request
*/
type WaypointGetBuildParams struct {

	/*NamespaceID
	  namespace_id is the id of the namespace.

	*/
	NamespaceID string
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

// WithTimeout adds the timeout to the waypoint get build params
func (o *WaypointGetBuildParams) WithTimeout(timeout time.Duration) *WaypointGetBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint get build params
func (o *WaypointGetBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint get build params
func (o *WaypointGetBuildParams) WithContext(ctx context.Context) *WaypointGetBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint get build params
func (o *WaypointGetBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint get build params
func (o *WaypointGetBuildParams) WithHTTPClient(client *http.Client) *WaypointGetBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint get build params
func (o *WaypointGetBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespaceID adds the namespaceID to the waypoint get build params
func (o *WaypointGetBuildParams) WithNamespaceID(namespaceID string) *WaypointGetBuildParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint get build params
func (o *WaypointGetBuildParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithRefID adds the refID to the waypoint get build params
func (o *WaypointGetBuildParams) WithRefID(refID string) *WaypointGetBuildParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the waypoint get build params
func (o *WaypointGetBuildParams) SetRefID(refID string) {
	o.RefID = refID
}

// WithRefSequenceApplicationApplication adds the refSequenceApplicationApplication to the waypoint get build params
func (o *WaypointGetBuildParams) WithRefSequenceApplicationApplication(refSequenceApplicationApplication *string) *WaypointGetBuildParams {
	o.SetRefSequenceApplicationApplication(refSequenceApplicationApplication)
	return o
}

// SetRefSequenceApplicationApplication adds the refSequenceApplicationApplication to the waypoint get build params
func (o *WaypointGetBuildParams) SetRefSequenceApplicationApplication(refSequenceApplicationApplication *string) {
	o.RefSequenceApplicationApplication = refSequenceApplicationApplication
}

// WithRefSequenceApplicationProject adds the refSequenceApplicationProject to the waypoint get build params
func (o *WaypointGetBuildParams) WithRefSequenceApplicationProject(refSequenceApplicationProject *string) *WaypointGetBuildParams {
	o.SetRefSequenceApplicationProject(refSequenceApplicationProject)
	return o
}

// SetRefSequenceApplicationProject adds the refSequenceApplicationProject to the waypoint get build params
func (o *WaypointGetBuildParams) SetRefSequenceApplicationProject(refSequenceApplicationProject *string) {
	o.RefSequenceApplicationProject = refSequenceApplicationProject
}

// WithRefSequenceNumber adds the refSequenceNumber to the waypoint get build params
func (o *WaypointGetBuildParams) WithRefSequenceNumber(refSequenceNumber *string) *WaypointGetBuildParams {
	o.SetRefSequenceNumber(refSequenceNumber)
	return o
}

// SetRefSequenceNumber adds the refSequenceNumber to the waypoint get build params
func (o *WaypointGetBuildParams) SetRefSequenceNumber(refSequenceNumber *string) {
	o.RefSequenceNumber = refSequenceNumber
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointGetBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

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
