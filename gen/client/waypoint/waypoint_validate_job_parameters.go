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

// NewWaypointValidateJobParams creates a new WaypointValidateJobParams object
// with the default values initialized.
func NewWaypointValidateJobParams() *WaypointValidateJobParams {
	var ()
	return &WaypointValidateJobParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointValidateJobParamsWithTimeout creates a new WaypointValidateJobParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointValidateJobParamsWithTimeout(timeout time.Duration) *WaypointValidateJobParams {
	var ()
	return &WaypointValidateJobParams{

		timeout: timeout,
	}
}

// NewWaypointValidateJobParamsWithContext creates a new WaypointValidateJobParams object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointValidateJobParamsWithContext(ctx context.Context) *WaypointValidateJobParams {
	var ()
	return &WaypointValidateJobParams{

		Context: ctx,
	}
}

// NewWaypointValidateJobParamsWithHTTPClient creates a new WaypointValidateJobParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointValidateJobParamsWithHTTPClient(client *http.Client) *WaypointValidateJobParams {
	var ()
	return &WaypointValidateJobParams{
		HTTPClient: client,
	}
}

/*WaypointValidateJobParams contains all the parameters to send to the API endpoint
for the waypoint validate job operation typically these are written to a http.Request
*/
type WaypointValidateJobParams struct {

	/*Body*/
	Body *models.HashicorpWaypointValidateJobRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint validate job params
func (o *WaypointValidateJobParams) WithTimeout(timeout time.Duration) *WaypointValidateJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint validate job params
func (o *WaypointValidateJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint validate job params
func (o *WaypointValidateJobParams) WithContext(ctx context.Context) *WaypointValidateJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint validate job params
func (o *WaypointValidateJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint validate job params
func (o *WaypointValidateJobParams) WithHTTPClient(client *http.Client) *WaypointValidateJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint validate job params
func (o *WaypointValidateJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint validate job params
func (o *WaypointValidateJobParams) WithBody(body *models.HashicorpWaypointValidateJobRequest) *WaypointValidateJobParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint validate job params
func (o *WaypointValidateJobParams) SetBody(body *models.HashicorpWaypointValidateJobRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointValidateJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
