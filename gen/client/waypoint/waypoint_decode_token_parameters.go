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

// NewWaypointDecodeTokenParams creates a new WaypointDecodeTokenParams object
// with the default values initialized.
func NewWaypointDecodeTokenParams() *WaypointDecodeTokenParams {
	var ()
	return &WaypointDecodeTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointDecodeTokenParamsWithTimeout creates a new WaypointDecodeTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointDecodeTokenParamsWithTimeout(timeout time.Duration) *WaypointDecodeTokenParams {
	var ()
	return &WaypointDecodeTokenParams{

		timeout: timeout,
	}
}

// NewWaypointDecodeTokenParamsWithContext creates a new WaypointDecodeTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointDecodeTokenParamsWithContext(ctx context.Context) *WaypointDecodeTokenParams {
	var ()
	return &WaypointDecodeTokenParams{

		Context: ctx,
	}
}

// NewWaypointDecodeTokenParamsWithHTTPClient creates a new WaypointDecodeTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointDecodeTokenParamsWithHTTPClient(client *http.Client) *WaypointDecodeTokenParams {
	var ()
	return &WaypointDecodeTokenParams{
		HTTPClient: client,
	}
}

/*WaypointDecodeTokenParams contains all the parameters to send to the API endpoint
for the waypoint decode token operation typically these are written to a http.Request
*/
type WaypointDecodeTokenParams struct {

	/*Body*/
	Body *models.HashicorpWaypointDecodeTokenRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint decode token params
func (o *WaypointDecodeTokenParams) WithTimeout(timeout time.Duration) *WaypointDecodeTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint decode token params
func (o *WaypointDecodeTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint decode token params
func (o *WaypointDecodeTokenParams) WithContext(ctx context.Context) *WaypointDecodeTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint decode token params
func (o *WaypointDecodeTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint decode token params
func (o *WaypointDecodeTokenParams) WithHTTPClient(client *http.Client) *WaypointDecodeTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint decode token params
func (o *WaypointDecodeTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint decode token params
func (o *WaypointDecodeTokenParams) WithBody(body *models.HashicorpWaypointDecodeTokenRequest) *WaypointDecodeTokenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint decode token params
func (o *WaypointDecodeTokenParams) SetBody(body *models.HashicorpWaypointDecodeTokenRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointDecodeTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
