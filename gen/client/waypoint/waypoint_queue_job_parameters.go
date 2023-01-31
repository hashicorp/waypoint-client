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

	"github.com/hashicorp/waypoint-client/gen/models"
)

// NewWaypointQueueJobParams creates a new WaypointQueueJobParams object
// with the default values initialized.
func NewWaypointQueueJobParams() *WaypointQueueJobParams {
	var ()
	return &WaypointQueueJobParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointQueueJobParamsWithTimeout creates a new WaypointQueueJobParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointQueueJobParamsWithTimeout(timeout time.Duration) *WaypointQueueJobParams {
	var ()
	return &WaypointQueueJobParams{

		timeout: timeout,
	}
}

// NewWaypointQueueJobParamsWithContext creates a new WaypointQueueJobParams object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointQueueJobParamsWithContext(ctx context.Context) *WaypointQueueJobParams {
	var ()
	return &WaypointQueueJobParams{

		Context: ctx,
	}
}

// NewWaypointQueueJobParamsWithHTTPClient creates a new WaypointQueueJobParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointQueueJobParamsWithHTTPClient(client *http.Client) *WaypointQueueJobParams {
	var ()
	return &WaypointQueueJobParams{
		HTTPClient: client,
	}
}

/*WaypointQueueJobParams contains all the parameters to send to the API endpoint
for the waypoint queue job operation typically these are written to a http.Request
*/
type WaypointQueueJobParams struct {

	/*Body*/
	Body *models.HashicorpWaypointQueueJobRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint queue job params
func (o *WaypointQueueJobParams) WithTimeout(timeout time.Duration) *WaypointQueueJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint queue job params
func (o *WaypointQueueJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint queue job params
func (o *WaypointQueueJobParams) WithContext(ctx context.Context) *WaypointQueueJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint queue job params
func (o *WaypointQueueJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint queue job params
func (o *WaypointQueueJobParams) WithHTTPClient(client *http.Client) *WaypointQueueJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint queue job params
func (o *WaypointQueueJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint queue job params
func (o *WaypointQueueJobParams) WithBody(body *models.HashicorpWaypointQueueJobRequest) *WaypointQueueJobParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint queue job params
func (o *WaypointQueueJobParams) SetBody(body *models.HashicorpWaypointQueueJobRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointQueueJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}